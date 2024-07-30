## 部署本地环境bitcoin&Rooch开发环境

docker 运行task01的bitcoin镜像 详细按参考

`rooch server start `

![1722232867448](image/readme/1722232867448.png)




`rooch rpc request --method rooch_queryObjectStates --params '[{"object_type":"0x4::utxo::UTXO"},  null, "20", {"descending": true,"showDisplay":false}]'`

![1722233020499](image/readme/1722233020499.png)



## 部署 BTC_HOLDER_COIN合约&调用

更新官方例子[btc_holder_coin](https://github.com/rooch-network/rooch/tree/main/examples/btc_holder_coin) 部署 btc_holder_coin 合约
![1722233737001](image/readme/1722233737001.png)

完成初始化 `rooch move run --function default::holder_coin::init`

完成质押

`rooch move run --function default::holder_coin::stake --args object_id:0x826a5e56581ba5ab84c39976f27cf3578cf524308b4ffc123922dfff507e514dd1abd52c43111f84bec296235b697d756846708c3a5b9b3975daff7496793a9f`


![1722235270103](image/readme/1722235270103.png)


rooch rpc request --method rooch_queryObjectStates --params '[{"object_type":"0x845260d38635e139408002cb71fab4c32b70b195cfe06dfa976d4e18902c5966::holder_coin::CoinInfoHolder"},  null, "2", {"descending": true,"showDisplay":false}]'

![1722235422756](image/readme/1722235422756.png)


`rooch move run --function default::holder_coin::claim --args object_id:0x845260d38635e139408002cb71fab4c32b70b195cfe06dfa976d4e18902c5966 --args object_id:0x826a5e56581ba5ab84c39976f27cf3578cf524308b4ffc123922dfff507e514dd1abd52c43111f84bec296235b697d756846708c3a5b9b3975daff7496793a9f`

![1722235955547](image/readme/1722235955547.png)

rooch account balance --json

![1722236004567](image/readme/1722236004567.png)
