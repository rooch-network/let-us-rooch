# Rooch 开发基础

## 开发环境搭建

**Mac 系统安装 `rooch` 命令**

下载源码

```ts
git clone https://github.com/rooch-network/rooch.git
```

编译项目, 前提需要安装 `rust`

```ts
cd rooch && cargo build
```

设置环境变量: 参考 `Sui` 和 `Rust` 的做法，在 `$HOME` 目录下创建 `.rooch` 目录, 并将 `rooch` 编译结果目录中的 `rooch/target/debug/rooch` 复制到 `$HOME/.rooch/bin` 目录中

```bash
mkdir -p ~/.rooch/bin

cp ~/Desktop/rooch/target/debug/rooch ~/.rooch/bin
```

如果你的 `terminal` 使用 `zsh shell`, 打开 `.zshrc`, 否则打开 `.bash_profile`,

```bash
vi ~/.zshrc
# or
vi ~/.bash_profile
```

并写入以下内容

```sh
# rooch
export ROOCH_HOME="$HOME/.rooch"
export PATH="$ROOCH_HOME/bin:$PARH"
```

保存后, 通过命令 `source ~/.zshrc` 或重开窗口以加载命令

输入

```bash
rooch -V # rooch 0.6.1

# 初始化
rooch init
```

## rooch 命令介绍

```bash
rooch move new
```

## 创建简单的合约
