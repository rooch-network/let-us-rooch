# publish
```
rooch move publish -p . --sender-account default --named-addresses mint_nft=default
```

# publish log
```
Publish modules to address: rooch16tf98cay0h6rt4kaew4kcg5u2acr793fty7kdumq2224x69km9cqkezhzc(0xd2d253e3a47df435d6ddcbab6c229c57703f1629593d66f36052955368b6d970)
{
  "sequence_info": {
    "tx_order": "3",
    "tx_order_signature": "0x0158c0e3cf317ba78b8114166b14be87ed2e590c9fc43bd93679139c682d5bf2c8503ab3c0fc794c838022301924f98199a4d476ee54b826aad29450af92a5b4c3021705aa6decf580728ebf36afa7b2f60abcc3c8050a5e4116e00428fcfbd9d6e9",
    "tx_accumulator_root": "0xb22eddbdb80fb4c88e2cc5a211bb8a08612e29a058ada145811902b1e917bc54",
    "tx_timestamp": "1721222017389"
  },
  "execution_info": {
    "tx_hash": "0x1d44c351ae78330bcc308142148a64381ad96ed5c6b305565c3e72c7b5463817",
    "state_root": "0x7b4d07456f28ce354f5d2d2b3ab0c87febc4a2ccca0cbf8d24746efedab77456",
    "event_root": "0xef46cff2a0773f23e67ba4562466da9070a97792cc9cb8ce159eeb30a51253c4",
    "gas_used": 551666,
    "status": {
      "type": "miscellaneouserror"
    }
  },
  "output": {
    "status": {
      "type": "miscellaneouserror"
    },
    "changeset": {
      "root_metadata": {
        "id": "0x",
        "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqcm08j4",
        "flag": 1,
        "state_root": "0x1e617cc0390ee563fccf12e328d56ab727a76f5030dece06b285cc597357f652",
        "size": 39,
        "created_at": 0,
        "updated_at": 0,
        "value_type": "0x2::object::Root"
      },
      "changes": [
        {
          "metadata": {
            "id": "0x05921974509dbe44ab84328a625f4a6580a5f89dff3e4e2dec448cb2b1c7f5b9",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "flag": 1,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": 0,
            "created_at": 1721007466187,
            "updated_at": 1721222017389,
            "value_type": "0x2::object::Timestamp"
          },
          "value": {
            "modify": "0x6db1d3c090010000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpsd68l8x",
            "flag": 0,
            "state_root": "0x9ae2d96968441c46c3671bd6f75840717a27497fd78d49719e05543736308b9b",
            "size": 3,
            "created_at": 0,
            "updated_at": 0,
            "value_type": "0x3::address_mapping::RoochToBitcoinAddressMapping"
          },
          "value": null,
          "fields": [
            {
              "metadata": {
                "id": "0x5024c060f254a47033bd9ce9043854d6b26f1f5929de9cb0b7038aa486c0ff0a7c415bf4e85c837124451fc41dc37f71053a6cb7164fc22c2eaf080ec66c99cd",
                "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
                "flag": 0,
                "state_root": null,
                "size": 0,
                "created_at": 1721222017389,
                "updated_at": 1721222017389,
                "value_type": "0x2::object::DynamicField<address, 0x3::bitcoin_address::BitcoinAddress>"
              },
              "value": {
                "new": "0xd2d253e3a47df435d6ddcbab6c229c57703f1629593d66f36052955368b6d97022020148fc56c9bceb1be070c5ebf16cc1ed679d14c7a6aa2cc38e50af81eef0dc87dd"
              },
              "fields": []
            }
          ]
        },
        {
          "metadata": {
            "id": "0x56e708f560452ee267f8ae7f7a745ba6ff8eb6ba3fd56d5624c41ddef712d41d",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhxqaen",
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": 0,
            "created_at": 1721007466187,
            "updated_at": 1721222017389,
            "value_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0x25a59b000000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xd2d253e3a47df435d6ddcbab6c229c57703f1629593d66f36052955368b6d970",
            "owner": "rooch16tf98cay0h6rt4kaew4kcg5u2acr793fty7kdumq2224x69km9cqkezhzc",
            "flag": 0,
            "state_root": null,
            "size": 0,
            "created_at": 1721222017389,
            "updated_at": 1721222017389,
            "value_type": "0x2::account::Account"
          },
          "value": {
            "new": "0xd2d253e3a47df435d6ddcbab6c229c57703f1629593d66f36052955368b6d9700100000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xdf9438a4df4398f180e887c0165940470de398c7d07c031b993cc8bde0980ca0",
            "owner": "rooch16tf98cay0h6rt4kaew4kcg5u2acr793fty7kdumq2224x69km9cqkezhzc",
            "flag": 0,
            "state_root": null,
            "size": 0,
            "created_at": 1721222017389,
            "updated_at": 1721222017389,
            "value_type": "0x3::coin_store::CoinStore<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "new": "0x0e76ed050000000000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        },
        {
          "metadata": {
            "id": "0xf81628c3bf85c3fc628f29a3739365d4428101fbbecca0dcc7e3851f34faea6b",
            "owner": "rooch1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpsd68l8x",
            "flag": 0,
            "state_root": "0x5350415253455f4d45524b4c455f504c414345484f4c4445525f484153480000",
            "size": 0,
            "created_at": 1721007466187,
            "updated_at": 1721222017389,
            "value_type": "0x3::coin::CoinInfo<0x3::gas_coin::GasCoin>"
          },
          "value": {
            "modify": "0x53303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e0e526f6f63682047617320436f696e03524743080002661cf35a0000000000000000000000000000000000000000000000000000"
          },
          "fields": []
        }
      ]
    },
    "events": [
      {
        "event_id": {
          "event_handle_id": "0x358779b791ef606d7f07df8881c1939f26de95119486b60745a9c3127ae8fd37",
          "event_seq": 2
        },
        "event_type": "0x3::coin::MintEvent",
        "event_data": "0x53303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e00e1f50500000000000000000000000000000000000000000000000000000000",
        "event_index": 0,
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0xdebc7ccc8fa8855fad9fdd2919e875e06bcfa9b11cdc53c3247e0f81239852e2",
          "event_seq": 3
        },
        "event_type": "0x3::coin_store::CreateEvent",
        "event_data": "0x01df9438a4df4398f180e887c0165940470de398c7d07c031b993cc8bde0980ca053303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e",
        "event_index": 1,
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": 4
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x01df9438a4df4398f180e887c0165940470de398c7d07c031b993cc8bde0980ca053303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696e00e1f50500000000000000000000000000000000000000000000000000000000",
        "event_index": 2,
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x8e3089f2c059cc5377a1b6b7c3dcefba8a586697c35de27c2a4b68f81defb69c",
          "event_seq": 2
        },
        "event_type": "0x3::coin_store::WithdrawEvent",
        "event_data": "0x01df9438a4df4398f180e887c0165940470de398c7d07c031b993cc8bde0980ca053303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696ef26a080000000000000000000000000000000000000000000000000000000000",
        "event_index": 3,
        "decoded_event_data": null
      },
      {
        "event_id": {
          "event_handle_id": "0x6ab771425e05fad096ce70d6ca4903de7cca732ee4c9f6820eb215be288e98dd",
          "event_seq": 5
        },
        "event_type": "0x3::coin_store::DepositEvent",
        "event_data": "0x0156e708f560452ee267f8ae7f7a745ba6ff8eb6ba3fd56d5624c41ddef712d41d53303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333a3a6761735f636f696e3a3a476173436f696ef26a080000000000000000000000000000000000000000000000000000000000",
        "event_index": 4,
        "decoded_event_data": null
      }
    ],
    "gas_used": 551666,
    "is_upgrade": false
  }
}
```
