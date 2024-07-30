module hello_simon::hello_simon {
    use moveos_std::account;
    use std::string;
    struct HelloMessage has key {
        text: string::String
    }
    entry fun say_hello(owner: &signer) {
        let hello = HelloMessage { text: string::utf8(b"Hello Simon!") };
        account::move_resource_to(owner, hello);
    }
}