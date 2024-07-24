# 学习日志

## Task2

### 安装 Rooch

```bash
$ wget https://github.com/rooch-network/rooch/releases/latest/download/rooch-ubuntu-latest.zip
$ unzip rooch-ubuntu-latest.zip
$ cd rooch-artifacts && sudo cp rooch /usr/local/bin

$ rooch -V
rooch 0.6.1
```

### 查看命令行帮助

```bash
$ rooch -h
Usage: rooch <COMMAND>

Commands:
  account      Tool for interacting with accounts
  init         Tool for init with rooch
  move
  server       Start Rooch network
  state        Get states by accessPath
  object
  resource     Get account resource by tag
  transaction  Tool for interacting with transaction
  event        Tool for interacting with event
  abi
  env          Interface for managing multiple environments
  session-key  Session key Commands
  rpc
  statedb      Statedb Commands
  indexer      Indexer Commands
  genesis      Statedb Commands
  help         Print this message or the help of the given subcommand(s)

Options:
  -h, --help     Print help
  -V, --version  Print version
```

### 账号初始化

```bash
$ rooch init
Creating client config file ["/root/.rooch/rooch_config/rooch.yaml"].
Enter a password to encrypt the keys. Press enter to leave it an empty password:
Generated new keypair for address [rooch1u9wflnw3mp34n45rpkkdgda7nk6tfmg9077eu43970u60h8zfwwq3whypw]
Secret Recovery Phrase : [auction clap monitor february oblige impact blush hurdle fire squeeze zebra soft]
Rooch client config file generated at /root/.rooch/rooch_config/rooch.yaml
```

### 查看账号列表

```bash
$ rooch account list --json
{
  "default": {
    "address": "rooch1u9wflnw3mp34n45rpkkdgda7nk6tfmg9077eu43970u60h8zfwwq3whypw",
    "hex_address": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c",
    "bitcoin_address": "bcrt1p5d3muf793ccy22ujfss5mh0kfc5k8zyxf7hv5dp5vcpmhxprtzlsqr2gcj",
    "nostr_public_key": "npub1j3daem4crewzz4nfauc423ryqd98chtkqh58c36r6uuhhu9yd5aqwpmdz3",
    "public_key": "AQOUW9zuuB5cIVZp7zFVRGQDSnxddgXofEdD1zl78KRtOg==",
    "has_session_key": false,
    "active": true
  }
}
```

### 启动本地网络

```bash
$ rooch server start

INFO rooch_rpc_server: JSON-RPC HTTP Server start listening 0.0.0.0:6767
INFO rooch_rpc_server: Available JSON-RPC methods : ["rooch_getObjectStates", "rooch_queryObjectStates", "rooch_getBalance", "rooch_sendRawTransaction", "rooch_getTransactionsByHash", "rooch_queryEvents", "rooch_executeRawTransaction", "rooch_getEventsByEventHandle", "rooch_queryTransactions", "rooch_executeViewFunction", "rooch_listFieldStates", "rooch_getBalances", "rooch_getModuleABI", "rooch_listStates", "rooch_getTransactionsByOrder", "btc_queryInscriptions", "rooch_getChainID", "rooch_getFieldStates", "btc_queryUTXOs", "rpc.discover", "rooch_getStates"]
The active env was successfully switched to `local`
```

### 确认当前网络为本地网络

```bash
$ rooch env list
       Env Alias    |              RPC URL                |    Websocket URL   |  Active Env
---------------------------------------------------------------------------------------------
         local      |        http://0.0.0.0:6767          |         Null       |     True
          dev       |   https://dev-seed.rooch.network    |         Null       |
          test      |  https://test-seed.rooch.network    |         Null       |
```

### 创建合约

```
$ rooch move new counter
Success
```

### 实现合约

```bash
module counter::counter {

   use moveos_std::account;

   struct Counter has key {
      value:u64,
   }

   fun init() {
      let signer = moveos_std::signer::module_signer<Counter>();
      account::move_resource_to(&signer, Counter { value: 0 });
   }

   public fun increase_() {
      let counter = account::borrow_mut_resource<Counter>(@counter);
      counter.value = counter.value + 1;
   }

   public entry fun increase() {
      Self::increase_()
   }

   public fun value(): u64 {
      let counter = account::borrow_resource<Counter>(@counter);
      counter.value
   }
}
```

### 编译部署合约

