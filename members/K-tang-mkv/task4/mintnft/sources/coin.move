module mintnft::coin {

    use std::string;

    use moveos_std::object::{Self, Object, ObjectID};
    use moveos_std::signer;
    use rooch_framework::account_coin_store;
    use rooch_framework::coin;
    use rooch_framework::coin_store::{Self, CoinStore};

    const TOTAL_SUPPLY: u256 = 210_000_000_000u256;
    const DECIMALS: u8 = 10u8;

    struct GANTCoin has key, store {}

    struct Treasury has key, store {
        coin_store: Object<CoinStore<GANTCoin>>
    }


    fun init() {
        let coin_info_obj = coin::register_extend<GANTCoin>(
            string::utf8(b"GANT Coin"),
            string::utf8(b"GNT"),
            DECIMALS,
        );
        let coin = coin::mint_extend<GANTCoin>(&mut coin_info_obj, TOTAL_SUPPLY);
        object::to_frozen(coin_info_obj);
        let coin_store_obj = coin_store::create_coin_store<GANTCoin>();
        coin_store::deposit(&mut coin_store_obj, coin);
        let treasury_obj = object::new_named_object(Treasury { coin_store: coin_store_obj });
        object::to_shared(treasury_obj);
    }

    public fun get_treasury_obj_id(): ObjectID {
        object::named_object_id<Treasury>()
    }

    public entry fun faucet(account: &signer) {
        let account_addr = signer::address_of(account);
        let treasury_obj = object::borrow_mut_object_shared<Treasury>(get_treasury_obj_id());
        let treasury = object::borrow_mut(treasury_obj);
        let coin = coin_store::withdraw(&mut treasury.coin_store, 10000);
        account_coin_store::deposit(account_addr, coin);
    }

    public entry fun deposit_to_treasury(account: &signer, amount: u256) {
        let treasury_obj = object::borrow_mut_object_shared<Treasury>(get_treasury_obj_id());
        let treasury = object::borrow_mut(treasury_obj);
        let coin = account_coin_store::withdraw<GANTCoin>(account, amount);
        coin_store::deposit(&mut treasury.coin_store, coin);
    }
}
