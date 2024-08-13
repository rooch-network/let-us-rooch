// Copyright (c) RoochNetwork
// SPDX-License-Identifier: Apache-2.0
// Author: Jason Jo

import {
    useConnectWallet,
    useCreateSessionKey,
    useCurrentAddress,
    UseSignAndExecuteTransaction,
    useWallets,
    useWalletStore
} from "@roochnetwork/rooch-sdk-kit";
import "./App.css";
import {shortAddress} from "./utils.ts";
import {Badge, Button, Card, Divider, message} from "antd";
import {Icon} from "@iconify/react";
import React, {useEffect, useState} from "react";
import CountUp from "react-countup";
import * as Api from "./request/all";
import * as MemAPi from "./request/mempool"
import {Args, ThirdPartyAddress, Transaction} from "@roochnetwork/rooch-sdk";
import * as Rpc from "./rpc"

const baseUrl = import.meta.env.VITE_API_BASEURL;
const iframeSrc = `${baseUrl}/seed/seedHTML`;

// const theme1 = "#f2be45";
const theme2 = "#52616d";

const contractAddr = import.meta.env.VITE_CONTRACT_ADDR;

function TitleDivider({title, end}: { title: string, end?: React.ReactNode }) {
    return (
        <div>
            <div className="flex items-center text-2xl">{title} {end}</div>
            <Divider className="m-0 mt-2 mb-6 bg-white"/>
        </div>
    );
}

function Mint({address, onMintComplete}: { address: ThirdPartyAddress | undefined, onMintComplete?: () => void }) {
    const [messageApi] = message.useMessage();

    const [loading, setLoading] = useState(false)
    const [mintState, setMinted] = useState("MINT")


    const addr = address?.toStr()
    const handleClick = async () => {
        try {
            setLoading(true)
            const gasFee = Math.ceil((await MemAPi.getFeeRate()).economyFee)
            const saveOrderResp = await Api.saveOrder(addr!, gasFee)
            await window.unisat.sendBitcoin(
                saveOrderResp.payAddress,
                saveOrderResp.estimateFee,
            );
            await Api.executeOrder(saveOrderResp.orderId);
            setMinted("MINTED");
            if (onMintComplete) {
                onMintComplete();
            }
            messageApi.open({
                type: 'success',
                content: 'Mint inscription success',
            });
        } finally {
            setLoading(false)
        }
    };


    return (
        <Button disabled={!addr || mintState == "MINTED"} type="link" loading={loading} onClick={handleClick}
                className="font-bold  text-xl">{mintState}</Button>
    )
}

interface OrderListResp {
    num: number;
    bitcoin_tx: string;
    hSeed: string;
    coinNum: number;
    isOpen: boolean;
    objectId: string;
}

