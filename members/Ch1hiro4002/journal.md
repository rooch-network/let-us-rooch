# 学习日志

## Task1：搭建比特币节点

通过 `docker` 运行 `bitcore` 节点

```shell
# 拉取镜像
sudo docker pull lncm/bitcoind:v25.1

# 运行
sudo docker run lncm/bitcoind:v25.1
```

运行结果：

![](./imgs/01.png)
<<<<<<< HEAD

## Task2：部署第一个合约

1. 安装 `rooch` 编译器

   ```
   # 安装依赖项
   sudo apt install git curl cmake make gcc lld pkg-config libssl-dev libclang-dev libsqlite3-dev g++ protobuf-compiler
   # 安装 Rust
   curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
   # 克隆源码
   git clone https://github.com/rooch-network/rooch.git
   # 安装 Rooch
   cd rooch && cargo build && cp target/debug/rooch ~/.cargo/bin/
   ```

2. 连接到 `Rooch`

   ```
   # 初始化
   rooch init
   ```

3. 切换到测试网
   查看环境列表：`rooch env list`
   切换到 test：`rooch env switch --alias test`

4. 创建第一个项目
   `rooch move new hello_rooch`

5. 创建合约并添加 move 代码

   ```rust
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

6. 部署合约
   `rooch move publish`

![](./imgs/02.png)
=======
