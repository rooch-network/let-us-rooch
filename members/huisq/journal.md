# 学习日志

## Week 1
1. 掌握比特币全节点运行
2. 安装rooch

## Week 2 
### 任务：基于Rooch Object实现和部署一个Move Contract
[部署log](./task2/publish_log.md)

1. Connect to devnet
- check list of environments in configuration
```
rooch env list
```
- switch to devnet using:
```
rooch env switch --alias dev
```
2. Initialize project with example template:
```
rooch move new hello_rooch
```
This will create a new Rooch contract project directory with Move.toml file.
Go ahead and create a new .move file in the sources folder, and start writing. 
Remember to change your module name in the Move.toml file based on the new module name you set in your Move contract module. 

3. 一些有用的文档，能帮助你顺利写出第一个Rooch合约:
    - [CLI](https://rooch.network/build/reference/rooch-cli)
    - libraries：
        - std: 0x1 [MoveStdlib](https://github.com/rooch-network/rooch/blob/main/frameworks/move-stdlib/doc)
        - moveos_std: 0x2 [MoveosStdlib](https://github.com/rooch-network/rooch/blob/main/frameworks/moveos-stdlib/doc)
        - rooch_framework: 0x3 [RoochFramework](https://github.com/rooch-network/rooch/blob/main/frameworks/rooch-framework/doc)
        - bitcoin_move: 0x4 [BitcoinMove](https://github.com/rooch-network/rooch/blob/main/frameworks/bitcoin-move/doc)
    - Rooch Object Model 和 Sui&Aptos 几个move生态的object model分析和对比:
        - [Week 2 Summary](./summary.md)
        

## Week 3
### 任务：构思一个 Bitcoin 生态的应用或者游戏，可以利用 Rooch 提供的特性，写成文章或者 Github Issue
[Safeswap构思](./task3/safeswap.md)
关于SessionKey的一些分析：[Week 3 Summary](./summary.md)

## Week 4
### 列举出至少三个Rooch Move和 SUI/APTOS Move 的区别
- Rooch Object Model 和 Sui&Aptos 几个move生态的object model分析和对比:
    - [Week 2 Summary](./summary.md)

## Week 5
### 使用pnpm 创建一个 Counter 工程
[img5](./img/task5.png)