# Rooch Move和 SUI/APTOS Move 的区别
## 1.模块命名和结构不同：
MoveStdlib: 0x01
MoveosStdlib: 0x02
RoochFramework: 0x03
BitcoinMove: 0x04
首先是地址标识符在move中不同，然后是Rooch Move 中的一些模块具有独特的命名，例如 account_authentication.move、onchain_coinfig.move 和 simple_rng.move。这些模块在 SUI/APTOS Move 中并不存在，或存在名称和功能上的差异。

SUI/APTOS Move 使用标准的 Move 库模块，如 account.move、vector.move 等，而 Rooch Move 则在此基础上进行了定制和扩展，添加了许多特定于 Rooch 平台的模块。

## 2.功能扩展：
Rooch Move 包含了一些特定的功能扩展模块，例如 BitcoinMove 模块，包括 bitcoin.move、bitcoin_hash.move 等，这些模块专门处理与比特币相关的功能。而 SUI/APTOS Move 没有这类扩展。
Rooch Move 提供了 address_mapping.move 用于处理多链地址到 Rooch 地址的映射，这在 SUI/APTOS Move 中没有。

## 3.配置和管理：
Rooch Move 提供了专门的 onchain_coinfig.move 模块，用于存储链上配置，这可能是 Rooch 平台特有的管理方式。而 SUI/APTOS Move 则没有提及类似的专用配置模块。
Rooch Move 中的 session_key.move 模块用于管确保安全，这是 Rooch 上独特的安全管理功能，而 SUI/APTOS Move 没有这一功能。
