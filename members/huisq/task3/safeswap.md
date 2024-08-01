## Week 3 
### 任务：构思一个 Bitcoin 生态的应用或者游戏，可以利用 Rooch 提供的特性，写成文章或者 Github Issue

#### 应用名字：safeswap 
- 于利用Rooch特性的AMM swap 平台

#### 利用的Rooch特性：
##### 1. SessionKey
每个用户交互都是生成一个新的SessionKey，scope里确保只和相关contract address和module交互，并且设有过期日期以避免账号长期给予权限而有风险。

```move
struct SessionKey has store,copy,drop {
    /// 应用名称
    app_name: std::string::String,
    /// 应用网站 URL
    app_url: std::string::String,
    /// 会话密钥的认证密钥，也作为会话密钥的 ID
    authentication_key: vector<u8>,
    /// 会话密钥的作用域
    scopes: vector<SessionScope>,
    /// 会话密钥的创建时间，当前时间戳（秒）
    create_time: u64,
    /// 会话密钥的最后活动时间，以秒为单位
    last_active_time: u64,
    /// 会话密钥的最长非活动时间段，以秒为单位
    /// 如果会话密钥在此期间未激活，它将过期
    /// 如果 max_inactive_interval 为 0，则会话密钥永不过期
    max_inactive_interval: u64,
}
```
这个library能够让各个应用和合约方便的嵌入设计中，精简的保护用户安全。
library中一些函数也可简单便捷调用和管理需要的资料：

```
public fun is_expired_session_key(account_address: address, authentication_key: vector<u8>): bool

public(friend) fun in_session_scope(session_key: &session_key::SessionKey): bool

public(friend) fun active_session_key(authentication_key: vector<u8>)
```

##### 2. Move on Rooch
Rooch 在 Diem 开发的 Move 之上延伸出 Move on Rooch保有 Move的特点如合约内置安全性，且加上其他独有的特点。
- 安全性：Move 具有内置的安全特性，并支持资源稀缺性，这使得它非常适合区块链应用程序，在这些应用程序中，资产和应用程序逻辑紧密相关。
- #[private_generics(T)] 是一个函数注解，确保带有此注解的函数只能在定义 T 的模块内被调用。在safeswap的合约设计里也会将个别函数套上此注解以确保不被外部合约攻击。