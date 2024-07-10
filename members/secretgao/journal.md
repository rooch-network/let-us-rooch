# 学习日志
####  linux ubuntu 下面安装命令
##### 安装依赖

```
sudo apt install git curl gcc lld pkg-config libssl-dev libclang-dev libsqlite3-dev g++
```
##### 安装 Rust
```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```
##### 克隆源码
```
git clone https://github.com/rooch-network/rooch.git
```
##### 编译并安装 Rooch
```
cd rooch && cargo build && cp target/debug/rooch ~/.cargo/bin/
```

##### 安装遇到的错误
* 问题 1 
```
= note: collect2: fatal error: ld terminated with signal 9 [Killed]
compilation terminated.

error: could not compile framework-release (bin "framework-release") due to 1 previous error
warning: build failed, waiting for other jobs to finish...

```

* 1.解决办法
```
export CARGO_BUILD_JOBS=1
cargo build
```

* 问题 2
```
e" "-Wl,-z,relro,-z,now" "-nodefaultlibs"
= note: collect2: fatal error: ld terminated with signal 9 [Killed]
compilation terminated.

error: could not compile rooch (bin "rooch") due to 1 previous error
```
* 2.解决办法
```
sudo fallocate -l 4G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
sudo swapon --show
```