```bash
$ rooch move publish
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
INCLUDING DEPENDENCY MoveStdlib
INCLUDING DEPENDENCY MoveosStdlib
INCLUDING DEPENDENCY RoochFramework
BUILDING counter
Publish modules to address: rooch1u9wflnw3mp34n45rpkkdgda7nk6tfmg9077eu43970u60h8zfwwq3whypw(0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c)
{
  "sequence_info": {
    "tx_order": "1",
    "tx_order_signature": "0x011ffecb89a5a94422e95d04855032a7610b16c65d513e4db53e687c301e1aabab643ef40379299471d33684fca1cfbdf7ec746fbc751e93f81b69453bf4bba94703945bdceeb81e5c215669ef31554464034a7c5d7605e87c4743d7397bf0a46d3a",
    "tx_accumulator_root": "0x1104758772e77ab3ac172044987cff39a8bce8c4f5cf57d1d3ca196acca78fa8",
    "tx_timestamp": "1721706285132"
  },
  "execution_info": {
    "tx_hash": "0xff3300adbb1e3d5faedd34241e5bad22c3f049b6b46b50bef4cc41a4a3683654",
    "state_root": "0xdaf90e18b980a2ddd803ed18f3af94c393ab415bf8e9ef6c40999af7e3f8e13f",
    "event_root": "0x226d21ba8efadc1b315f56aeb8e2e1f748e9931258189776f0332837a98fb53c",
    "gas_used": "2441316",
    "status": {
      "type": "executed"
    }
  },
  "output": {
    "status": {
      "type": "executed"
    },
    "changeset": {
      "state_root": "0xdaf90e18b980a2ddd803ed18f3af94c393ab415bf8e9ef6c40999af7e3f8e13f",
      "global_size": "33",
      "changes": [
        {
          "metadata": {
            "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdb",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 1,
            "state_root": "0x73ff467ea867d1c5c19c0f48329099a31ea6460b96a214ed7b9faee1cffbc03b",
            "size": "6",
            "created_at": "0",
            "updated_at": "0",
            "object_type": "0x2::module_store::ModuleStore"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdbe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c",
                "owner": "rooch1u9wflnw3mp34n45rpkkdgda7nk6tfmg9077eu43970u60h8zfwwq3whypw",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0x592cb5c80cb8039c8f0d39e866b82e36c90857bae35bf7b914f318546b5a1c46",
                "size": "1",
                "created_at": "1721706285132",
                "updated_at": "1721706285132",
                "object_type": "0x2::module_store::Package"
              },
              "value": {
                "new": "0x00"
              },
              "fields": [
                {
                  "metadata": {
                    "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdbe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c07d5213b89daff24b7cb8b6a39c220bae391dcb598ced039e9b6ecb0b4640758",
                    "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                    "owner_bitcoin_address": null,
                    "flag": 0,
                    "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                    "size": "0",
                    "created_at": "1721706285132",
                    "updated_at": "1721706285132",
                    "object_type": "0x2::object::DynamicField<0x1::string::String, 0x2::move_module::MoveModule>"
                  },
                  "value": {
                    "new": "0x07636f756e746572c503a11ceb0b060000000b010006020604030a2c043608053e1b0759800108d90140069902220abb02050cc002520d920302000001010102000308000004000000000500000000060000000007000100010804050108020900060100010a07000108010b04080108040305030603070300010301070800010800010501070900010c02060c09000106090007636f756e746572076163636f756e74067369676e657207436f756e74657208696e63726561736509696e6372656173655f04696e69740576616c756513626f72726f775f6d75745f7265736f757263650d6d6f64756c655f7369676e6572106d6f76655f7265736f757263655f746f0f626f72726f775f7265736f75726365e15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c00000000000000000000000000000000000000000000000000000000000000020520e15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c000201070300010400000211010201010000020c070038000c000a00100014060100000000000000160b000f00150202000000060738010c000e0006000000000000000012003802020301000000050700380310001402000000"
                  },
                  "fields": []
                }
              ]
            }
          ]
        },
        {
          "metadata": {
            "id": "0x4e8d2c243339c6e02f8b7dd34436a1b1eb541b0fe4d938f845f4dbb9d9f218a2",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 1,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721706285132",
            "updated_at": "1721706285132",
            "object_type": "0x2::timestamp::Timestamp"
          },
          "value": {
            "modify": "0x4c08b1dd90010000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpsd68l8x",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x83606f610d94411fd16a3811a4adc2739dfaca6f5e1cdd7a6ac6456980a0a85f",
            "size": "2",
            "created_at": "0",
            "updated_at": "0",
            "object_type": "0x3::address_mapping::RoochToBitcoinAddressMapping"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a55ff8a0b6b5feacc7a7226ca8171a3e55a6ac564b7c1704711f51b3e084373cf",
                "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                "size": "0",
                "created_at": "1721706285132",
                "updated_at": "1721706285132",
                "object_type": "0x2::object::DynamicField<address, 0x3::bitcoin_address::BitcoinAddress>"
              },
              "value": {
                "new": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c220201a363be27c58e30452b924c214dddf64e296388864faeca34346603bb982358bf"
              },
              "fields": []
            }
          ]
        },
        {
          "metadata": {
            "id": "0xb53bc58fed718fb32633c9ec566888543d8db3ce1a4908e2b8006adf8a896f91",
            "owner": "rooch1u9wflnw3mp34n45rpkkdgda7nk6tfmg9077eu43970u60h8zfwwq3whypw",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721706285132",
            "updated_at": "1721706285132",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "new": "0x9ca0d0050000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xc777c90e88125df5e24ba3f745dbd1f38fc26c600a4764b3263859d00d443aa8",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721706285132",
            "updated_at": "1721706285132",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0x644025000000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c",
            "owner": "rooch1u9wflnw3mp34n45rpkkdgda7nk6tfmg9077eu43970u60h8zfwwq3whypw",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0xa2e29a6ae93b85c576b4dcf2245a038af03222d7020f7a510b8008affdf768d8",
            "size": "1",
            "created_at": "1721706285132",
            "updated_at": "1721706285132",
            "object_type": "0x2::account::Account"
          },
          "value": {
            "new": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c0100000000000000"
          },
          "fields": [
            {
              "metadata": {
                "id": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c38129d4da977ab579bd1cfc36208def2a45fd089e251cda61f80ebfe8ddc7aa3",
                "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                "size": "0",
                "created_at": "1721706285132",
                "updated_at": "1721706285132",
                "object_type": "0x2::object::DynamicField<0x1::string::String, 0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c::counter::Counter>"
              },
              "value": {
                "new": "0x52653135633966636464316438363335396436383330646163643433376265396462346234656430353766626439653536323566336639613764636532346239633a3a636f756e7465723a3a436f756e7465720000000000000000"
              },
              "fields": []
            }
          ]
        },
        {
          "metadata": {
            "id": "0xf81628c3bf85c3fc628f29a3739365d4428101fbbecca0dcc7e3851f34faea6b",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpsd68l8x",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721706285132",
            "updated_at": "1721706285132",
            "object_type": "0x3::coin::CoinInfo<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0x53303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e0e526f6f63682047617320436f696e035247430800217016f35a0000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        }
      ]
    },
    "events": [
      {
        "event_id": {
          "event_handle_id": "0x358779b791ef606d7f07df8881c1939f26de95119486b60745a9c3127ae8fd37",
          "event_seq": "1"
        },
        "event_type": "0x3::coin::MintEvent",
        "event_data": "0x53303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e00e1f50500000000000000000000000000000000000000000000000000000000",
        "event_index": "0",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0xdebc7ccc8fa8855fad9fdd2919e875e06bcfa9b11cdc53c3247e0f81239852e2",
          "event_seq": "2"
        },
        "event_type": "0x3::coin_store::CreateEvent",
        "event_data": "0x01b53bc58fed718fb32633c9ec566888543d8db3ce1a4908e2b8006adf8a896f9153303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e",
        "event_index": "1",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": "1"
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x01b53bc58fed718fb32633c9ec566888543d8db3ce1a4908e2b8006adf8a896f9153303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e00e1f50500000000000000000000000000000000000000000000000000000000",
        "event_index": "2",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x8e3089f2c059cc5377a1b6b7c3dcefba8a586697c35de27c2a4b68f81defb69c",
          "event_seq": "0"
        },
        "event_type": "0x3::coin_store::WithdrawEvent",
        "event_data": "0x01b53bc58fed718fb32633c9ec566888543d8db3ce1a4908e2b8006adf8a896f9153303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e6440250000000000000000000000000000000000000000000000000000000000",
        "event_index": "3",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": "2"
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x01c777c90e88125df5e24ba3f745dbd1f38fc26c600a4764b3263859d00d443aa853303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e6440250000000000000000000000000000000000000000000000000000000000",
        "event_index": "4",
        "decoded_event_data": null
      }
    ],
    "gas_used": "2441316",
    "is_upgrade": false
  }
}
```

