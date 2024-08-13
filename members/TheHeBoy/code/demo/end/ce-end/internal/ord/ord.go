package ord

import (
	"fmt"
	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/mempool"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
	"gohub/pkg/btcapi"
)

type InscriptionData struct {
	ContentType string
	Body        []byte
	Destination string
}

type InscriptionRequest struct {
	TxOutPointList []*wire.OutPoint
	TxOutList      []*wire.TxOut
	TxPrivateKey   *btcec.PrivateKey
	FeeRate        int64
	Data           InscriptionData
	RevealOutValue int64
	ChargeFee      int64
}

type InscriptionTool struct {
	net                 *chaincfg.Params
	client              btcapi.BtcApiClient
	feeRate             int64
	txPrevOutputFetcher *txscript.MultiPrevOutFetcher
	txPrivateKey        *btcec.PrivateKey
	txCtxData           *inscriptionTxCtxData
	revealTx            *wire.MsgTx
}

type inscriptionTxCtxData struct {
	address             *btcutil.AddressTaproot
	inscriptionScript   []byte
	controlBlockWitness []byte
}

const (
	defaultSequenceNum    = wire.MaxTxInSequenceNum - 10
	defaultRevealOutValue = int64(546)

	MaxStandardTxWeight = blockchain.MaxBlockWeight / 10
)

func NewInscriptionTool(request *InscriptionRequest) (*InscriptionTool, error) {
	tool := &InscriptionTool{
		net:                 btcapi.NetParams,
		client:              btcapi.Client,
		feeRate:             request.FeeRate,
		txPrevOutputFetcher: txscript.NewMultiPrevOutFetcher(nil),
		txPrivateKey:        request.TxPrivateKey,
	}
	return tool, tool._initTool(btcapi.NetParams, request)
}

func (tool *InscriptionTool) _initTool(net *chaincfg.Params, request *InscriptionRequest) error {
	revealOutValue := defaultRevealOutValue
	if request.RevealOutValue > 0 {
		revealOutValue = request.RevealOutValue
	}
	txCtxData, err := createInscriptionTxCtxData(net, request.TxPrivateKey, request.Data)
	if err != nil {
		return err
	}
	tool.txCtxData = txCtxData

	destination, err := btcutil.DecodeAddress(request.Data.Destination, net)
	if err != nil {
		return err
	}

	err = tool.buildRevealTx(destination, revealOutValue, request)
	if err != nil {
		return err
	}

	return err
}

func createInscriptionTxCtxData(net *chaincfg.Params, privateKey *btcec.PrivateKey, data InscriptionData) (*inscriptionTxCtxData, error) {
	inscriptionBuilder := txscript.NewScriptBuilder().
		AddData(schnorr.SerializePubKey(privateKey.PubKey())).
		AddOp(txscript.OP_CHECKSIG).
		AddOp(txscript.OP_FALSE).
		AddOp(txscript.OP_IF).
		AddData([]byte("ord")).
		// Two OP_DATA_1 should be OP_1. However, in the following link, it's not set as OP_1:
		// https://github.com/casey/ord/blob/0.5.1/src/inscription.rs#L17
		// Therefore, we use two OP_DATA_1 to maintain consistency with go-ord-tx.
		AddOp(txscript.OP_DATA_1).
		AddOp(txscript.OP_DATA_1).
		AddData([]byte(data.ContentType)).
		AddOp(txscript.OP_0)
	maxChunkSize := 520
	bodySize := len(data.Body)
	for i := 0; i < bodySize; i += maxChunkSize {
		end := i + maxChunkSize
		if end > bodySize {
			end = bodySize
		}
		// to skip txscript.MaxScriptSize 10000
		inscriptionBuilder.AddFullData(data.Body[i:end])
	}
	inscriptionScript, err := inscriptionBuilder.Script()
	if err != nil {
		return nil, err
	}
	// to skip txscript.MaxScriptSize 10000
	inscriptionScript = append(inscriptionScript, txscript.OP_ENDIF)

	leafNode := txscript.NewBaseTapLeaf(inscriptionScript)
	proof := &txscript.TapscriptProof{
		TapLeaf:  leafNode,
		RootNode: leafNode,
	}

	tapHash := proof.RootNode.TapHash()
	address, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootOutputKey(privateKey.PubKey(), tapHash[:])), net)
	if err != nil {
		return nil, err
	}

	controlBlock := proof.ToControlBlock(privateKey.PubKey())
	controlBlockWitness, err := controlBlock.ToBytes()
	if err != nil {
		return nil, err
	}

	return &inscriptionTxCtxData{
		address:             address,
		inscriptionScript:   inscriptionScript,
		controlBlockWitness: controlBlockWitness,
	}, nil
}

