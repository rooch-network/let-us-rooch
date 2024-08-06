# Rooch 与 btc

## 客户端验证

在矿工打包区块时会进行交易验证，并将新区块广播到整个网络节点中。当全节点或轻节点客户端接收到新区块时，如何保证接收到的新区块是有效的。例如有欺诈节点发送了一个区块到其已知的全节点 ip, 如果该全节点不进行验证，那么其他人调用该全节点的 rpc 服务必然也会出错。

所以需要进行验证，这种验证就称作客户端验证。如果验证通过则会将新区块加入客户端的区块链副本中，否则拒绝这个新区块。

验证内容包括验证区块头、UTXO 有效性(避免双花)、`scriptSig` 是否可以解开 `scriptPubKey`、`coinbase` 交易、区块大小及容量等。

更进一步, `RGB` 协议亦使用客户端验证模型来对比特币进行拓展，将 `RGB` 合约执行和状态存储放在链下, 只将状态变更的证明记录在比特币交易中(`OP_RETURN`)。

因此 `RGB` 客户端在进行客户端验证时, 可以读取区块中包含 `OP_RETURN` 的 `UTXO`, 获取内容并判断是否属于 `RGB` 交易, 并根据原始的比特币交易判断在 `RGB` 上的资产归属。

## Sequencer

在 `layer2` 中, 通过排序器接收交易并将交易放入待处理的交易序列中，同时为交易序列中的每笔交易生成一个**排序证明**。这个证明包含交易在序列中的位置信息,以及能验证这个位置的加密证据。排序器处理交易序列并将交易序列打包(`batches`)发送到 `layer1`

排序器的作用是为了保证交易顺序, 防止 `MEV`。例如某笔交易已获取到排序证明，但最终链上没有该笔交易或被 `MEV` 攻击。此时可以用排序证明申请仲裁。

## Stackable L2

堆叠式 L2, `Rooch` 同步整个比特币的数据到 `Rooch` 网络。同时会将这些数据转化为 `Object`, 包括 `Block`

```rust
struct BitcoinBlockStore has key{
    /// The genesis start block
    genesis_block: BlockHeightHash,
    latest_block: Option<BlockHeightHash>,
    /// block hash -> block header
    blocks: Table<address, Header>,
    /// block height -> block hash
    height_to_hash: Table<u64, address>,
    /// block hash -> block height
    hash_to_height: Table<address, u64>,
    /// tx id -> tx
    txs: Table<address, Transaction>,
    /// tx id -> block height
    tx_to_height: Table<address, u64>,
    /// tx id list, we can use this to scan txs
    tx_ids: TableVec<address>,
}
```

以及`UTXO`

```rust
struct UTXO has key {
    /// The txid of the UTXO
    txid: address,
    /// The vout of the UTXO
    vout: u32,
    /// The value of the UTXO
    value: u64,
    /// Protocol seals
    seals: SimpleMultiMap<String, ObjectID>
}
```

比特币交易输出 `UTXO` 集合或数组

- `txid` 交易集合
- `vout` 在 `UTXO` 集中的索引
- `value` 存储聪的数量

## Rooch 合约读取 UTXO

下面示例读取交易中的包括 `OP_RETURN` 的 `UTXO`

> 每个比特币交易输出中只能包含一个拥有 OP_RETURN 的 UTXO

```rust
fun read_tx_op_return_msg(txid: address) {
    // 根据 txid 获取交易
    let tx = option::destroy_some(bitcoin::get_tx(txid));
    // 获取交易输出
    let outputs = tx_output(&tx);
    // 获取该笔交易输出的 utxo 数量
    let len = vector::length(outputs);
    // 遍历 utxo
    while (len > 0){
        let utxo = vector::borrow(outputs, len - 1);
        // 获取 utxo 中的 scriptPubKey
        let script_pubkey = txout_script_pubkey(utxo);
        // 判断 scriptPubKey 是否包含 OP_RETURN
        if(is_op_return(script_pubkey)){
            // 读取 OP_RETURN 后面的信息
            let msg = witness_program(script_pubkey);
        };
        len = len - 1;
    };
}
```
