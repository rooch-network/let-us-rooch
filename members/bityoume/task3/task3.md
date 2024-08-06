# 任务3

## 3.1 搭建本地 Bitcoin 以及 Rooch 开发环境

>   https://github.com/rooch-network/rooch/blob/main/scripts/bitcoin/README.md

### 启动本地bitcoin节点

```bash
$ cd scripts/bitcoin/ && ./node/run_local_node_docker.sh
Error response from daemon: No such container: bitcoind_regtest
Error response from daemon: No such container: bitcoind_regtest
Unable to find image 'lncm/bitcoind:v25.1' locally
v25.1: Pulling from lncm/bitcoind
579b34f0a95b: Pull complete
c749c9fee6fb: Pull complete
af31b96e4577: Pull complete
0639e51d2f37: Pull complete
c2c84033476e: Pull complete
Digest: sha256:6562182e029221fe21c352f138540d8016963671c31b376e2ebad84914d9bed3
Status: Downloaded newer image for lncm/bitcoind:v25.1
43c9fa341e8365d8850dab35906db110d4cca2fc46da72d54e8c58742a90de2c
```

### 设置命令别名

```bash
$ source ./cmd_alias.sh
```

### rooch初始化

```bash
$ rooch init
Creating client config file ["~/.rooch/rooch_config/rooch.yaml"].
Enter a password to encrypt the keys. Press enter to leave it an empty password:
Generated new keypair for address [rooch1umumkfcuew6fdxunvghrsqlgqlygpdxtqc5sgs478czaaj9nc97syxcqhr]
Secret Recovery Phrase : [suffer tattoo funny manage face word fog ask taste swear peace prevent]
Rooch client config file generated at ~/.rooch/rooch_config/rooch.yaml
```

### 本地启动rooch服务并连接本地bitcoin节点

