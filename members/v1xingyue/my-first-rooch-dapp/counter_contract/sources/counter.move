module quick_start_counter::quick_start_counter {

    use moveos_std::account;
    use std::signer;

    struct Counter has key {
        count_value: u64
    }
    
    entry fun mint(signer:&signer){
        account::move_resource_to(signer, Counter { count_value: 0 });
    }

    entry fun increase(signer:&signer) {
        let counter = account::borrow_mut_resource<Counter>(signer::address_of(signer));
        counter.count_value = counter.count_value + 1;
    }

    public fun value(addr:address): u64{
        let counter = account::borrow_mut_resource<Counter>(addr);
        counter.count_value
    }
}