### 记录环境变量

```bash
export PACKAGE_ID=0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c
export JASON=0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c
```

### 调用合约方法

```bash
rooch move run --function $PACKAGE_ID::counter::increase --sender-account $JASON
{
  "sequence_info": {
    "tx_order": "2",
    "tx_order_signature": "0x01a13c4d270b2e238723e6ca7e157537cee3cf254973f13f1164437b3bd02a3dab15c6f6c31e0ef21e64910717b55ae26080b1841d63650027dcacd532a721356803945bdceeb81e5c215669ef31554464034a7c5d7605e87c4743d7397bf0a46d3a",
    "tx_accumulator_root": "0x4e580fe49efcf74b842029e1cd7431d85413d3d369bda572d60340178d7fb9da",
    "tx_timestamp": "1721706678682"
  },
  "execution_info": {
    "tx_hash": "0x0bddd6c4e87f7628959e178f2306f7bd9897745c117e20f230ab6f4b3395cd1b",
    "state_root": "0xfafef7be077f252eb079e60a92ecabf4a5afa3be0a35f5922496cb0478e2d527",
    "event_root": "0x74362f2ed645a9389958bbcd891e3fafa64d5683c4667e56e2fb632a573fc1f4",
    "gas_used": "558759",
    "status": {
      "type": "executed"
    }
  },
  "output": {
    "status": {
      "type": "executed"
    },
    "changeset": {
      "state_root": "0xfafef7be077f252eb079e60a92ecabf4a5afa3be0a35f5922496cb0478e2d527",
      "global_size": "33",
      "changes": [
        {
          "metadata": {
            "id": "0x4e8d2c243339c6e02f8b7dd34436a1b1eb541b0fe4d938f845f4dbb9d9f218a2",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 1,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721706285132",
            "updated_at": "1721706678682",
            "object_type": "0x2::timestamp::Timestamp"
          },
          "value": {
            "modify": "0x9a09b7dd90010000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xb53bc58fed718fb32633c9ec566888543d8db3ce1a4908e2b8006adf8a896f91",
            "owner": "rooch1u9wflnw3mp34n45rpkkdgda7nk6tfmg9077eu43970u60h8zfwwq3whypw",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721706285132",
            "updated_at": "1721706678682",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0xf519c8050000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xc777c90e88125df5e24ba3f745dbd1f38fc26c600a4764b3263859d00d443aa8",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721706285132",
            "updated_at": "1721706678682",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0x0bc72d000000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c",
            "owner": "rooch1u9wflnw3mp34n45rpkkdgda7nk6tfmg9077eu43970u60h8zfwwq3whypw",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x3863fc852c59965f98734dcdbe5f6d1857b88ee7ab1dbf5f7dfeb09851ee972c",
            "size": "1",
            "created_at": "1721706285132",
            "updated_at": "1721706678682",
            "object_type": "0x2::account::Account"
          },
          "value": {
            "modify": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c0200000000000000"
          },
          "fields": [
            {
              "metadata": {
                "id": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c38129d4da977ab579bd1cfc36208def2a45fd089e251cda61f80ebfe8ddc7aa3",
                "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                "size": "0",
                "created_at": "1721706285132",
                "updated_at": "1721706678682",
                "object_type": "0x2::object::DynamicField<0x1::string::String, 0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c::counter::Counter>"
              },
              "value": {
                "modify": "0x52653135633966636464316438363335396436383330646163643433376265396462346234656430353766626439653536323566336639613764636532346239633a3a636f756e7465723a3a436f756e7465720100000000000000"
              },
              "fields": []
            }
          ]
        }
      ]
    },
    "events": [
      {
        "event_id": {
          "event_handle_id": "0x8e3089f2c059cc5377a1b6b7c3dcefba8a586697c35de27c2a4b68f81defb69c",
          "event_seq": "1"
        },
        "event_type": "0x3::coin_store::WithdrawEvent",
        "event_data": "0x01b53bc58fed718fb32633c9ec566888543d8db3ce1a4908e2b8006adf8a896f9153303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696ea786080000000000000000000000000000000000000000000000000000000000",
        "event_index": "0",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": "3"
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x01c777c90e88125df5e24ba3f745dbd1f38fc26c600a4764b3263859d00d443aa853303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696ea786080000000000000000000000000000000000000000000000000000000000",
        "event_index": "1",
        "decoded_event_data": null
      }
    ],
    "gas_used": "558759",
    "is_upgrade": false
  }
}
```

### 查看结果

```bash
$ rooch resource --address $JASON --resource $PACKAGE_ID::counter::Counter
{
  "id": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c38129d4da977ab579bd1cfc36208def2a45fd089e251cda61f80ebfe8ddc7aa3",
  "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
  "owner_bitcoin_address": null,
  "flag": 0,
  "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
  "size": "0",
  "created_at": "1721706285132",
  "updated_at": "1721706678682",
  "object_type": "0x2::object::DynamicField<0x1::string::String, 0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c::counter::Counter>",
  "value": "0x52653135633966636464316438363335396436383330646163643433376265396462346234656430353766626439653536323566336639613764636532346239633a3a636f756e7465723a3a436f756e7465720100000000000000",
  "decoded_value": {
    "abilities": 12,
    "type": "0x2::object::DynamicField<0x1::string::String, 0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c::counter::Counter>",
    "value": {
      "name": "e15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c::counter::Counter",
      "value": {
        "abilities": 8,
        "type": "0xe15c9fcdd1d86359d6830dacd437be9db4b4ed057fbd9e5625f3f9a7dce24b9c::counter::Counter",
        "value": {
          "value": "1"
        }
      }
    }
  },
  "display_fields": null
}
```
