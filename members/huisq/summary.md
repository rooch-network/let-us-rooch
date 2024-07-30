# 学习成果

## Week 2:
### 1. 分析Move生态里的Object Model

| **Feature**                 | **Aptos**                                                               | **Sui**                                                              | **Rooch**                                                           |
|-----------------------------|-------------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|
| **Ownership Model**         | Account-based                                                           | Object-based                                                        | Object-based with ownership types                                   |
| **Object Representation**   | Assets managed within accounts as resources                             | Each asset is a unique object with a specific ID                    | Each asset is an object with specific ownership properties          |
| **Ownership Management**    | Managed within accounts and controlled via smart contracts              | Direct ownership recorded in object's metadata                      | Objects can be system-owned or user-owned, with specific ownership flags |
| **Object Types**            | Resources and tokens within accounts                                    | Unique objects with explicit ownership metadata                     | SystemOwnedObject (owner as 0x0), UserOwnedObject (owner non-equal to 0x0) |
| **Object Flags**            | Not applicable                                                          | Address-Owned Objects, Immutable Objects, Shared Objects, Wrapped Objects | Normal (default for UserOwnedObject), SharedObject, FrozenObject    |
| **Transfer Mechanism**      | Transactions between accounts                                           | Transactions update object's metadata to reflect new owner          | Transactions update 'owner' 'ObjectEntity' to reflect new owner     |
| **Programming Paradigm**    | Move (resource-oriented)                                                | Move (object-centric)                                               | Move (object-centric with scope and ownership type management)      |
| **Granularity of Control**  | At the account level                                                    | At the object level                                                 | At the object level, with additional control through ownership types and flags |
| **Shared Ownership**        | Managed through smart contracts and multisig accounts                   | SharedObject allows multiple accounts to access and mutate          | SharedObject flag allows everyone to get a &mut reference           |
| **Restricted Access**       | Managed through account permissions and smart contract logic            | Immutable object - can't be mutated, transferred, or deleted        | FrozenObject flag restricts &mut reference access                   |
| **Storage Abstraction**     | Hierarchical key-value storage model; each account maintains separate storage for its data and resources | Flexible storage model; each object encapsulates its own data, allowing efficient access and modification | Hybrid storage model; combines on-chain and off-chain storage, with on-chain metadata pointing to off-chain data for scalability and efficiency |

### 2. 常用的Object-related Method列表对应表

| **Function/Purpose**        | **Aptos**                                                                       | **Sui**                                           | **Rooch**                                                                    |
|-----------------------------|----------------------------------------------------------------------------------|--------------------------------------------------|-------------------------------------------------------------------------------|
| **Create Object**           | `public fun create_object(owner_address: address): ConstructorRef`               | `public fun new(ctx: &mut TxContext): UID`        | `#[private_generics(T)] public fun new<T: key>(value: T): Object<T>`          |
| **Named Object**            | `public fun create_named_object(creator: &signer, seed: vector<u8>): object::ConstructorRef` | Not specified                                    | `#[private_generics(T)] object::new_named_object<T: key>(T): Object<T>`       |
| **Account Named Object**    | Not specified                                                                    | Not specified                                    | `#[private_generics(T)] object::new_account_named_object<T: key>(address, T): Object<T>` |
| **Borrow Object**           | Not specified                                                                    | `public fun borrow<T: key>(self: &Object<T>): &T` | `public fun borrow_object<T: key>(object_id: ObjectID): &Object<T>`           |
| **Borrow Mutable Object**   | Not specified                                                                    | `public fun borrow_mut<Name: copy + drop + store, Value: key + store>(object: &mut UID, name: Name): &mut Value` | `public fun borrow_mut_object<T: key>(owner: &signer, object_id: ObjectID): &mut Object<T>` |
| **Take Object**             | `public fun object_from_constructor_ref<T: key>(ref: &ConstructorRef): Object<T>` | Not specified                                    | `public fun take_object<T: key + store>(owner: &signer, object_id: ObjectID): Object<T>` |
| **Borrow Inner Value**      | Not specified                                                                    | `public fun borrow<Name: copy + drop + store, Value: key + store>(object: &UID, name: Name): &Value` | `public fun borrow<T: key>(self: &Object<T>): &T`                             |
| **Borrow Mutable Inner Value** | Not specified                                                                | `public fun borrow_mut<Name: copy + drop + store, Value: key + store>(object: &mut UID, name: Name): &mut Value` | `public fun borrow_mut<T: key>(self: &mut Object<T>): &mut T`                 |
| **Transfer Object**         | `public fun transfer_with_ref(ref: LinearTransferRef, to: address) acquires ObjectCore, TombStone` | `public fun transfer<T: key>(obj: T, recipient: address)` | `public fun transfer<T: key + store>(self: Object<T>, new_owner: address)`    |
| **Public Transfer Object**  | Not specified                                                                    | `public fun public_transfer<T: key + store>(obj: T, recipient: address)` | `public fun take_object<T: key + store>(owner: &signer, object_id: ObjectID): Object<T>` |
| **Share Object**            | Not specified                                                                    | `public fun share_object<T: key>(obj: T)`         | Not specified                                                                |
| **Delete Object**           | `public fun delete(ref: DeleteRef) acquires Untransferable, ObjectCore`          | `public fun delete(id: UID)`                      | `#[private_generics(T)] public fun remove<T: key>(self: Object<T>): T`        |

