# 学习成果
### 初始化 rooch init
```
parallels@ubuntu-linux-22-04-desktop:~$ rooch init
Creating client config file ["/home/parallels/.rooch/rooch_config/rooch.yaml"].
Enter a password to encrypt the keys. Press enter to leave it an empty password:
Generated new keypair for address [rooch1aqh300s4az4ezzge8qyz6n73cplg84n5l2f2gqaqtr4fts53rassldx4kd]
Secret Recovery Phrase : [coach verify ride proud habit kitten render enhance success champion aware hundred]
Rooch client config file generated at /home/parallels/.rooch/rooch_config/rooch.yaml
```

#### 启动本地节点
```
rooch server start
2024-07-09T05:56:59.357425Z  INFO moveos::moveos: execute genesis tx state_root:0xa96fb97e73549e5af4207a8dfab01832dcbfa24cae695ae06462ba6aa48cba53, state_size:32
2024-07-09T05:57:00.750151Z  INFO moveos::moveos: execute genesis tx state_root:0xa96fb97e73549e5af4207a8dfab01832dcbfa24cae695ae06462ba6aa48cba53, state_size:32
2024-07-09T05:57:01.102601Z  INFO rooch_rpc_server: The latest Root object state root: 0xa96fb97e73549e5af4207a8dfab01832dcbfa24cae695ae06462ba6aa48cba53, size: 32
2024-07-09T05:57:01.113311Z  INFO rooch_rpc_server: RPC Server sequencer address: rooch1aqh300s4az4ezzge8qyz6n73cplg84n5l2f2gqaqtr4fts53rassldx4kd(0xe82f17be15e8ab91091938082d4fd1c07e83d674fa92a403a058ea95c2911f61)
2024-07-09T05:57:01.113452Z  INFO rooch_sequencer::actor::sequencer: Load latest sequencer order 0
2024-07-09T05:57:01.113507Z  INFO rooch_sequencer::actor::sequencer: Load latest sequencer accumulator info AccumulatorInfo { accumulator_root: 0xe450cab4e625b1e7a87860f684957d05e2c6ae52c6abea86550fd5fd03ffd758, frozen_subtree_roots: [0xe450cab4e625b1e7a87860f684957d05e2c6ae52c6abea86550fd5fd03ffd758], num_leaves: 1, num_nodes: 1 }
2024-07-09T05:57:01.114517Z  INFO rooch_rpc_server: RPC Server proposer address: rooch1aqh300s4az4ezzge8qyz6n73cplg84n5l2f2gqaqtr4fts53rassldx4kd(0xe82f17be15e8ab91091938082d4fd1c07e83d674fa92a403a058ea95c2911f61)
2024-07-09T05:57:01.115447Z  INFO rooch_rpc_server: acl=Const("*")
2024-07-09T05:57:01.117289Z  INFO rooch_rpc_server: JSON-RPC HTTP Server start listening 0.0.0.0:6767
2024-07-09T05:57:01.117304Z  INFO rooch_rpc_server: Available JSON-RPC methods : ["rooch_getObjectStates", "rooch_queryObjectStates", "rooch_getBalance", "rooch_sendRawTransaction", "rooch_getTransactionsByHash", "rooch_queryEvents", "rooch_executeRawTransaction", "rooch_getEventsByEventHandle", "rooch_queryTransactions", "rooch_executeViewFunction", "rooch_listFieldStates", "rooch_getBalances", "rooch_getModuleABI", "rooch_listStates", "rooch_getTransactionsByOrder", "btc_queryInscriptions", "rooch_getChainID", "rooch_getFieldStates", "btc_queryUTXOs", "rpc.discover", "rooch_getStates"]
The active env was successfully switched to `local`
```
 
