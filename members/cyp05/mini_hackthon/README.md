<div align="center">
  <h1>PuzzleFi
</h1>

 <p>Puzzlefi是基于rooch的全链博弈游戏平台</p>

 <p>
    <a href="https://github.com/qShirley/puzzlefi"><img src="https://badgen.net/badge/icon/github?icon=github&label" alt="GitHub" /></a>
  </p>

</div>

## PuzzleFi

<div>
    <a href="https://puzzlefi.vercel.app/">The first Puzzlefi example</a>
</div>

## 简介

Puzzlefi是基于rooch的全链博弈游戏平台, 独创的游戏收益分发模型可以将博弈双方从平台和玩家转化为玩家与玩家，接入的第一个游戏为石头剪刀布。




## 功能与模块

**Puzzlefi 一共有三个模块**
- ```btc_swap```模块提供提供订单簿市场让用户在rooch挂单通过btc支付交易
- ```puzzle_game```模块提供stake与redeem功能将游戏创作者发行的代币转化为PFC范型代币，转化博弈主体
- ```puzzle_coin```模块提供铸造和销毁PFC范型代币的功能，只有puzzle_game模块可以调用
### ```btc_swap```

- 此模块基于rooch example orderbook实现，可挂单rooch上任意public token并通过btc链支付token来购买
- 交易流程为
  1. 调用```create_market<BaseAsset: key + store>()```方法初始化一个用于交易```BaseAsset```的市场(每种token只需要创建一次)
  2. 卖家调用```list<BaseAsset: key + store>()```方法上架rooch上的token，token会根据单价排序存储在```Marketplace```的```CritbitTree```结构中
  3. 前端通过```query_order<BaseAsset: key + store>()```查询所有订单,通过```query_user_order```查询用户的订单
  4. 买家调用```confirm_order<BaseAsset: key + store>()```来锁定订单，订单默认锁定12小时。锁定订单后在btc链上进行支付，获取支付的txid后调用```buy```方法来解析btc交易，如果付款大于等于订单价格则完成订单将rooch上的token转给买家
  5. 取消订单调用```cancel_order<BaseAsset: key + store>()```处于锁定的订单不能被取消 但会标记成```pending_cancel```状态, 买家不能锁定处于```pending_cancel```的订单
- 目前只支持p2pkh的btc地址进行比特币链上的支付，```asset_transaction_sender```方法会解析输入交易的解锁脚本对应的```pubkey```来确认发送比特币的地址为锁单的地址，通过```effective_transaction_amount```方法来获取卖家收到的金额
### ```puzzle_game```

- 此模块提供代币转化功能，将代币存入池子转化为PFC范型代币，当池子盈利token数量变多时，每个PFC范型代币能换回的token也变多，可以将PFC范型代币理解成池子的股权。
- 石头剪刀币游戏流程
  1. 质押玩家或项目方调用```stake```充值token进StakePool成为进行猜拳玩家的对手方
  2. 猜拳玩家调用```new_finger_game```来创建新游戏，每一轮游戏会创建一个```FingerGame```将池子与玩家的资金都存入其中，一轮结束后发送给胜利方
  3. ```stake/redeem/new_finger_game```都会先结算上一轮游戏再执行对应方法
  4. 质押玩家可以调用```redeem```方法来销毁PFC范型代币赎回池子中token，目前赎回过程中协议会收取0.5%的费用
- 结算逻辑使用了rooch链上随机数作为随机种子，因为玩家无法知道```tx_accumulator_root```，此方法可假设安全,若测试过程中发现其他风险可替换为```drand```

### ```puzzlefi_coin```
- 此模块提供PFC代币的注册，铸造，销毁功能，任何人都可以注册和销毁PFC范型token,但只有```puzzlefi_game```模块才可以铸造代币

## TODO
- 在btc swap中支持更多类型的比特币地址进行交易
- 支持卖家自定义锁仓时间，卖家选择买家是否需要押金
- 实现btc swap的前端
- 支持更多类型的游戏，统计游戏数据
- 修改前端样式，支持多种类型池子，解偶puzzle_game模块，实现所有游戏的通用接入方式
- 为玩家参与游戏提供更多激励方式
- 社区运营

## Q&A
- 如何解析BTC Transaction
  - rooch可以读取比特币交易，解析交易中解锁脚本的pubkey发款对象，再从txout里解析发送给收款方的utxo数量
- 如何保证比特币支付后能收到rooch上token
  - 用户需要先锁单再支付比特币，锁单后卖家必须等锁定时间过后才可以解锁下架，只要在锁定过程中确认了购买就能获取成功
- PFC\<CoinType> 是什么
  - PFC\<CoinType>是一种范型代币，每一种CoinType都可以构建一种自己的PFC代币，等价于StakePool的股权
- PFC和token的兑换比例如何确定
  - current_stake_coin/current_pfc_amount == (coin_amount+current_stake_coin)/(new_pfc_amount + current_pfc_amount)
- 为什么选择在rooch上
  - rooch能读取比特币的状态，同时特有的session key机制为游戏模型提供了良好的体验
