mononoke@Air rooch_object % rooch move publish   
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
INCLUDING DEPENDENCY MoveStdlib
INCLUDING DEPENDENCY MoveosStdlib
INCLUDING DEPENDENCY RoochFramework
BUILDING rooch_object
Publish modules to address: rooch1gpk8gnlz0d2cgnk2u5a742dzj2vlrgkhzml66q6wcufxec5lflhswajqj0(0x406c744fe27b55844ecae53beaa9a29299f1a2d716ffad034ec7126ce29f4fef)
{
  "sequence_info": {
    "tx_order": "3",
    "tx_order_signature": "0x0167ee3b851fc6fb34696e36300d4235e8b50a85e8b2497917e2dd7e6f6f5681ec381faa87be1de3fcd47a9cb7e5d9988507d27393715c6225e225ddae29e297de026c9e5a00643a706d3826424f766bbbb08adada4dc357c1b279ad4662d2fd1e2e",
    "tx_accumulator_root": "0x7acea7fcf08f7aab86dbd23c4ec4290ba9e57275d7b30af0e960851bae5f62ea",
    "tx_timestamp": "1721324586717"
  },
  "execution_info": {
    "tx_hash": "0xe07a345ee700cc024e27ecabb4b4dad36c12f20e80bd35148e5adb93ecf12e0f",
    "state_root": "0xf092cd155f43f7f707c8a13dc15a443ff0aa5a72d984fb3379797f10c5f49e12",
    "event_root": "0xdcf9f017d4c0eb6359bac67b6998f02b0a281f5e08e42c57060b83c907fdee1e",
    "gas_used": "1695869",
    "status": {
      "type": "executed"
    }
  },
  "output": {
    "status": {
      "type": "executed"
    },
    "changeset": {
      "state_root": "0xf092cd155f43f7f707c8a13dc15a443ff0aa5a72d984fb3379797f10c5f49e12",
      "global_size": "33",
      "changes": [
        {
          "metadata": {
            "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdb",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 1,
            "state_root": "0x10d10905e1391ce27af2b41a7f80afd8888ead7837c319a173362fea89939adf",
            "size": "6",
            "created_at": "0",
            "updated_at": "0",
            "object_type": "0x2::module_store::ModuleStore"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdb406c744fe27b55844ecae53beaa9a29299f1a2d716ffad034ec7126ce29f4fef",
                "owner": "rooch1gpk8gnlz0d2cgnk2u5a742dzj2vlrgkhzml66q6wcufxec5lflhswajqj0",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0xc23eedf9ec86779664891b5bfa87b7d174a060bc3b0af5a6d7a9e2d01c685c74",
                "size": "1",
                "created_at": "1721324586717",
                "updated_at": "1721324586717",
                "object_type": "0x2::module_store::Package"
              },
              "value": {
                "new": "0x00"
              },
              "fields": [
                {
                  "metadata": {
                    "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdb406c744fe27b55844ecae53beaa9a29299f1a2d716ffad034ec7126ce29f4fefa6a1e48eba790756b5a045b0ef2e8ede31b666822a06958654a68d4495d9bde9",
                    "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                    "owner_bitcoin_address": null,
                    "flag": 0,
                    "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                    "size": "0",
                    "created_at": "1721324586717",
                    "updated_at": "1721324586717",
                    "object_type": "0x2::object::DynamicField<0x1::string::String, 0x2::move_module::MoveModule>"
                  },
                  "value": {
                    "new": "0x0c726f6f63685f6f626a656374fd04a11ceb0b060000000a01000e020e1603244104650a056f3b07aa01b60108e0026006c003330af3030d0c80044f000001010102020302040205020600070c0000080300040a0c010001040f0700021307000009000100000b000200000c000100061001030004110501010c040b06070108040e0809010003120601010302140c0b000515000b0002160d0100010c0e010100040405040604070a0b0b010300010b020108000105010800020b0201090005010900010b0201090001060b02010900010803010801010804010a02020708040804010609000c726f6f63685f6f626a65637405646562756706737472696e67056576656e74066f626a6563740c737472696e675f7574696c730a74785f636f6e7465787409436f6f6b6965426f78084e65774576656e740a6372656174655f626f78064f626a656374036e6577057072696e740576616c7565026964084f626a65637449440673656e646572087472616e7366657204656d697406537472696e6704757466380d746f5f737472696e675f75363406617070656e64406c744fe27b55844ecae53beaa9a29299f1a2d716ffad034ec7126ce29f4fef000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020a021d1c43726561746564206e657720636f6f6b696520626f782077697468200a02100f20636f6f6b69657320696e2069742e0002010d030102020e08030d030001040001070a001101110338000b0011020201010000020b0a00120038010c010e0138020b00120138030b0102020100000b0e070011080c010d010b001109110a0d0107011108110a0e0138040200"
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
            "created_at": "1721324586717",
            "updated_at": "1721324586717",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0x7de019000000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0x406c744fe27b55844ecae53beaa9a29299f1a2d716ffad034ec7126ce29f4fef",
            "owner": "rooch1gpk8gnlz0d2cgnk2u5a742dzj2vlrgkhzml66q6wcufxec5lflhswajqj0",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721324586717",
            "updated_at": "1721324586717",
            "object_type": "0x2::account::Account"
          },
          "value": {
            "new": "0x406c744fe27b55844ecae53beaa9a29299f1a2d716ffad034ec7126ce29f4fef0100000000000000"
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
            "updated_at": "1721324586717",
            "object_type": "0x2::timestamp::Timestamp"
          },
          "value": {
            "modify": "0xddc6f0c690010000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpsd68l8x",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5084e1d481544f494c209d1c387fa92aeaa3d8776aaed32298b33b09f2cc7a90",
            "size": "3",
            "created_at": "0",
            "updated_at": "0",
            "object_type": "0x3::address_mapping::RoochToBitcoinAddressMapping"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a06952ae931f8680fc811362aaf8c118dbfc8f5e1e76490303ef35c9b88f2979a",
                "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                "size": "0",
                "created_at": "1721324586717",
                "updated_at": "1721324586717",
                "object_type": "0x2::object::DynamicField<address, 0x3::bitcoin_address::BitcoinAddress>"
              },
              "value": {
                "new": "0x406c744fe27b55844ecae53beaa9a29299f1a2d716ffad034ec7126ce29f4fef2202016387ad010469ffebec8eda8aa9e4b04c06b3f326abbec9669cc5d33c782700a7"
              },
              "fields": []
            }
          ]
        },
        {
          "metadata": {
            "id": "0xe88c6b6be18c427c61b6c93ea3e8e16dddc19b073310fe173055357f31300511",
            "owner": "rooch1gpk8gnlz0d2cgnk2u5a742dzj2vlrgkhzml66q6wcufxec5lflhswajqj0",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721324586717",
            "updated_at": "1721324586717",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "new": "0x8300dc050000000000000000000000000000000000000000000000000000000000"
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
            "created_at": "1721324586717",
            "updated_at": "1721324586717",
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
        "event_data": "0x01e88c6b6be18c427c61b6c93ea3e8e16dddc19b073310fe173055357f3130051153303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e",
        "event_index": "1",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": "1"
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x01e88c6b6be18c427c61b6c93ea3e8e16dddc19b073310fe173055357f3130051153303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e00e1f50500000000000000000000000000000000000000000000000000000000",
        "event_index": "2",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x8e3089f2c059cc5377a1b6b7c3dcefba8a586697c35de27c2a4b68f81defb69c",
          "event_seq": "0"
        },
        "event_type": "0x3::coin_store::WithdrawEvent",
        "event_data": "0x01e88c6b6be18c427c61b6c93ea3e8e16dddc19b073310fe173055357f3130051153303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e7de0190000000000000000000000000000000000000000000000000000000000",
        "event_index": "3",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": "2"
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x0136b66e328827e3f63b94bf4596902e99a5dff17308336904357ed247fd194be453303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e7de0190000000000000000000000000000000000000000000000000000000000",
        "event_index": "4",
        "decoded_event_data": null
      }
    ],
    "gas_used": "1695869",
    "is_upgrade": false
  }
}