func CreateAccount(net *chaincfg.Params, data InscriptionData) (*btcec.PrivateKey, *btcutil.AddressTaproot, error) {
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, nil, err
	}
	ctxData, err := createInscriptionTxCtxData(net, privateKey, data)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, ctxData.address, nil
}

func (tool *InscriptionTool) buildRevealTx(destination btcutil.Address, revealOutValue int64, request *InscriptionRequest) error {
	// 构建输入
	totalSenderAmount := btcutil.Amount(0)
	tx := wire.NewMsgTx(wire.TxVersion)
	for i := range request.TxOutList {
		txOut := request.TxOutList[i]
		txOutPoint := request.TxOutPointList[i]
		tool.txPrevOutputFetcher.AddPrevOut(*txOutPoint, txOut)
		in := wire.NewTxIn(txOutPoint, nil, nil)
		in.Sequence = defaultSequenceNum
		tx.AddTxIn(in)
		totalSenderAmount += btcutil.Amount(txOut.Value)
	}

	// 构建输出
	totalOutputAmount := btcutil.Amount(0)
	pkScript, err := txscript.PayToAddrScript(destination)
	if err != nil {
		return errors.WithStack(err)
	}
	tx.AddTxOut(wire.NewTxOut(revealOutValue, pkScript))
	// service fee, rest sats is gas fee
	err = tool.chargeServiceFee(request.ChargeFee, tx)
	if err != nil {
		return err
	}
	for i := range tx.TxOut {
		totalOutputAmount += btcutil.Amount(tx.TxOut[i].Value)
	}

	// 添加witness
	for i := range tx.TxIn {
		if i == 0 {
			witnessArray, err := txscript.CalcTapscriptSignaturehash(txscript.NewTxSigHashes(tx, tool.txPrevOutputFetcher),
				txscript.SigHashDefault, tx, i, tool.txPrevOutputFetcher, txscript.NewBaseTapLeaf(tool.txCtxData.inscriptionScript))
			if err != nil {
				return errors.WithStack(err)
			}

			signature, err := schnorr.Sign(tool.txPrivateKey, witnessArray)
			if err != nil {
				return errors.WithStack(err)
			}
			tx.TxIn[i].Witness = wire.TxWitness{signature.Serialize(), tool.txCtxData.inscriptionScript, tool.txCtxData.controlBlockWitness}
		} else {
			txOut := tool.txPrevOutputFetcher.FetchPrevOutput(tx.TxIn[i].PreviousOutPoint)
			witness, err := txscript.TaprootWitnessSignature(tx, txscript.NewTxSigHashes(tx, tool.txPrevOutputFetcher),
				i, txOut.Value, txOut.PkScript, txscript.SigHashDefault, tool.txPrivateKey)
			if err != nil {
				return err
			}
			tx.TxIn[i].Witness = witness
		}
	}

	fee := btcutil.Amount(mempool.GetTxVirtualSize(btcutil.NewTx(tx))) * btcutil.Amount(request.FeeRate)
	changeAmount := totalSenderAmount - totalOutputAmount - fee
	if changeAmount < 0 {
		return errors.New("insufficient balance")
	}

	tool.revealTx = tx
	return nil
}

func (tool *InscriptionTool) chargeServiceFee(chargeFee int64, tx *wire.MsgTx) error {
	if chargeFee == 0 {
		return nil
	}
	if chargeFee <= 500 {
		return errors.New("charge fee must be greater than 500")
	}
	pkScript, err := txscript.PayToAddrScript(btcapi.ChargeAddress)
	if err != nil {
		return errors.WithStack(err)
	}

	tx.AddTxOut(wire.NewTxOut(chargeFee, pkScript))
	return nil
}

func (tool *InscriptionTool) CalculateFee() int64 {
	fees := int64(0)
	for _, in := range tool.revealTx.TxIn {
		fees += tool.txPrevOutputFetcher.FetchPrevOutput(in.PreviousOutPoint).Value
	}
	for _, out := range tool.revealTx.TxOut {
		fees -= out.Value
	}
	return fees
}

func (tool *InscriptionTool) EstimateFee() int64 {
	var total int64

	for _, out := range tool.revealTx.TxOut {
		total += out.Value
	}

	return mempool.GetTxVirtualSize(btcutil.NewTx(tool.revealTx))*tool.feeRate + total
}

func (tool *InscriptionTool) sendRawTransaction(tx *wire.MsgTx) (*chainhash.Hash, error) {
	return tool.client.BroadcastTx(tx)
}

func (tool *InscriptionTool) Inscribe() (revealTxHash *chainhash.Hash, inscriptionId string, fees int64, err error) {
	fees = tool.CalculateFee()
	revealTxHash, err = tool.sendRawTransaction(tool.revealTx)
	if err != nil {
		return nil, "", fees, errors.WithMessage(err, "send reveal tx error")
	}
	inscriptionId = fmt.Sprintf("%si0", revealTxHash)

	return revealTxHash, inscriptionId, fees, nil
}
