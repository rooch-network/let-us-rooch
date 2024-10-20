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

## 4. 调研rooch和sui、aptos的区别

# Rooch StarTrek Q&A

## 概念相关

### 现在 Rooch 最新发行版是？

目前最新发行版为 [![GitHub Release](https://img.shields.io/github/v/release/rooch-network/rooch)](https://github.com/rooch-network/rooch/releases)

### Rooch 跟 Bitcoin 是什么关系？

Rooch 是比特币二层（Layer2），扩展比特币的应用生态，借助 Move 语言的资产安全优势，开启比特币经济的无限可能。

### 有测试网吗？

目前 Rooch 已经上线测试网，`devnet` 和 `testnet` 网络，`devnet` 会定期清理数据，`testnet` 将作为永久激励测试网。

目前 Rooch 发布的发行版可以在 [GitHub release 页面]下载(https://github.com/rooch-network/rooch/releases)。

### Rooch 与 Aptos/Sui 的区别？

都使用 Move 语言进行智能合约编程，Rooch 是 BTC L2，属于应用扩展层。Aptos/Sui 是公链跟 BTC 属于同一层级。

### 一般用什么钱包？

Rooch 作为 BTC L2 需要使用 Bitcoin 相关的钱包，常用的有 UniSat 和 OKX。

目前 Rooch Portal 仅支持 UniSat，后续将支持更多钱包，详细使用请参考[Rooch Portal](https://rooch.network/zh-CN/learn/miscellaneous/portal)。

### Unisat 怎么切换测试网？

在 `设置 -> Network` 可以切换，`LIVENET` 为比特币主网，`TESTNET` 为比特币测试网。

### 测试网以什么开头，用什么类型的比特币地址？

Rooch 使用 Taproot 类型的比特币地址，以 `tb1p` 开头。

### Bitcoin 主网和测试网地址不一样？

不一样，但是可以转换。

## 技术相关

### Rooch Move 与 Aptos Move / Sui Move 区别？

Rooch 的 Move 是基于 [Core Move](https://github.com/move-language/move) 修改的版本，跟 Aptos、Sui 都有不同，详细的使用可以参考 [Counter quick start](../../build/tutorial/counter) 来详细了解 Rooch Move 的用法。

参考文档：

- [Rooch Object](https://rooch.network/zh-CN/learn/core-concepts/objects/object)
- [Rooch Framework 源码](https://github.com/rooch-network/rooch/tree/main/frameworks)

### BTC 资产与 Rooch 资产问题

Q：BTC 资产仓库变成 Rooch 资产之后可以改变他的状态或者内容吗？比如再套一层之类的，如果可以改变，那返回的时候是不是也要考虑需要转换成原始资产模式返回？

A：把 BTC 资产变成 Move 的 Object，可以通过动态字段往里面写状态。

Q：资产的处理，比如 BTC 资产如果过来，我们想对资产进行更多的处理的时候，比如加一些标记，那这个资产回到 BTC 的时候是不是需要有额外的处理？

A：关于这方面，[资产跃迁](https://rooch.network/zh-CN/learn/in-depth-tech/client-side-validation.zh-CN#%E8%B5%84%E4%BA%A7%E8%B7%83%E8%BF%81%E5%8D%8F%E8%AE%AEasset-leap-protocol)的文档有详细介绍。

### 关于 Rooch 的 Bitcoin 扩展方案概述看这两篇博客

1. [Stackable L2 — 一种新的区块链扩容方案](https://rooch.network/zh-CN/blog/stackable-l2)
2. [Rooch Network - Bitcoin 的应用层](https://rooch.network/zh-CN/blog/the-application-layer-of-bitcoin)

### 如何用 Move 针对 Bitcoin 编程以及例子参看文档

- [扩展 Bitcoin](https://rooch.network/zh-CN/build/bitcoin)

## 相关链接

**[Website](https://rooch.network/) | [Discord](https://discord.com/invite/rooch) | [Twitter](https://x.com/RoochNetwork) | [Telegram](https://t.me/roochnetwork) | [Github](https://github.com/rooch-network/)**


## 5. 部署一个rooch dapp

# 第一个 Rooch dApp

本篇教程主要介绍如何使用 Rooch 提供的前端模板来实现一个简易计数器 dApp。

![](/docs/first-dapp/dapp-main.png)

[模板 Repo](https://github.com/rooch-network/my-first-rooch-dapp)

## 克隆模板源码

```bash
git clone https://github.com/rooch-network/my-first-rooch-dapp
```

## 初始化项目

安装 dApp 所需依赖：

```bash
cd my-first-rooch-dapp
bun install
## if you are using yarn
yarn install
```

运行 dApp：
```bash
bun dev
```

当一切顺利完成后，在浏览器访问本地的预览链接，就能看到如下效果：

![](/docs/first-dapp/dapp-counter.png)

恭喜你！完成以上步骤后，说明你的 dApp 已经成功运行了，为了完成完整的链上交互，我们还需要部署合约。

## 确认 Rooch 当前的 Network

我们使用 Testnet 测试网来部署 Counter 应用，使用 `rooch env switch` 来切换网络：

```bash
rooch env switch --alias test

The active environment was successfully switched to `test`
```

## 部署合约

在 `counter_contract` 目录里，我们可以看到 Counter dApp 的合约源码。
进入 `counter_contract` 目录，使用如下命令来部署合约。

请注意，在部署前请确保自己的地址下有足够的 Gas Fee，我们可以使用如下命令来查询：

```bash
rooch account balance

## output

      Coin Type        |      Symbol      | Decimals |  Balance              
--------------------------------------------------------------------
0x3::gas_coin::GasCoin |       RGC        |     8    | 1939625968 
```

确认有足够的 Gas Fee 后，就可以使用下面的命令来部署合约了。

```bash
## in counter_contract directory
rooch move publish --named-addresses quick_start_counter=default
```

在部署完成后，我们可以看到命令行的输出：

```bash
BUILDING quick_start_counter
Publish modules to address: rooch1e7qm7jqangukl37qs49ckv7j4w47zyu5cr2gd9tmzal89q9sudqqzhy92t
(0xcf81bf481d9a396fc7c0854b8b33d2ababe11394c0d486957b177e7280b0e340)
...
```

如上命令行的输出，Counter 就被部署在了 `0xcf81bf481d9a396fc7c0854b8b33d2ababe11394c0d486957b177e7280b0e340` 这个地址上了。


## 修改前端配置

找到前端项目中的 `src/App.tsx` 文件，修改 `counterAddress` 这个常量：

```tsx
// Publish address of the counter contract
const counterAddress = "YOUR_COUNTER_ADDRESS";
```

替换上一步部署的 Counter 合约地址：

```bash
const counterAddress = "0xcf81bf481d9a396fc7c0854b8b33d2ababe11394c0d486957b177e7280b0e340"
```

> 注意：这里的合约地址会与部署者的地址有关，请确认。

## 合约交互

连接上钱包，我们可以在 Session key 区域点击 Create 按钮，创建会话密钥：

创建完成后，即可看到 Session Key 的相关信息：

![](/docs/first-dapp/dapp-integration.png)

当替换完合约地址后, dApp Integration 区域里也可以看到 `Counter Value` 的计数。

完成上述步骤后，就可以在 dApp Integration 区域里点击 `Increase Counter Value` 按钮，调用合约并增加计数器的数值了。

## 总结

完成上述步骤后，你已经掌握了 **创建 dApp**，**部署合约** 以及 **前端与合约进行基本的交互** 的技能了。

