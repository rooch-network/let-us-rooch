# 学习日志

## task1 搭建比特币节点

1. 下载程序

下载地址： [https://bitcoincore.org/en/download/](https://bitcoincore.org/en/download/)
根据你的操作系统选择对应的程序。

2. 运行 bitcoind 启动节点

下载解压，目录结构如下：

![alt text](images/image.png)

3. 启动节点

```shell
./bin/bitcoind
```

节点同步状态如下:

先经历一个比较漫长的 `Pre-synchronizing` 阶段
![alt text](images/image-1.png)

再经历一个比较漫长的 `Synchronizing blockheaders` 阶段
![alt text](images/image-3.png)

最后，开始同步区块

![alt text](images/image-4.png)

4. 查看节点信息

```shell
./bin/bitcoin-cli -getinfo
```

![alt text](images/image-2.png)

获取同步的区块高度

```shell
./bin/bitcoin-cli -getblockcount
```

![alt text](images/image-5.png)

5. 方便 rpc 请求，配置文件 `bitcoin.conf` 添加 rpc 授权信息

```conf
rpcuser=admin
rpcpassword=admin
```

启动的时候，使用全路径指定配置文件:

```shell
./bin/bitcoind -conf=/root/bitcoin-27.1/bitcoin.conf
```

curl 使用 rpc 调用，获取节点信息:

```shell
 curl -u admin:admin --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getblockchaininfo", "params": [] }'  http://127.0.0.1:8332/ | jq
```

![alt text](images/image-6.png)

## task2 部署一个和 bitcoin 交互的 move 合约

### 合约简单说明

合约实验通过修改基本的 `Counter` 合约来实现。
合约部署完成后，会在当前模块的地址创建一个 `Counter` 的实例。
`Counter` 实例中包含了， 创建模块的时候比特币公链的块高。

后续 increment 方法可以对 Counter 实例中的值进行加一操作。这里会判断下调用合约模块时候的比特币块高。
块高必须满足一定的条件,才会执行。以此借助 btc 的 pow 能力，控制 合约的执行。

### 实现步骤

1. 准备基本环境

a. 创建 rooch 的基本账号,领取一定的 gas token 。

使用 rooch 命令行工具，创建一个账号。

```shell
rooch init
```

注意，保存好，你的助记词和私钥，一定不要泄露。
gas token 需要从 `Discord`中获取。

使用 `rooch account balance` 确认 gas token .

![alt text](images/image_gas.png)

b. 修改配置文件，确认网络状态。

配置文件目录: ~/.rooch/rooch_config/rooch.yaml

内容如下:

```yaml
keystore_path: /Users/yourname/.rooch/rooch_config/rooch.keystore
active_address: 0xaef440f63dbbad67c3ce25dfc863e18dc603565e57f009eaa846d31356299558
envs:
  - alias: local
    rpc: http://0.0.0.0:50051
    ws: null
  - alias: dev
    rpc: https://dev-seed.rooch.network:443/
    ws: null
  - alias: test
    rpc: https://test-seed.rooch.network:443/
    ws: null
active_env: test
```

其中 `active_address` 配置默认账户，可以根据自己的账户地址进行修改。
`active_env` 配置网络环境，这里使用的是 test 网络。

2. 编写合约

创建项目:

```shell
rooch move new btc-locker
```

修改 Move.toml 配置

```toml
[package]
name = "btc-locker"
version = "0.0.1"

[dependencies]
MoveStdlib = { local = "../../rooch/frameworks/move-stdlib" }
MoveosStdlib = { local = "../../rooch/frameworks/moveos-stdlib" }
RoochFramework = { local = "../../rooch/frameworks/rooch-framework" }
BitcoinMove = { local = "../../rooch/frameworks/bitcoin-move" }

[addresses]
btc_locker = "0xaef440f63dbbad67c3ce25dfc863e18dc603565e57f009eaa846d31356299558"
std = "0x1"
moveos_std = "0x2"
rooch_framework = "0x3"
bitcoin_move =  "0x4"

```

其中 dependencies 模块默认使用 github 远程配置，可以使用本地的 local 配置，加速合约的编译。
合约中会使用 bitcoin 的相关功能，所以加入 BitcoinMove 和 bitcoin_move 两个定义。

3. 编译合约

Move 合约代码如下：

```move
// Copyright (c) RoochNetwork
// SPDX-License-Identifier: Apache-2.0

module btc_locker::counter {
    use moveos_std::account;
    use bitcoin_move::bitcoin;
    use bitcoin_move::types::{BlockHeightHash,unpack_block_height_hash};
    use std::option::{Self};

    const ErrorBitcoinClientError: u64 = 10;
    const ErrorBitcoinBlockError: u64 = 11;

    struct Counter has key {
        count_value: u64,
        last_block: u64
    }

    fun init() {
        let signer = moveos_std::signer::module_signer<Counter>();
        account::move_resource_to(&signer, Counter { count_value: 0,last_block:0 });
    }

    entry fun increase() {
        let latest_block = bitcoin::get_latest_block();
        assert!(option::is_some<BlockHeightHash>(&latest_block), ErrorBitcoinClientError);
        let (height,_hash) = unpack_block_height_hash(option::destroy_some(latest_block));
        let counter = account::borrow_mut_resource<Counter>(@btc_locker);
        assert!(height % 10 >= 2, ErrorBitcoinBlockError);
        counter.count_value = counter.count_value + height % 10;
        counter.last_block = height;
    }
}
```

其中 Counter 对象，分配到发布的地址上，供所有使用者共享。其中 increse 操作的时候，会判断 当前比特币的块高，个位数一定要 大于等于 2 ，才会往下执行。

** rooch 上的 btc 高度相比于 btc 本身的高度，延后了 3 个区块。 **

4. 合约编译、发布

```shell
rooch move build
rooch move publish --max-gas-amount 50000000
```

如果发布失败，适度的调整 最大 gas 消耗。

发布成功后，就可以通过 cli 来调用合约获取 Counter 对象。

```shell
rooch resource --address 0xaef440f63dbbad67c3ce25dfc863e18dc603565e57f009eaa846d31356299558 --resource 0xaef440f63dbbad67c3ce25dfc863e18dc603565e57f009eaa846d31356299558::counter::Counter
```

![alt text](images/btc_locker_counter.png)

cli 调用 increment 方法

```shell
rooch move run --function 0xaef440f63dbbad67c3ce25dfc863e18dc603565e57f009eaa846d31356299558::counter::increase --max-gas-amount 50000000 | jq ".execution_info"
```

status 可以表示交易是否生效。

![alt text](images/btc_locker_call.png)