```bash
$ rooch server start --btc-rpc-url http://127.0.0.1:18443 --btc-rpc-username roochuser --btc-rpc-password roochpass

2024-08-03T15:13:54.564481Z  INFO moveos_common::utils: set max open fds 8192
2024-08-03T15:13:54.682624Z  INFO moveos::moveos: execute genesis tx state_root:0x8af77116cbb8cc5506ab3b7198366126fd9d057e7248b1fcbe8c2d00cd4060b9, state_size:32
2024-08-03T15:13:54.749489Z  INFO moveos::moveos: execute genesis tx state_root:0x8af77116cbb8cc5506ab3b7198366126fd9d057e7248b1fcbe8c2d00cd4060b9, state_size:32
2024-08-03T15:13:54.755423Z  INFO rooch_rpc_server: The latest Root object state root: 0x8af77116cbb8cc5506ab3b7198366126fd9d057e7248b1fcbe8c2d00cd4060b9, size: 32
2024-08-03T15:13:54.757793Z  INFO rooch_rpc_server: RPC Server sequencer address: rooch1umumkfcuew6fdxunvghrsqlgqlygpdxtqc5sgs478czaaj9nc97syxcqhr(0xe6f9bb271ccbb4969b93622e3803e807c880b4cb06290442be3e05dec8b3c17d)
2024-08-03T15:13:54.757816Z  INFO rooch_sequencer::actor::sequencer: Load latest sequencer order 0
2024-08-03T15:13:54.758011Z  INFO rooch_sequencer::actor::sequencer: Load latest sequencer accumulator info AccumulatorInfo { accumulator_root: 0x3bc8e3054011101a591a951ceae3e413771b8e3e1337dc9c8dc1f172c63c6880, frozen_subtree_roots: [0x3bc8e3054011101a591a951ceae3e413771b8e3e1337dc9c8dc1f172c63c6880], num_leaves: 1, num_nodes: 1 }
2024-08-03T15:13:54.758418Z  INFO rooch_rpc_server: RPC Server proposer address: rooch1umumkfcuew6fdxunvghrsqlgqlygpdxtqc5sgs478czaaj9nc97syxcqhr(0xe6f9bb271ccbb4969b93622e3803e807c880b4cb06290442be3e05dec8b3c17d)
2024-08-03T15:13:54.783011Z  INFO rooch_relayer::actor::relayer: BitcoinRelayer started
2024-08-03T15:13:54.783069Z  INFO rooch_rpc_server: acl=Const("*")
2024-08-03T15:13:54.783925Z  INFO rooch_rpc_server: JSON-RPC HTTP Server start listening 0.0.0.0:6767
2024-08-03T15:13:54.783931Z  INFO rooch_rpc_server: Available JSON-RPC methods : ["btc_broadcastTX", "rooch_getObjectStates", "rooch_queryObjectStates", "rooch_getBalance", "rooch_sendRawTransaction", "rooch_getTransactionsByHash", "rooch_queryEvents", "rooch_executeRawTransaction", "rooch_getEventsByEventHandle", "rooch_queryTransactions", "rooch_executeViewFunction", "rooch_listFieldStates", "rooch_getBalances", "rooch_getModuleABI", "rooch_listStates", "rooch_getTransactionsByOrder", "btc_queryInscriptions", "rooch_getChainID", "rooch_getFieldStates", "btc_queryUTXOs", "rpc.discover", "rooch_getStates"]
The active env was successfully switched to `local`
2024-08-03T15:13:55.795062Z  INFO rooch_relayer::actor::bitcoin_relayer: BitcoinRelayer process block, height: 0, hash: 0f9188f13cb7b2c71f2a335e3a4fc328bf5beb436012afca590b1a11466e2206, tx_size: 1, time: 1296688602
2024-08-03T15:13:55.796833Z  INFO rooch_sequencer::actor::sequencer: sequencer tx: 0x10d5…6d8e order: 1
2024-08-03T15:13:55.825186Z  INFO rooch_relayer::actor::relayer: Relayer execute relay block(hash: 06226e46111a0b59caaf126043eb5bbf28c34f3a5e332a1fc7b2b73cf188910f, height: 0) success
2024-08-03T15:13:55.826345Z  INFO rooch_sequencer::actor::sequencer: sequencer tx: 0x771d…7348 order: 2
2024-08-03T15:13:55.827402Z  INFO rooch_relayer::actor::relayer: Relayer execute relay tx(txid: 3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a) success
2024-08-03T15:13:59.761222Z  INFO rooch_proposer::actor::proposer: [ProposeBlock] block_number: 0, batch_size: 2
```

### 获取bitcoin地址

```json
rooch account list --json
{
  "default": {
    "address": "rooch1umumkfcuew6fdxunvghrsqlgqlygpdxtqc5sgs478czaaj9nc97syxcqhr",
    "hex_address": "0xe6f9bb271ccbb4969b93622e3803e807c880b4cb06290442be3e05dec8b3c17d",
    "bitcoin_address": "bcrt1pvxgaqz6fhd4l4j0fv25mdyuqz55469hxhntdluwm2syzpyym7nqsqrrtwz",
    "nostr_public_key": "npub1vtcsrwfq9te7mvtvjdm9hn6w9l4g7k8wrav03d5dwhrvvw9u7z3q8s7mxn",
    "public_key": "AQJi8QG5ICrz7bFsk3Zbz04v6o9Y7h9Y+LaNdcbGOLzwog==",
    "has_session_key": false,
    "active": true
  }
}
```

### 使用该地址构造比特币区块

```bash
$ bitcoin-cli generatetoaddress 101 bcrt1pvxgaqz6fhd4l4j0fv25mdyuqz55469hxhntdluwm2syzpyym7nqsqrrtwz

[
  "3ca0f6c288474a4c74eee00a9dff5d39c3a501176f43ac8817ec88195b178ff9",
  "505f28704f2cceacd4cbf31a2a8dad4adc7cafc8e9878ae2bdd6178308db2be6",
  "64f9d2ab76f8124e8c3cf2239c60fe243f9b6536007817326f0fea132330d172",
  "2edd7149d5a67b06858260af45ac4164383d4688a2311bfe37259ee83d27b226",
  "4725c2cccd94de8f82efd41dc7b72d46a6d5e5ac538f83df0ca425c4ee3673ab",
  "51af98f0194a1fc127e678376050f87c3b7d782d0614522fdc2e753858b51374",
  ......
  "53a78c104b490631e5869d25c7f8349f15e1d3d44496af761080b1e00d4e8fe3",
  "4d2cac160234dcba6f66d4979305ef6046380650d3bc971a168eaa21def7ba6c"
]
```