## Week 3:
### 分析SessionKey
Rooch 的 SessionKey设计是move生态独有的一个设计。
它是一个临时密钥，方便用户在设定的时间、合约和模块范围内和应用产生链上交互。相当于在web2场境内，使用网银app一段时间内不活跃便会被自动登出。SessionKey能自动限制用户给予应用权限的时间范围，不需要用户时不时手动revoke。

#### 功能
1. **Session Key 生成**：在用户第一次与 Rooch 应用程序交互时，系统会生成一个 Session Key。这个密钥包含了一些关键信息，比如应用的名称、网址、密钥的创建时间和最大不活动时间间隔等。

2. **过期机制**：Session Key 具有过期时间。如果用户在设定的不活动时间内没有进行交互，该 Session Key 将失效。这一机制确保了 Session Key 的安全性，并防止长期不使用带来的安全隐患。

3. **自动签名**：在session期间，只要 Session Key 没有过期，用户不需要再次调用钱包进行签名。Rooch 系统会自动在每一步操作时为链上的交易签名。这极大地简化了用户操作，提升了用户体验。

### 优势
1. **简化签名流程**：Session Key 减少了频繁调用钱包进行签名的步骤。以很多链上游戏为例，每一步都需要发送交易，如果每次都需要调用钱包签名，会极大地降低用户体验。而使用 Session Key，只需在游戏开始时签名一次，后续步骤中系统会自动签名。

2. **提高安全性**：传统方法中，有些钱包使用委托代理形式预付 Gas，当预算用完时需要一次性上传签名交易。这种方式存在金融风险和游戏作弊风险。而 Session Key 机制避免了这些风险，因为所有的操作都由 Rooch 系统自动签名，且密钥在一定时间内有效。

3. **权限管理**：在首次签名时，可以定义 Session Key 的权限，决定不同账户地址可以使用哪些功能。这种细粒度的权限控制提高了系统的安全性和灵活性。

4. **用户体验**：Session Key 提供了 Web2 级别的流畅用户体验。用户只需在开始时进行一次签名，后续所有操作系统自动处理，使 Web3 应用的使用体验更加接近 Web2 应用。

### 实现细节
1. **密钥管理**：Session Key 的生成、存储和验证都由 Rooch 系统管理。密钥存储在用户设备上，并通过安全协议进行传输和验证。

2. **过期检查**：每次操作时，系统会检查 Session Key 是否过期。如果过期，用户需要重新生成 Session Key 并进行签名。

3. **自动签名服务**：Rooch 系统内置自动签名服务，确保在会话期间的每一步操作都能自动完成签名并提交交易。

### 使用场景
1. **游戏应用**：在游戏中，每一步操作都需要与区块链交互。Session Key 简化了操作，使用户体验更加流畅。
2. **电商平台**：在电商平台上，用户需要频繁签名进行支付、确认等操作。Session Key 可以减少签名步骤，提高操作效率。
3. **社交网络**：在区块链社交网络中，用户发布、评论等操作都需要签名。Session Key 可以提高用户活跃度，减少操作复杂性。
3. **DeFi**：在DeFi应用中，使用 Session Key 可以使用户不用长期赋予应用权限，一定程度上提高用户安全性和体验。

总的来说，Rooch 的 Session Key 通过简化签名流程、提高安全性和优化用户体验，为 Web3 应用提供了一种高效、安全的解决方案。