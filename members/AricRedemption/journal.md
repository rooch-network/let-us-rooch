# 学习日志

## task01 运行一个比特币全节点

### 1. 配置bitcoin本地环境
```shell
brew install bitcoin
```

### 2. 启动bitcoin服务并同步区块和交易信息
```shell
bitcoind -daemon -txindex
```

### 3. 查看交易信息
```shell
bitcoin-cli getblockhash <height> # 获取区块高度哈希值

bitcoin-cli getblock <blockhash> # 获取区块详细信息
```
