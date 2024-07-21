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
export PATH="$ROOCH_HOME/bin:$PATH"
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
# account
rooch account create # 创建账号
rooch account list # 列出所有账户
rooch account switch -a <ADDRESS> # 切换账号
rooch account balance # 获取余额
rooch account export # 导出私钥

# env
rooch env list # 列出环境列表, 默认有local/dev/test
rooch env add --alias <ALIAS> --rpc <RPC> # 添加环境
rooch env switch --alias <ALIAS> # 切换环境
rooch env remove --env <ENV> # 移除环境

# move
rooch move new <NAME> # 创建项目
rooch move build # 编译项目
rooch move info
rooch move test # 运行单测
rooch move run --function <FUNCTION>
rooch move publish # 部署到链上

# rpc
rooch rpc request --method <METHOD> --params <PARAMS> # 发送 rpc 请求
```

## 创建简单的合约

通过下面的名称创建项目

```bash
rooch move new first_rooch_package
code first_rooch_package
```

`sources` 目录下创建 `first_module.move` 文件, 写入下列内容

```rust
module first_rooch_package::first_module {
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

编译并部署

```bash
rooch move build

rooch move publish
```

terminal 输出

```bash
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
UPDATING GIT DEPENDENCY https://github.com/rooch-network/rooch.git
INCLUDING DEPENDENCY MoveStdlib
INCLUDING DEPENDENCY MoveosStdlib
INCLUDING DEPENDENCY RoochFramework
BUILDING first_rooch_package
Publish modules to address: rooch1s09h4u6k2swv4y6gdzfkmge053327yzpg8wz9sqlhehgeudx2gjqxvsuru(0x83cb7af356541cca934868936da32fa462af104141dc22c01fbe6e8cf1a65224)
Enter the password to publish:
{
  "sequence_info": {
    "tx_order": "4",
    "tx_order_signature": "0x010f39a578646078839cf5646d3595d8585a011b03b47852b73cefe12c73f3a41a5d6907035aff95ae8e0dc5eabe3990ec95d2ee44c6a9fce5f7a6140507bf729a026c9e5a00643a706d3826424f766bbbb08adada4dc357c1b279ad4662d2fd1e2e",
    "tx_accumulator_root": "0x8c834e01af8a2ba0e20534bb46dc1e59a666f2658325a8ae5e74f5bdb28f57c8",
    "tx_timestamp": "1721583481336"
  },
  "execution_info": {
    "tx_hash": "0x01dd26b0e6a515d51974796455bcbd7b240db615fe4fa012ffbbca815033b465",
    "state_root": "0xb196e4095d1fb022bcb47193ab68c03057f17a377a0da9c6b2249671cecdbd95",
    "event_root": "0x5d970bec69d99bf0b0247e8a481dc964f7ee96d00bb85b74be33afbc2d4059c3",
    "gas_used": "1674177",
    "status": {
      "type": "executed"
    }
  }
  ....
}
```
