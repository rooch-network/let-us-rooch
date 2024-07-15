# 学习日志

## 搭建比特币节点

1. 下载程序

下载地址： [https://bitcoincore.org/en/download/](https://bitcoincore.org/en/download/)
根据你的操作系统选择对应的程序。

2. 运行 bitcoind 启动节点

下载解压，目录结构如下： 

![alt text](image.png)

3. 启动节点

```shell
./bin/bitcoind 
```

节点同步状态如下:

先经历一个比较漫长的 `Pre-synchronizing` 阶段
![alt text](image-1.png)

再经历一个比较漫长的 `Synchronizing blockheaders` 阶段
![alt text](image-3.png)

最后，开始同步区块

![alt text](image-4.png)

4. 查看节点信息

```shell
./bin/bitcoin-cli -getinfo
```

![alt text](image-2.png)

获取同步的区块高度

```shell
./bin/bitcoin-cli -getblockcount
```

![alt text](image-5.png)