### 查询对象状态

```bash
$ rooch rpc request --method rooch_queryObjectStates --params '[{"object_type":"0x4::utxo::UTXO"},  null, "2", {"descending": true,"showDisplay":false}]'

Data:
----------------------------------------------------------------------------------------------------
  objectId       | 0x826a5e56581ba5ab84c39976f27cf3578cf524308b4ffc123922dfff507e514dc286ad491cb19a79e63441398862c9aa006cc4550af8b7301b85c8fff97d7a03
  type           | 0x4::utxo::UTXO
  owner          | rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzqf47q0l
  owner(bitcoin) | None
  status         | UserOwned
  tx_order       | 204
  state_index    | 10
----------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------
  objectId       | 0x826a5e56581ba5ab84c39976f27cf3578cf524308b4ffc123922dfff507e514d141c06077a9b326bca31c9f1922b5cb361a52dcde5eab30d5ea74176150748df
  type           | 0x4::utxo::UTXO
  owner          | rooch1umumkfcuew6fdxunvghrsqlgqlygpdxtqc5sgs478czaaj9nc97syxcqhr
  owner(bitcoin) | Some("bcrt1pvxgaqz6fhd4l4j0fv25mdyuqz55469hxhntdluwm2syzpyym7nqsqrrtwz")
  status         | UserOwned
  tx_order       | 204
  state_index    | 9
----------------------------------------------------------------------------------------------------

Next cursor:
    IndexerStateID[tx order: 204, state index: 9]

Has next page: true
```

## 3.2 部署一个和 Bitcoin 数据交互的 Move 智能合约

```rust
rooch move publish
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
INCLUDING DEPENDENCY BitcoinMove
INCLUDING DEPENDENCY MoveStdlib
INCLUDING DEPENDENCY MoveosStdlib
INCLUDING DEPENDENCY RoochFramework
BUILDING btc_holder_coin
Publish modules to address: rooch1umumkfcuew6fdxunvghrsqlgqlygpdxtqc5sgs478czaaj9nc97syxcqhr(0xe6f9bb271ccbb4969b93622e3803e807c880b4cb06290442be3e05dec8b3c17d)
Execution info:
    status: Executed
    gas used: 1324730
    tx hash: 0x8be44f7176353e60ede93812b97c7cbe4ca9077c468e998438af3a589d375fd3
    state root: 0x716633ec3cbe637af613a93eb5ceba60332b976ec54ddea6b27cf5f14b73b3b2
    event root: 0x1737caca26dd98b5ea30026babd4043fb82f0b51a86a018fab633d43be09697f

New modules:
    0xe6f9bb271ccbb4969b93622e3803e807c880b4cb06290442be3e05dec8b3c17d::holder_coin

Updated modules:
    None
```

## 3.3 并且进行调用

### 获取UTXO

```bash
$ rooch rpc request --method btc_queryUTXOs --params '[{"owner":"0xe6f9bb271ccbb4969b93622e3803e807c880b4cb06290442be3e05dec8b3c17d"}, null, "1", true]'
```

### 质押UTXO

```bash
export UTXO_ID=0x826a5e56581ba5ab84c39976f27cf3578cf524308b4ffc123922dfff507e514da7f9e8d832ae586e704a76ba8726aeaf44c37274f01779a2ac43b1662a6b1930
rooch move run --function default::holder_coin::stake  --args object_id:$UTXO_ID 
```

### 获取奖励

```bash
$ rooch move run --function default::holder_coin::claim --args object_id:$HOLDER_ID --args object_id:$UTXO_ID | jq ".execution_info
```
