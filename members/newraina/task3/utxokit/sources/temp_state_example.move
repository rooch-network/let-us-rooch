module utxokit::temp_state_example {
    use std::option;

    use bitcoin_move::bitcoin;
    use bitcoin_move::types;
    use bitcoin_move::utxo::{Self, UTXO};
    use moveos_std::event;
    use moveos_std::object::Object;
    use moveos_std::timestamp;

    const ErrorBitcoinClientError: u64 = 0;
    const ErrorAlreadyStaked: u64 = 1;
    const ErrorDataNotMatch: u64 = 2;

    struct UTXOTempState has store, drop {
        start_time: u64,
        height: u64,
    }

    struct TempStatePassedCheckEvent has copy, store, drop {
        start_time: u64,
        height: u64,
    }

    public fun get_block_height(): u64 {
        let height_hash = bitcoin::get_latest_block();
        assert!(option::is_some(&height_hash), ErrorBitcoinClientError);
        let (height, _hash) = types::unpack_block_height_hash(option::destroy_some(height_hash));
        height
    }

    public fun add_temp_state(utxo: &mut Object<UTXO>) {
        assert!(!utxo::contains_temp_state<UTXOTempState>(utxo), ErrorAlreadyStaked);
        let height = get_block_height();
        let now = timestamp::now_seconds();
        let temp_state = UTXOTempState { start_time: now, height };
        utxo::add_temp_state(utxo, temp_state);
    }

    public fun get_temp_state(utxo: &Object<UTXO>): UTXOTempState {
        let temp_state = utxo::borrow_temp_state<UTXOTempState>(utxo);
        let state = UTXOTempState { start_time: temp_state.start_time, height: temp_state.height };

        state
    }

    public entry fun main(utxo: &mut Object<UTXO>) {
        add_temp_state(utxo);
        let temp_state = get_temp_state(utxo);
        assert!(temp_state.start_time == timestamp::now_seconds(), ErrorDataNotMatch);

        event::emit(TempStatePassedCheckEvent { start_time: temp_state.start_time, height: temp_state.height });
    }
}
