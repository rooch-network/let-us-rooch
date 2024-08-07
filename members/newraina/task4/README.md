一个在 mint 时需要消耗自定义 Coin (SnowCoin) 的 NFT 合约。

## 目录

- [1. 获取 mint NFT 所需的 SnowCoin](#1-获取-mint-nft-所需的-snowcoin)
- [2. 创建 Collection](#2-创建-collection)
- [3. mint NFT](#3-mint-nft)
- [4. 查询账号拥有的 NFT 信息](#4-查询账号拥有的-nft-信息)

## 1. 获取 mint NFT 所需的 SnowCoin

```bash
rooch move run --function 0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::coin::faucet
```

执行成功后查询余额。

```bash
curl --location 'https://dev-seed.rooch.network:443' --header 'Content-Type: application/json' --data '{
 "id":101,
 "jsonrpc":"2.0",
 "method":"rooch_getBalance",
 "params":["0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe", "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::coin::SnowCoin"]
}' | jq
```

可以看到目前余额为 10000，成功获取到了 SnowCoin。

```json
{
  "jsonrpc": "2.0",
  "result": {
    "coin_type": "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::coin::SnowCoin",
    "name": "Snow Coin",
    "symbol": "SNW",
    "decimals": 10,
    "supply": "210000000000",
    "balance": "10000"
  },
  "id": 101
}
```

## 2. 创建 Collection

```bash
rooch move run --function 0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::collection::create_collection_entry
```

观察输出中出现了 CreateCollectionEvent，说明创建 Collection 成功。

```json
{
  "event_id": {
    "event_handle_id": "0x1353995413e39b6f04528b645355c944c661a5cce376d310f062fa37eae87920",
    "event_seq": "0"
  },
  "event_type": "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::collection::CreateCollectionEvent",
  "event_data": "0x01175fccfc454a288ea37f54fa00dd79e8bad13f2de5cfeb24d9787bb0ca0ebeef13536e6f77204e465420436f6c6c656374696f6ef6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe0140420f00000000001b5468697320697320536e6f77204e465420436f6c6c656374696f6e",
  "event_index": "2",
  "decoded_event_data": null
}
```

使用 event_id 查询 Collection 的 ObjectId

```bash
curl --location --request POST 'https://dev-seed.rooch.network:443' --header 'Content-Type: application/json' --data-raw '{
"id":101,
"jsonrpc":"2.0",
"method":"rooch_getEventsByEventHandle", "params":["0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::collection::CreateCollectionEvent", null, "1000", false, {"decode":true}]
}' | jq
```

可以得到 Collection ObjectId 为 `0x175fccfc454a288ea37f54fa00dd79e8bad13f2de5cfeb24d9787bb0ca0ebeef`。

<details>
<summary>点击展开详细输出</summary>

```json
{
  "jsonrpc": "2.0",
  "result": {
    "data": [
      {
        "event_id": {
          "event_handle_id": "0x1353995413e39b6f04528b645355c944c661a5cce376d310f062fa37eae87920",
          "event_seq": "0"
        },
        "event_type": "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::collection::CreateCollectionEvent",
        "event_data": "0x01175fccfc454a288ea37f54fa00dd79e8bad13f2de5cfeb24d9787bb0ca0ebeef13536e6f77204e465420436f6c6c656374696f6ef6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe0140420f00000000001b5468697320697320536e6f77204e465420436f6c6c656374696f6e",
        "event_index": "2",
        "decoded_event_data": {
          "abilities": 3,
          "type": "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::collection::CreateCollectionEvent",
          "value": {
            "creator": "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe",
            "description": "This is Snow NFT Collection",
            "maximum": {
              "abilities": 7,
              "type": "0x1::option::Option<u64>",
              "value": {
                "vec": ["1000000"]
              }
            },
            "name": "Snow NFT Collection",
            "object_id": "0x175fccfc454a288ea37f54fa00dd79e8bad13f2de5cfeb24d9787bb0ca0ebeef"
          }
        }
      }
    ],
    "next_cursor": "0",
    "has_next_page": false
  },
  "id": 101
}
```

</details>

## 3. mint NFT

使用第一步得到的 Collection ObjectId mint NFT：

```bash
rooch move run --function default::nft::mint_entry --args object_id:0x175fccfc454a288ea37f54fa00dd79e8bad13f2de5cfeb24d9787bb0ca0ebeef --args String:'Snow NFT 0001'
```

观察输出中 status 为 executed，说明 mint NFT 成功。

```json
"execution_info": {
    "tx_hash": "0x78ad49e4979bd4a4183264dbcff2e6043dd26aa4d6775f848ce60be23935f7ca",
    "state_root": "0x42a1309909a1643f5f246908ac05d5d7084d723596302828faeea3c16a55139a",
    "event_root": "0x64f9a6ff1772dff2746e05e8c1a6e85a877c74b1955ecb16814c4247963daceb",
    "gas_used": "540399",
    "status": {
      "type": "executed"
    }
  }
```

## 4. 查询账号拥有的 NFT 信息

```bash
curl --location --request POST 'https://dev-seed.rooch.network:443' --header 'Content-Type: application/json' --data-raw '{
"id":101,
"jsonrpc":"2.0",
"method":"rooch_getObjectsByOwner", "params":["0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe", "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::nft::NFT", null, "1000", false, {"decode":true}]
}' | jq
```

可以从输出中看到 NFT 的 name 为 Snow NFT 0001，这就是刚刚 mint 的 NFT。

```json
{
  "id": "0xc96e51b7e4de27ee3b7b0c7606dc2bc90bf257cbe176ee28eb23c4e5c916a136",
  "owner": "rooch17mpzg2393tr3rpznxpc4739fw4zh5chkwkr4fr2cxk4602ul80lqey9hv9",
  "owner_bitcoin_address": "bcrt1p9wawj3spyyl79d0py3rmlgn7pjlpqsvdxcgtk9cenaynlvcga5xs9lpvmq",
  "flag": 0,
  "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
  "size": "0",
  "created_at": "1723019418984",
  "updated_at": "1723019418984",
  "object_type": "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::nft::NFT",
  "value": "0x0d536e6f77204e4654203030303101175fccfc454a288ea37f54fa00dd79e8bad13f2de5cfeb24d9787bb0ca0ebeeff6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe",
  "decoded_value": {
    "abilities": 12,
    "type": "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe::nft::NFT",
    "value": {
      "collection": "0x175fccfc454a288ea37f54fa00dd79e8bad13f2de5cfeb24d9787bb0ca0ebeef",
      "creator": "0xf6c2242a258ac711845330715f44a975457a62f67587548d5835aba7ab9f3bfe",
      "name": "Snow NFT 0001"
    }
  },
  "tx_order": "7",
  "state_index": "6",
  "display_fields": null
}
```
