一个保存用户的头像的合约，用户可以设置自己地址对应头像的地址。

合约 init 时会为发起交易的地址默认设置一个头像，头像的 url 为 `https://example.com/avatar.png`。

## 1. 部署

合约部署地址： `0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f`

<details>
<summary>点击查看详细结果</summary>

```shell
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
INCLUDING DEPENDENCY MoveStdlib
INCLUDING DEPENDENCY MoveosStdlib
INCLUDING DEPENDENCY RoochFramework
BUILDING avatar
Publish modules to address: rooch10whhuddqjsqsqccq3uqsnk0rxac3ypq9vk6rr3a6mumyxy0el4lsgg4u9w(0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f)
{
  "sequence_info": {
    "tx_order": "3570341",
    "tx_order_signature": "0x01f1825a7c3e18bc9986d128b5d004132bc231550833a79620e5ca3cb36f77a3c207395ee29a5e90adb09bd096364d4beb46093bff8d08ad2f2d961b8bf402f8d8026c9e5a00643a706d3826424f766bbbb08adada4dc357c1b279ad4662d2fd1e2e",
    "tx_accumulator_root": "0x1ebc927f3156f9b48a294453f956e2b02c6f4725bb56f051ad1967777b325cfc",
    "tx_timestamp": "1721635671353"
  },
  "execution_info": {
    "tx_hash": "0x2492afb2a6f9fe595a555becff341a0b7b70ee50b74ea70db72fce60216ec4a4",
    "state_root": "0x6c2baf29478e19ad97288379af7d21704a316446d4d04d219db8dcbaaaa1c6bf",
    "event_root": "0x44185e499113228ee9bf07a3c0f01cab3e8a6d77543f636f2bd32a28050f3ea7",
    "gas_used": "2094452",
    "status": {
      "type": "executed"
    }
  },
  "output": {
    "status": {
      "type": "executed"
    },
    "changeset": {
      "state_root": "0x6c2baf29478e19ad97288379af7d21704a316446d4d04d219db8dcbaaaa1c6bf",
      "global_size": "446",
      "changes": [
        {
          "metadata": {
            "id": "0x1bfd3970bdc6d8e9ad089741d35b2339933c596707a365f1719aa999fa0cede3",
            "owner": "rooch10whhuddqjsqsqccq3uqsnk0rxac3ypq9vk6rr3a6mumyxy0el4lsgg4u9w",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721635659443",
            "updated_at": "1721635671353",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0x8cd47a3b0000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdb",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 1,
            "state_root": "0x7e983fc1d81e92ad5276524a4c81b9f819f1e19d702984b407d3b81727e1e6db",
            "size": "18",
            "created_at": "1720876964000",
            "updated_at": "1720876964000",
            "object_type": "0x2::module_store::ModuleStore"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdb7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f",
                "owner": "rooch10whhuddqjsqsqccq3uqsnk0rxac3ypq9vk6rr3a6mumyxy0el4lsgg4u9w",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0xfa4c7e1ec7358815b3c5348732d9e5d5fe8e3d0950acbe1299d66ae354066a39",
                "size": "1",
                "created_at": "1721635671353",
                "updated_at": "1721635671353",
                "object_type": "0x2::module_store::Package"
              },
              "value": {
                "new": "0x00"
              },
              "fields": [
                {
                  "metadata": {
                    "id": "0x2214495c6abca5dd5a2bf0f2a28a74541ff10c89818a1244af24c4874325ebdb7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7fae917075d70b875a436137db7d45977223ef5882081ce2ce2aba92efa8fdb25e",
                    "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                    "owner_bitcoin_address": null,
                    "flag": 0,
                    "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                    "size": "0",
                    "created_at": "1721635671353",
                    "updated_at": "1721635671353",
                    "object_type": "0x2::object::DynamicField<0x1::string::String, 0x2::move_module::MoveModule>"
                  },
                  "value": {
                    "new": "0x06617661746172aa04a11ceb0b060000000a01000a020a1603203104510805593a079301970108aa0260068a03220aac030c0cb803440000010101020203020400050e00000607000407070002080700040e0c0100010009000100000a020300000b020300040f060701080110020800040d0901010003110603010304120b03010c02130c0d0003050505060a070502060c080301080201060c00030b04010800080205010800010900010b04010900010501060b04010900010801020b0401090005010a0201080306617661746172067369676e657206737472696e67056576656e74066f626a6563740641766174617212417661746172437265617465644576656e74084f626a656374494406537472696e67066372656174650e6372656174655f64656661756c7404696e69740375726c026964064f626a656374036e65770a616464726573735f6f6604656d6974087472616e7366657204757466387baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020a021f1e68747470733a2f2f6578616d706c652e636f6d2f6176617461722e706e670002010c08030102010d08020001000004120b01120038000c020b0011040c040e0238010c030a03120138020b020b0438030b03020101040003060b0007001108110001020200000003030b0011010200"
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
            "created_at": "1720876964000",
            "updated_at": "1721635671353",
            "object_type": "0x2::timestamp::Timestamp"
          },
          "value": {
            "modify": "0x398d7bd990010000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpsd68l8x",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x90875c2fbf698ce0037b8b8312ec2af3940b2e9a115676684a550bfc37d5df37",
            "size": "227499",
            "created_at": "1720876964000",
            "updated_at": "1720876964000",
            "object_type": "0x3::address_mapping::RoochToBitcoinAddressMapping"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a1136a0ccb4926b299cd2283f0879fc421b2aac7cb5d729e35195c9efbf104b32",
                "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                "owner_bitcoin_address": null,
                "flag": 0,
                "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
                "size": "0",
                "created_at": "1721635671353",
                "updated_at": "1721635671353",
                "object_type": "0x2::object::DynamicField<address, 0x3::bitcoin_address::BitcoinAddress>"
              },
              "value": {
                "new": "0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f2202011f7538675558addbe5f113c450a854324483ebc383639414c2a0b2375f78d47a"
              },
              "fields": []
            }
          ]
        },
        {
          "metadata": {
            "id": "0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f",
            "owner": "rooch10whhuddqjsqsqccq3uqsnk0rxac3ypq9vk6rr3a6mumyxy0el4lsgg4u9w",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721635671353",
            "updated_at": "1721635671353",
            "object_type": "0x2::account::Account"
          },
          "value": {
            "new": "0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f0100000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xc65fc6df9a5957159fd0f8134a4c474068382f382e80ebb2fb915b1f19034960",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1720876964000",
            "updated_at": "1721635671353",
            "object_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0xb7b3f84a0600000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xe469bc7fd50c7e6fcd8945e1df7666839932f160e4ae541378131ff3bb34a31d",
            "owner": "rooch10whhuddqjsqsqccq3uqsnk0rxac3ypq9vk6rr3a6mumyxy0el4lsgg4u9w",
            "owner_bitcoin_address": null,
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": "0",
            "created_at": "1721635671353",
            "updated_at": "1721635671353",
            "object_type": "0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f::avatar::Avatar"
          },
          "value": {
            "new": "0x1e68747470733a2f2f6578616d706c652e636f6d2f6176617461722e706e67"
          },
          "fields": []
        }
      ]
    },
    "events": [
      {
        "event_id": {
          "event_handle_id": "0xe90b1d3753fcaa3b7339e8a0b0b5acb8e13e287284b4272683f68cee7f834e83",
          "event_seq": "0"
        },
        "event_type": "0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f::avatar::AvatarCreatedEvent",
        "event_data": "0x01e469bc7fd50c7e6fcd8945e1df7666839932f160e4ae541378131ff3bb34a31d",
        "event_index": "0",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x8e3089f2c059cc5377a1b6b7c3dcefba8a586697c35de27c2a4b68f81defb69c",
          "event_seq": "35485"
        },
        "event_type": "0x3::coin_store::WithdrawEvent",
        "event_data": "0x011bfd3970bdc6d8e9ad089741d35b2339933c596707a365f1719aa999fa0cede353303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e74f51f0000000000000000000000000000000000000000000000000000000000",
        "event_index": "1",
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": "35567"
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x01c65fc6df9a5957159fd0f8134a4c474068382f382e80ebb2fb915b1f1903496053303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e74f51f0000000000000000000000000000000000000000000000000000000000",
        "event_index": "2",
        "decoded_event_data": null
      }
    ],
    "gas_used": "2094452",
    "is_upgrade": false
  }
}
```

