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

## task02 部署rooch合约

### 1. 配置rooch本地环境
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
module hello_rooch::hello_rooch {
    use moveos_std::account;
    use std::string;
    struct HelloMessage has key {
        text: string::String
    }
    entry fun say_hello(owner: &signer) {
        let hello = HelloMessage { text: string::utf8(b"Hello Rooch!") };
        account::move_resource_to(owner, hello);
    }
}
```

### 4. 编译并部署合约
```shell
cd quick_start_counter

rooch move build

rooch move publish
```
