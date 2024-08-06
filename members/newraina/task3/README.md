这是一个测试 UTXO 临时状态存储的合约。在入口函数 main 中，会把当前区块时间和高度存到指定 UTXO 的临时状态中。然后再取出来，比较是否一致。如果一致，则触发 TempStatePassedCheckEvent 事件。

## 1. 搭建本地 Bitcoin 以及 Rooch 开发环境

在 bitcoind_regtest 启动之后，启动 rooch server 会有以下错误 (rooch@0.6.4)：
```bash
2024-07-29T23:57:22.626767Z ERROR rooch_relayer::actor::bitcoin_relayer: BitcoinRelayer sync block error: JSON-RPC error: RPC error response: RpcError { code: -5, message: "Block not found", data: None }
```

这个错误并不影响后面测试 UTXO 合约，可以先忽略。

然后为即将用来测试合约的 btc 地址创建一些交易记录：

```bash
bitcoin-cli generatetoaddress 101 <bitcoin address>
```

## 2. 部署 UTXO 临时状态存储合约

确认 rooch env 在本地测试网之后， 在 `utxokit` 目录下，运行以下命令：

```bash
rooch move publish
```

<details>
<summary>展开执行结果</summary>

```bash
rooch move publish
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
INCLUDING DEPENDENCY BitcoinMove
INCLUDING DEPENDENCY MoveStdlib
INCLUDING DEPENDENCY MoveosStdlib
INCLUDING DEPENDENCY RoochFramework
BUILDING utxokit
Publish modules to address: rooch1p3wyd925p9dpvrukj67adfvcjut9d322qgqg3qrlvcnzvpz46ndszg7dzt(0x0c5c469554095a160f9696bdd6a598971656c54a020088807f6626260455d4db)
Execution info:
    status: Executed
    gas used: 1658749
    tx hash: 0x45cf7481e2dee3162617da56710f60464d5789a39b345bb2ab728350f74b710d
    state root: 0xd54a852cc3fc57ca0a583084461847f9edef8f4be1a02b35e9e085fcad1708d7
    event root: 0x1f0e3e2d969714d5941df4e888c76580f775cf76b3952b84bd0f6065e6eb69a2

New modules:
    0xc5c469554095a160f9696bdd6a598971656c54a020088807f6626260455d4db::temp_state_example

Updated modules:
    None
```
</details>


## 3. 测试 UTXO 临时状态存储合约

使用以下命令获取 UTXO 列表：

```bash
rooch object -t 0x4::utxo::UTXO -o <上面生成交易记录的 bitcoin address>
```

在其中选择一条记录，复制其中的 id (这是 object id)。

```json lines
{
  "id": "0x826a5e56581ba5ab84c39976f27cf3578cf524308b4ffc123922dfff507e514da68584a191f4479a76029d27df994e6fe5d2a0ce7576b6f9deb19add86fbcb90",
  "owner": "rooch1p3wyd925p9dpvrukj67adfvcjut9d322qgqg3qrlvcnzvpz46ndszg7dzt",
  "owner_bitcoin_address": "bcrt1pjjkf5r0g7wn9vcgglnjz6tvuv6x4hc936ask3lwdyr8fxase44mq2u9xss",
  "flag": 0,
  "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
  "size": "0",
  "created_at": "1722178888249",
  "updated_at": "1722178888249",
  "object_type": "0x4::utxo::UTXO",
  "value": "0x65c69e446897e0ae51d720c2e513be7389fc613764ccaa4e251f6b4c8dc4229a0000000000f2052a0100000000",
  "decoded_value": {
    "abilities": 8,
    "type": "0x4::utxo::UTXO",
    "value": {
      "seals": {
        "abilities": 4,
        "type": "0x2::simple_multimap::SimpleMultiMap<0x1::string::String, 0x2::object::ObjectID>",
        "value": {
          "data": []
        }
      },
      "txid": "0x65c69e446897e0ae51d720c2e513be7389fc613764ccaa4e251f6b4c8dc4229a",
      "value": "5000000000",
      "vout": 0
    }
  },
  "tx_order": "109",
  "state_index": "11",
  "display_fields": null
}
```

使用以下命令调用合约：

```bash
rooch move run --function default::temp_state_example::main --args object_id:<上面复制的 id>
```

命令行输出里看到以下内容，表明执行成功：


```json lines
{
  "execution_info": {
    "tx_hash": "0xc34b8e1a4d5e1486bc2fe52c14b525e5012fe10a37907cfcb831a0d40876c07c",
    "state_root": "0xd15191079854d439f623a100ffbbde9c8c96bf0fd2d6de466574a61566c11e31",
    "event_root": "0x619f2e9858b2b7cc805d53d5a7199f3f2dd4167d27a36c2e236cdd17f05ca422",
    "gas_used": "2318910",
    "status": {
      "type": "executed"
    }
  }
}
```

命令行输出里，events 字段下，可以找到如下内容，表明 TempStatePassedCheckEvent 成功触发，说明 UTXO 临时状态存储合约测试成功：

```json lines
[
  {
    "event_id": {
      "event_handle_id": "0x8a1d507f504d4ee5932a9bb5f66b3df97467f28cd37ac7ef9b759a7cbcc5161b",
      "event_seq": "0"
    },
    "event_type": "0xc5c469554095a160f9696bdd6a598971656c54a020088807f6626260455d4db::temp_state_example::TempStatePassedCheckEvent",
    "event_data": "0x522aa866000000006500000000000000",
    "event_index": "2",
    "decoded_event_data": null
  }
]
```