</details>

## 2. 查看部署账户的头像地址

```shell
curl -X "POST" "https://test-seed.rooch.network:443/" \
     -H 'Content-Type: application/json' \
     -d $'{
  "id": 101,
  "jsonrpc": "2.0",
  "method": "rooch_getEventsByEventHandle",
  "params": [
    "0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f::avatar::AvatarCreatedEvent",
    null,
    null,
    true,
    {
      "decode": true
    }
  ]
}'
```

可以得到部署账户的 avatar ObjectId 为 `0xe469bc7fd50c7e6fcd8945e1df7666839932f160e4ae541378131ff3bb34a31d`。

接着查看 Object 的内容：

```shell
rooch object -id 0xe469bc7fd50c7e6fcd8945e1df7666839932f160e4ae541378131ff3bb34a31d
```

可以得到部署账户的头像地址为 `https://example.com/avatar.png` 。

<details>
<summary>点击查看详细结果</summary>

```json
{
  "data": [
    {
      "id": "0xe469bc7fd50c7e6fcd8945e1df7666839932f160e4ae541378131ff3bb34a31d",
      "owner": "rooch10whhuddqjsqsqccq3uqsnk0rxac3ypq9vk6rr3a6mumyxy0el4lsgg4u9w",
      "owner_bitcoin_address": "tb1pra6nse64tzkahe03z0z9p2z5xfzg867rsd3eg9xz5zerwhmc63aq4c08k0",
      "flag": 0,
      "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
      "size": "0",
      "created_at": "1721635671353",
      "updated_at": "1721635671353",
      "object_type": "0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f::avatar::Avatar",
      "value": "0x1e68747470733a2f2f6578616d706c652e636f6d2f6176617461722e706e67",
      "decoded_value": {
        "abilities": 14,
        "type": "0x7baf7e35a094010063008f0109d9e3377112040565b431c7badf364311f9fd7f::avatar::Avatar",
        "value": {
          "url": "https://example.com/avatar.png"
        }
      },
      "tx_order": "3570341",
      "state_index": "9",
      "display_fields": null
    }
  ],
  "next_cursor": {
    "tx_order": "3570341",
    "state_index": "9"
  },
  "has_next_page": false
}
```

</details>
