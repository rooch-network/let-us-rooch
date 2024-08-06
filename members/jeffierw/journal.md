# 学习日志

## task1
参考了![rooch](https://github.com/rooch-network/rooch/tree/main/scripts/bitcoin)老师提供的文档，但是不知为何运行的是regtest测试网络，后来尝试用
`brew install bitcoin`安装，并使用`bitcoind -daemon -txindex -chain=main`本地运行了bitcoin全节点，![fullnode](./task1/bitcoin_fullnode.jpg)

## task2
参考了[rooch](https://rooch.network/)文档，一步步来没遇到太多问题。

1. 先安装rooch（mac）

通过[github releases](https://github.com/rooch-network/rooch/releases)安装二进制文件

然后把文件放入本地bin文件夹即可

2. 获取账户和gas
```shell
rooch init
```
会创建一个rooch账户和助记词

通过[discord](https://discord.gg/rooch)获取gas

3. 创建合约
```shell
rooch move new hello_rooch
```

4. 编写代码

在 sources 目录里创建一个 counter.move 文件来编写我们的合约代码。
```move
module quick_start_counter::quick_start_counter {
    use moveos_std::account;
    struct Counter has key {
        count_value: u64
    }
    fun init() {
        let signer = moveos_std::signer::module_signer<Counter>();
        account::move_resource_to(&signer, Counter { count_value: 0 });
    }
    entry fun increase() {
        let counter = account::borrow_mut_resource<Counter>(@quick_start_counter);
        counter.count_value = counter.count_value + 1;
    }
}
```
类似sui move的结构，引用了`MoveOS`标准库中的`account`模块；

再定义一个包含count数据的`Counter`结构体，且具有key能力；

接下来是`init`函数，通过`moveos_std`的函数对`Counter`中的值初始化；

最后通过`entry`定义入口函数，内部逻辑就是对定义的`count_value`数据+1，使用到了`account`的`borrow_mut_resource`函数。

5. 编译部署
先确定环境是dev环境 如果不是可以切换一下
```shell
// 查看env
rooch env list
// 切换env
rooch env switch --alias dev
// 编译
rooch move build
// 部署
rooch move publish
```

6. 输出日志
```shell
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
INCLUDING DEPENDENCY MoveStdlib
INCLUDING DEPENDENCY MoveosStdlib
INCLUDING DEPENDENCY RoochFramework
BUILDING hello_rooch
Publish modules to address: rooch15xqe7ku3l4gaz5rhkyjhgfzdf9m4l05q9xl3svfvy3q0teelpecqctxxws(0xa1819f5b91fd51d15077b12574244d49775fbe8029bf18312c2440f5e73f0e70)
{
  "sequence_info": {
    "tx_order": "3",
    "tx_order_signature": "0x011e403b0d4d58a6d3b93fc8cb347a12de574d4a464d1b9e60776cb4147243bef62806602fa14082ae79d8613d08fe09bc9a305fea354d47166730c9fb436f9d81026c9e5a00643a706d3826424f766bbbb08adada4dc357c1b279ad4662d2fd1e2e",
    "tx_accumulator_root": "0x57c9bcf1f59fb6919208487ce9985d8df16cdfd9cb2795284b06a0d26cbe8b2e",
    "tx_timestamp": "1721223836435"
  },
  "execution_info": {
    "tx_hash": "0xf341446678f21d3bb63967fa09824a17e8b17ae362c417e06b18e2c068880502",
    "state_root": "0xa641978106c9dd2aee37b08331ca894ccf30988b8316384403e0723c64d33395",
    "event_root": "0x6fabd1d96e25db32e135804e297bcf914921cf48012db138feb6905e7133b0c4",
    "gas_used": "1673587",
    "status": {
      "type": "executed"
    }
  },
  "output": {
    "status": {
      "type": "executed"
    },
    "changeset": {
      "state_root": "0xa641978106c9dd2aee37b08331ca894ccf30988b8316384403e0723c64d33395",
      "global_size": "33",
      "changes": [
        {
          "metadata": {
            "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdb",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 1,
            "state_root": "0xce520d71116016b422f2e4c576343da7fadc2da63a66b4cbaef4c087ba28e65a",
            "size": "6",
            "created_at": "0",
            "updated_at": "0",
            "object_type": "0x2::module_store::ModuleStore"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdba1819f5b91fd51d15077b12574244d49775fbe8029bf18312c2440f5e73f0e70",
                "owner": "rooch15xqe7ku3l4gaz5rhkyjhgfzdf9m4l05q9xl3svfvy3q0teelpecqctxxws",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0xd0f48e346aa4837d2a31b679bc10042f28f0c01fc54f0f6ec3c8cc5d0de739c1",
                "size": "1",
                "created_at": "1721223836435",
                "updated_at": "1721223836435",
                "object_type": "0x2::module_store::Package"
              },
              "value": {
                "new": "0x00"
              },
              "fields": [
                {
                  "metadata": {
                    "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdba1819f5b91fd51d15077b12574244d49775fbe8029bf18312c2440f5e73f0e70051caf5bd0cc91315ebef20d397f2e114d72ce517ead479d697a175f8489b1c0",
                    "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                    "owner_bitcoin_address": null,
                    "flag": 0,
                    "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                    "size": "0",
                    "created_at": "1721223836435",
                    "updated_at": "1721223836435",
                    "object_type": "0x2::object::DynamicField<0x1::string::String, 0x2::move_module::MoveModule>"
                  },
                  "value": {
                    "new": "0x0b68656c6c6f5f726f6f6368bd02a11ceb0b060000000a010006020608030e10041e020520120732540886016006e601100af601060cfc0115000001010202000308000106070000040001000107030400020805010108020201060c00010800010a0201080102060c09000b68656c6c6f5f726f6f636806737472696e67076163636f756e740c48656c6c6f4d657373616765097361795f68656c6c6f047465787406537472696e670475746638106d6f76655f7265736f757263655f746fa1819f5b91fd51d15077b12574244d49775fbe8029bf18312c2440f5e73f0e70000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020a020d0c48656c6c6f20526f6f6368210002010508010000040002080700110112000c010b000b0138000200"
                  },
                  "fields": []
                }
              ]
            }
          ]
        },
        {
          "metadata": {
            "id": "0x36b66e328827e3f63b94bf4596902e99a5dff17308336904357ed247fd194be4",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721223836435",
            "updated_at": "1721223836435",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0x738919000000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0x4e8d2c243339c6e02f8b7dd34436a1b1eb541b0fe4d938f845f4dbb9d9f218a2",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 1,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1296688602000",
            "updated_at": "1721223836435",
            "object_type": "0x2::timestamp::Timestamp"
          },
          "value": {
            "modify": "0x1373efc090010000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpsd68l8x",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0xe11fa3ed75dd1086c04de626e4beb12d1418490a29a6b7db29c33ea695ca75ab",
            "size": "3",
            "created_at": "0",
            "updated_at": "0",
            "object_type": "0x3::address_mapping::RoochToBitcoinAddressMapping"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0ac38adffcb14f5cef63f37eedc9cb1dbf3c1a5f0e56f4b2433490d73f43fdf465",
                "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                "size": "0",
                "created_at": "1721223836435",
                "updated_at": "1721223836435",
                "object_type": "0x2::object::DynamicField<address, 0x3::bitcoin_address::BitcoinAddress>"
              },
              "value": {
                "new": "0xa1819f5b91fd51d15077b12574244d49775fbe8029bf18312c2440f5e73f0e702202013e5cbe0d1000d33c5fcfbca57070aa1bfab8f625e1d99a49d50755249579a3d2"
              },
              "fields": []
            }
          ]
        },
        {
          "metadata": {
            "id": "0xa1819f5b91fd51d15077b12574244d49775fbe8029bf18312c2440f5e73f0e70",
            "owner": "rooch15xqe7ku3l4gaz5rhkyjhgfzdf9m4l05q9xl3svfvy3q0teelpecqctxxws",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721223836435",
            "updated_at": "1721223836435",
            "object_type": "0x2::account::Account"
          },
          "value": {
            "new": "0xa1819f5b91fd51d15077b12574244d49775fbe8029bf18312c2440f5e73f0e700100000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xea71131b98bf6b6109fc67702a1a97e16116ef2151834421106e07bb65395b3d",
            "owner": "rooch15xqe7ku3l4gaz5rhkyjhgfzdf9m4l05q9xl3svfvy3q0teelpecqctxxws",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721223836435",
            "updated_at": "1721223836435",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "new": "0x8d57dc050000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xf81628c3bf85c3fc628f29a3739365d4428101fbbecca0dcc7e3851f34faea6b",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpsd68l8x",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721223836435",
            "updated_at": "1721223836435",
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
        "event_data": "0x01ea71131b98bf6b6109fc67702a1a97e16116ef2151834421106e07bb65395b3d53303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e",
        "event_index": "1",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": "1"
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x01ea71131b98bf6b6109fc67702a1a97e16116ef2151834421106e07bb65395b3d53303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e00e1f50500000000000000000000000000000000000000000000000000000000",
        "event_index": "2",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x8e3089f2c059cc5377a1b6b7c3dcefba8a586697c35de27c2a4b68f81defb69c",
          "event_seq": "0"
        },
        "event_type": "0x3::coin_store::WithdrawEvent",
        "event_data": "0x01ea71131b98bf6b6109fc67702a1a97e16116ef2151834421106e07bb65395b3d53303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e7389190000000000000000000000000000000000000000000000000000000000",
        "event_index": "3",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": "2"
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x0136b66e328827e3f63b94bf4596902e99a5dff17308336904357ed247fd194be453303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e7389190000000000000000000000000000000000000000000000000000000000",
        "event_index": "4",
        "decoded_event_data": null
      }
    ],
    "gas_used": "1673587",
    "is_upgrade": false
  }
}
```

## Week 3
### 任务：构思一个 Bitcoin 生态的应用或者游戏，可以利用 Rooch 提供的特性，写成文章或者 Github Issue
[Bitcoin Block Derby](./task3/block_derby.md)