function App() {
    const wallets = useWallets();
    const currentAddress = useCurrentAddress();
    const connectionStatus = useWalletStore((state) => {
        return state.connectionStatus;
    });
    const setWalletDisconnected = useWalletStore(
        (state) => state.setWalletDisconnected
    );
    const {mutateAsync: connectWallet} = useConnectWallet();
    const {mutateAsync: createSessionKey} = useCreateSessionKey();
    const {mutateAsync: signAndExecuteTransaction} = UseSignAndExecuteTransaction();

    const [randomUsableSeedState, setRandomUsableSeedState] = useState({
        num: 0,
        hSeed: "",
    });
    const [profileState, setProfileState] = useState<OrderListResp[]>([])
    const [balanceState, setBalanceState] = useState("0")
    const [poolState, setPoolState] = useState("0")
    const [blockHeight, setBlockHeight] = useState(0);

    useEffect(() => {
        // 定义一个异步函数
        const fetchRandomUsableSeed = async () => {
            if (connectionStatus === 'connected' && currentAddress) {
                setBalanceState(await Rpc.getBalance(currentAddress.genRoochAddress().toStr()));
                setProfileState(await Api.orderList(currentAddress.toStr()));
                try {
                    const seed = await Api.randomUsableSeed(currentAddress.toStr());
                    setRandomUsableSeedState(seed);
                } catch (error) {
                    console.error("Error fetching random usable seed:", error);
                }
            }
        };

        // 调用异步函数
        fetchRandomUsableSeed();
    }, [connectionStatus, currentAddress]);

    useEffect(() => {
        task();
    })

    // Refresh profile when profileState changes
    useEffect(() => {
        refreshProfile();
    }, [profileState]);

    const refreshProfile = async () => {
        if (currentAddress) {
            setBalanceState(await Rpc.getBalance(currentAddress.genRoochAddress().toStr()));
        }
        if (profileState.length > 0) {
            const stateList = await Rpc.checkInscriptionsState(profileState.filter(item => item.objectId).map(item => item.objectId))
            setProfileState(prevState =>
                prevState.map((item, i) => {
                    // 确保索引 i 在 stateList 中是有效的
                    const state = i < stateList.length ? stateList[i].value : {};
                    return {
                        ...item,
                        coinNum: state.coin_num || 0,
                        isOpen: state.is_open
                    };
                })
            );
        }
    }

    const task = async () => {
        setPoolState(await Rpc.getPoolCapacity())
        setBlockHeight(await Rpc.getLastBlockHeight());
    }

    setInterval(task, 60000);

    return (
        <div>
            <div className="w-full bg-white text-center">current block height: {blockHeight}</div>
            <div className="m-auto p-4 max-w-[1280px] min-w-[1280px]">
                <div className="flex w-full justify-between">
                    <div className="flex items-center">
                        <div className="text-4xl font-bold">ColorEgg</div>
                    </div>
                    <div>
                        <Button
                            size="large"
                            onClick={async () => {
                                if (connectionStatus === "connected") {
                                    setWalletDisconnected();
                                    return;
                                }
                                await connectWallet({wallet: wallets[0]});
                            }}
                        >
                            {connectionStatus === "connected"
                                ? shortAddress(currentAddress?.genRoochAddress().toStr(), 8, 6)
                                : "Connect Wallet"}
                        </Button>
                    </div>
                </div>
                <Divider className="bg-white"/>
                <div className="flex justify-center">
                    <Badge color={theme2} count="100+">
                        <Card className="w-[300px]">
                            <div className="text-theme2 text-4xl  flex">
                                <Icon className="text-theme1"
                                      icon="ph:coins"/>
                                <span className="m-auto">
                                    <CountUp
                                        start={0}
                                        end={Number(poolState)}
                                        separator=","
                                    />
                                </span>
                            </div>
                        </Card>
                    </Badge>
                </div>

                <div className="flex-col space-y-4">
                    <TitleDivider title="MINT"/>
                    <div className="flex justify-between">
                        <div className="w-1/2 ">
                            <p className="font-bold text-2xl mb-4">Color Egg grow up base on Bitcoin with Rooch</p>
                            <p className="text-pretty text-xl leading-loose">
                                Color Egg 是基于 Bitcoin 上的 Ordinals 协议 和 Rooch 链的项目。通过 Rooch 链上可以读取
                                Bitcoin 上状态的特点，用户可以在 Bitcoin 上 mint ColorEgg 铭文，
                                每个持有 ColorEgg 铭文的账户，可以在 Rooch 链上打开它，以此来得到 Egg Coin。每个 Egg 中都存储着
                                Coin，其 Coin 数量
                                随着区块高度的增加，存储的 Coin 就会越多，用户砸开 ColorEgg 后就可以得到其中的 Coin，但是
                                Coin 的数量不能大于 Coin Pool，
                                超出的部分用户并不能拿到，而 Coin Pool 的容量是每增加一个区块高度，就会加入 100 Coin。
                            </p>
                        </div>
                        <Card hoverable className="w-[250px]"
                              cover={<iframe className="h-[250px]"
                                             src={`${iframeSrc}/${randomUsableSeedState.hSeed || "4yrD3vWpBAqvzVlm"}`}/>}
                              actions={[<Mint onMintComplete={async () =>
                                  setProfileState(await Api.orderList(currentAddress!.toStr()))
                              }
                                              address={currentAddress}/>]}
                        >
                            <div className="flex font-bold text-theme2 justify-between">
                                <span className="truncate w-1/2"># {randomUsableSeedState.num}</span>
                                <span>COLOR EGG</span>
                            </div>
                        </Card>
                    </div>
                    <div>
                        <TitleDivider title="PROFILE" end={
                            <div className="flex items-center ml-4">
                                <Icon className="text-theme2" icon="ph:coins"/> {balanceState}
                            </div>}/>
                        {
                            <div className="flex flex-wrap gap-2">
                                {profileState.map((item) => (
                                    <Card
                                        key={item.num}
                                        hoverable
                                        className="w-[200px]"
                                        cover={
                                            <iframe className="h-[200px]" src={`${iframeSrc}/${item.hSeed}`}/>
                                        }
                                    >
                                        <div className="text-theme2 flex items-center justify-between">
                                            <div className="flex-col justify-between">
                                                <div className="flex items-center">
                                                    <Icon className="text-theme1 mr-2"
                                                          icon="fluent:number-symbol-16-filled"/> {item.num}
                                                </div>
                                                <div className="flex items-center">
                                                    <Icon className="text-theme1 mr-2" icon="ph:coins"/>
                                                    <span className="text-end">{item.coinNum}</span>
                                                </div>
                                            </div>
                                            <Button onClick={async () => {
                                                const defaultScopes = [`${contractAddr}::*`];

                                                await createSessionKey(
                                                    {
                                                        appName: "Color Egg",
                                                        appUrl: "http://localhost:3000",
                                                        maxInactiveInterval: 1000,
                                                        scopes: defaultScopes,
                                                    },
                                                    {
                                                        onError: (why) => {
                                                            console.log(why);
                                                        },
                                                    }
                                                )

                                                const txn = new Transaction();
                                                txn.callFunction({
                                                    address: contractAddr.split("::")[0],
                                                    module: contractAddr.split("::")[1],
                                                    function: "open",
                                                    args: [
                                                        Args.objectId(item.objectId),
                                                        Args.objectId("0x2b714929f9017449a2aabf1682c3e82ddd29bc77afbca40b85439c017c46bac7"),
                                                    ],
                                                });
                                                await signAndExecuteTransaction({transaction: txn});

                                            }} disabled={!item.objectId || item.isOpen} type="link"
                                                    className="p-0 rotating-box">
                                                {
                                                    item.isOpen ?
                                                        <Icon className="text-4xl text-gray-400"
                                                              icon={"solar:sledgehammer-bold"}/> :
                                                        <Icon className="text-theme1 text-4xl hover:text-[#72b0fe]"
                                                              icon={item.objectId ? "solar:sledgehammer-bold" : "formkit:time"}/>
                                                }
                                            </Button>
                                        </div>
                                    </Card>
                                ))}
                            </div>
                        }

                    </div>
                </div>
            </div>
        </div>
    );
}

export default App;
