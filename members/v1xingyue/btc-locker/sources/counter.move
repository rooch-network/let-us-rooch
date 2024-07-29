// Copyright (c) RoochNetwork
// SPDX-License-Identifier: Apache-2.0

module btc_locker::counter {
    use moveos_std::account;
    use bitcoin_move::bitcoin;
    use bitcoin_move::types::{BlockHeightHash,unpack_block_height_hash};
    use std::option::{Self};

    const ErrorBitcoinClientError: u64 = 10;
    const ErrorBitcoinBlockError: u64 = 11;

    struct Counter has key {
        count_value: u64,
        last_block: u64
    }

    fun init() {
        let signer = moveos_std::signer::module_signer<Counter>();
        account::move_resource_to(&signer, Counter { count_value: 0,last_block:0 });
    }

    entry fun increase() {
        let latest_block = bitcoin::get_latest_block();
        assert!(option::is_some<BlockHeightHash>(&latest_block), ErrorBitcoinClientError);
        let (height,_hash) = unpack_block_height_hash(option::destroy_some(latest_block));
        let counter = account::borrow_mut_resource<Counter>(@btc_locker);
        assert!(height % 10 >= 2, ErrorBitcoinBlockError);
        counter.count_value = counter.count_value + height % 10;
        counter.last_block = height;
    }
}
