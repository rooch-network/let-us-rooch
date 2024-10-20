module random::random {
    use moveos_std::timestamp::{Self, Timestamp};
    use moveos_std::object::{Self ,Object};
    use moveos_std::tx_context;
    use moveos_std::address;
    use moveos_std::signer;
    use moveos_std::bcs;
    use moveos_std::hash::sha3_256;
    use std::vector;

    public fun get_tx_hash(): vector<u8> {
        let tx_hash = tx_context::tx_hash();
        tx_hash
    }

    public fun u64_to_bytes(num: u64): vector<u8> {
        if (num == 0) {
            return b"0"
        };
        let bytes = vector::empty<u8>();
        while (num > 0) {
            let remainder = num % 10;
            num = num / 10;
            vector::push_back(&mut bytes, (remainder as u8) + 48);
        };
        vector::reverse(&mut bytes);
        bytes
    }

    public fun random_to_u64(bytes: vector<u8>): vector<u8>  {
        let len = vector::length(&bytes);

        let start_index = (len - 8);
        let selected_bytes = vector::empty<u8>();

        let i = 0;
        while (i < 8) {
            let byte = vector::borrow(&bytes, start_index + i);
            vector::push_back(&mut selected_bytes, *byte);
            i = i + 1;
        };
        vector::reverse(&mut selected_bytes);
        selected_bytes
    }

    public entry fun get_random(account: &signer, timestamp_obj: &Object<Timestamp>, max: u64) {
        let account_addr = signer::address_of(account);
        let timestamp = object::borrow(timestamp_obj);
        let now_seconds = timestamp::seconds(timestamp);
        let tx_hash = get_tx_hash();

        let random_vector = vector::empty<u8>();
        vector::append(&mut random_vector, address::to_bytes(&account_addr));
        vector::append(&mut random_vector, u64_to_bytes(now_seconds));
        vector::append(&mut random_vector, tx_hash);

        let temp1 = sha3_256(tx_hash);
        let tempnum = random_to_u64(temp1);

        let random_num_ex = bcs::to_u64(tempnum);
        let random_value = (random_num_ex % max);
        random_value
    }
}
