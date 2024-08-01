# 学习日志

## task01 运行一个比特币全节点

### 1. 配置 bitcoin 本地环境

```shell
brew install bitcoin
```

### 2. 启动 bitcoin 服务并同步区块和交易信息

```shell
bitcoind -daemon -txindex
```

### 3. 查看交易信息

```shell
bitcoin-cli getblockhash <height> # 获取区块高度哈希值

bitcoin-cli getblock <blockhash> # 获取区块详细信息
```

## task02 部署 rooch 合约

### 1. 配置 rooch 本地环境

```shell
# clone rooch
git clone https://github.com/rooch-network/rooch.git

# Building from source
cargo build && cp target/debug/rooch ~/.cargo/bin/

# initialize Rooch config
rooch init
```

### 2. 创建合约项目

```shell
# Creating a new Move project
rooch move new quick_start_counter
```

### 3. 创建合约

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

### 4. 编译并部署合约

```shell
cd quick_start_counter

rooch move build

rooch move publish
```

### 5. 常用 rooch 命令

```shell
# get account list
rooch account list

# get account balance
rooch account balance

# get env list
rooch env list
```

## task03 搭建本地 Bitcoin 以及 Rooch 开发环境，部署一个和 Bitcoin 数据交互的 Move 智能合约，并且进行调用。

### 1.搭建本地 Bitcoin 以及 Rooch 开发环境

参考 rooch 提供的官方[脚本](https://github.com/rooch-network/rooch/tree/main/scripts/bitcoin)，搭建 Bitcoin regtest 测试网以及启动 rooch server。  
pre-requirements:

1. 配置 docker 环境，本文通过安装[docker-desktop](https://www.docker.com/products/docker-desktop/)配置 docker

2. 查看 docker 状态，正常打开 docker-desktop 就能连接上

```shell
docker info
```

3. 执行脚本，将启动 bitcoin regtest 的 docker 容器

```shell
git clone https://github.com/rooch-network/let-us-rooch

cd /<your_path>/rooch/scripts/bitcoin/node & ./run_local_node_docker.sh
```

正常输出

```shell
bitcoind_regtest
```

查看容器运行状态

```shell
docker ps
```

将服务设置别名

```shell
alias bitcoin_regtest='docker exec -it bitcoind_regtest bitcoin-cli'
```

启动 rooch 服务并清除之前的数据

```shell
rooch server clean
rooch server start --btc-rpc-url http://127.0.0.1:18443 --btc-rpc-username roochuser --btc-rpc-password roochpass
```

### 2.部署一个和 Bitcoin 数据交互的 Move 智能合约，并且进行调用。

在这里使用 rooch 官方提供的 example 里面的[btc_holder_coin](https://github.com/rooch-network/rooch/tree/main/examples/btc_holder_coin)合约例子。这份合约实现了一个名为 btc_holder_coin::holder_coin 的 Move 模块，旨在通过质押（stake）比特币未花费交易输出（UTXO）来获取 BTC Holder Coin (HDC)。

1. 部署合约，进入到 rooch/examples/btc_holder_coin 目录，执行部署

```
cd /<your_path>/rooch/examples/btc_holder_coin
rooch move publish --named-addresses btc_holder_coin=default
```

2. 合约部署好了，调用 init 函数进行初始化

```
rooch move run --function default::holder_coin::init
```

注意记录发布合约的 contract_id

3. 初始化完之后，开始质押，质押需要 utxo，所以先给当前账户生成区块。

```
# 查看账户
rooch account list --json
# 生成区块
bitcoin-cli generatetoaddress 101 <bitcoin_address>
```

4. 区块生成之后，就可以查询它的 utxo id 了

```
rooch rpc request --method rooch_queryObjectStates --params '[{"object_type":"0x4::utxo::UTXO"},  null, "20", {"descending": true,"showDisplay":false}]'
```

注意选择 owner 是当前账户地址的 objectId

5. 接下来开始质押 utxo，调用质押函数

```
rooch move run --function default::holder_coin::stake --args object_id:<your_utxo_id>
```

6. 质押完成之后，调用 claim 函数，领取代币，需要知道 CoinInfoHolder 的 objectId。

### 查询 CoinInfoHolder 的 objectId

```
rooch rpc request --method rooch_queryObjectStates --params '[{"object_type":"<your_utxo_id>::holder_coin::CoinInfoHolder"},  null, "2", {"descending": true,"showDisplay":false}]'
```

### 调用 claim，领取代币

```
rooch move run --function default::holder_coin::claim --args object_id:<your_contract_id> --args object_id:<your_utxo_id>
```

7. 查询代币余额

```
rooch account balance --json
```
