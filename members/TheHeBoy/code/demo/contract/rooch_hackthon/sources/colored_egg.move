module colored_egg::colored_egg10 {

    use std::vector;
    use std::option;
    use std::string;
    use moveos_std::account;
    use moveos_std::account::{move_resource_to};
    use rooch_framework::account_coin_store;
    use moveos_std::tx_context;
    use bitcoin_move::bitcoin;
    use bitcoin_move::ord;
    use bitcoin_move::types::unpack_block_height_hash;
    use bitcoin_move::ord::Inscription;
    use moveos_std::object::{Self, Object};
    use rooch_framework::coin;
    use rooch_framework::coin::{CoinInfo};

    const ErrorAlreadyOpened: u64 = 0;
    const ErrorCoinNumEqualZero: u64 = 1;
    const ErrorPoolCapacityEqualZero: u64 = 2;

    const DECIMALS: u8 = 0;
    const BlockCapacity: u64 = 100;

    struct EGGS has key, store {}

    struct State has key {
        publish_height: u64,
        open_num: u256,
    }

    struct CoinInfoHolder has key {
        coin_info: Object<CoinInfo<EGGS>>,
    }

    struct InscriptionState has store {
        height: u64,
        amount: u64,
    }

    struct StatusResult {
        is_open: bool,
        coin_num: u64,
    }

    fun init() {
        let coin_info_obj = coin::register_extend<EGGS>(
            string::utf8(b"EGGS"),
            string::utf8(b"EGGS"),
            DECIMALS,
        );
        let coin_info_holder_obj = object::new_named_object(CoinInfoHolder { coin_info: coin_info_obj });
        // Make the coin info holder object to shared, so anyone can get mutable CoinInfoHolder object
        object::to_shared(coin_info_holder_obj);

        let signer = moveos_std::signer::module_signer<State>();

        let height = get_last_block_height();
        move_resource_to(&signer, State {
            publish_height: height,
            open_num: 0,
        })
    }

    fun get_state(): &mut State {
        account::borrow_mut_resource<State>(@colored_egg)
    }

    // open egg
    public entry fun open(egg_obj: &mut Object<Inscription>, coin_info_holder_obj: &mut Object<CoinInfoHolder>) {
        assert!(!is_open(egg_obj), ErrorAlreadyOpened);

        let coin_num = (get_coin_num(egg_obj) as u256);
        assert!(coin_num > 0, ErrorCoinNumEqualZero);

        let pool_capacity = get_pool_capacity();
        assert!(pool_capacity > 0, ErrorPoolCapacityEqualZero);
        if ((coin_num as u256) > pool_capacity) {
            coin_num = pool_capacity;
        };

        let egg_inscription = object::borrow_mut(egg_obj);
        check_inscription(egg_inscription);

        let height = get_last_block_height();
        let inscription_state = InscriptionState {
            height,
            amount: (coin_num as u64)
        };
        ord::add_permanent_state(egg_obj, inscription_state);

        let state = get_state();
        state.open_num = state.open_num + coin_num;
        let coin_info_holder = object::borrow_mut(coin_info_holder_obj);
        let coin = coin::mint_extend(&mut coin_info_holder.coin_info, coin_num);
        let sender = tx_context::sender();
        account_coin_store::deposit(sender, coin);
    }

    public fun get_coin_num(ins: &Object<Inscription>): u64 {
        if (is_open(ins)) {
            let state = ord::borrow_permanent_state<InscriptionState>(ins);
            state.amount
        }else {
            let tx_height = option::borrow(&bitcoin::get_tx_height(ord::txid(object::borrow(ins))));
            get_last_block_height() - *tx_height
        }
    }

    // last block height
    public fun get_last_block_height(): u64 {
        let block = option::extract(&mut bitcoin::get_latest_block());
        let (height, _) = unpack_block_height_hash(block);
        height
    }

    // contract block height when contract publish
    public fun get_publish_height(): u64 {
        get_state().publish_height
    }

    // number of coinss that can be minted
    public fun get_pool_capacity(): u256 {
        let state = get_state();
        (((get_last_block_height() - state.publish_height) * BlockCapacity) as u256) - state.open_num
    }

    // whether it is opened or not
    public fun is_open(egg_obj: &Object<Inscription>): bool {
        ord::contains_permanent_state<InscriptionState>(egg_obj)
    }

    public fun is_opens(egg_objs: &vector<Object<Inscription>>): vector<bool> {
        let v = vector::empty<bool>();
        vector::for_each_ref(egg_objs, |m| {
            vector::push_back(&mut v, is_open(m));
        });
        v
    }

    public fun check_inscription_status(egg_objs: &vector<Object<Inscription>>): vector<StatusResult> {
        let v = vector::empty<StatusResult>();

        vector::for_each_ref(egg_objs, |item| {
            vector::push_back(&mut v, StatusResult {
                is_open: is_open(item),
                coin_num: get_coin_num(item),
            });
        });

        v
    }

    fun check_inscription(_: &Inscription) {
        //todo
    }
}

