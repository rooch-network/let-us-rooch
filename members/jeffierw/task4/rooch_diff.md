列举出至少三个Rooch Move和 SUI/APTOS Move 的区别

1. 通过`#[data_struct]` 结构体注解，标记结构体为纯数据结构体，使其可以在合约中直接反序列化。
```move
module my_project::my_module {
    #[data_struct]
    struct MyData has copy, drop {
        value: u64,
        name: vector<u8>,
    }
}

// 直接反序列化 MyData 结构体
let data: MyData = moveos_std::bcs::from_bytes(bytes);
```

2. 通过`#[private_generics(T)]` 函数注解，保证了添加该注解的函数，只能在定义 T 的模块内调用。这个注解对开发基础合约库非常有用，它可以保证添加了该注解的函数不能被用户随意调用，只能通过上层合约封装的函数进行调用。
```move
module moveos_std::account{
    #[private_generics(T)]
   /// Borrow a mut resource from the account's storage
   /// This function equates to `borrow_global_mut<T>(address)` instruction in Move
   public fun borrow_mut_resource<T: key>(account: address): &mut T;
}
```

3. 内置了`bitcoin-move`框架，可以通过 bitcoin-move 框架来读取 Bitcoin 区块以及交易，利用区块和交易中携带的数据来进行编程。