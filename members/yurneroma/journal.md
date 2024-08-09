# 学习日志

## task1 搭建比特币节点

###  下载bitcoin code， 源码编译
    
   ![alt text](images/image.png)
   1. autogen
    ![alt text](images/image1.png)
   2. configure 
    ![alt text](images/image2.png)
   3. compile 
   ![alt text](images/image3.png)
   ![alt text](images/image5.png)
   
###  启动bitcoind 
   ![alt text](images/image6.png)
   预同步阶段
   ![alt text](images/image7.png)
   同步阶段
   ![alt text](images/image8.png)

### 查看节点信息

![alt text](images/image9.png)

### 获取区块高度

``` bash
    ./src/bitcoin-cli getblockcount 
```
![alt text](images/image10.png)


