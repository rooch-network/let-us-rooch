module mint_nft::mint_nft {

    use std::string;
    use bitcoin_move::utxo;
    use bitcoin_move::utxo::UTXO;
    use moveos_std::tx_context::sender;
    use moveos_std::object;
    use moveos_std::object::{Object, to_shared};
    use moveos_std::display;

    struct NFT has key, store{
        index: u64,
    }

    struct Global has key, store{
        index: u64
    }

    struct MintTemp has store, drop {}

    const ErrorAlreadyMint: u64 = 1;

    fun init() {
        let nft_display_object = display::display<NFT>();
        display::set_value(nft_display_object, string::utf8(b"name"), string::utf8(b"ROOCH NFT#{index}"));
        display::set_value(nft_display_object, string::utf8(b"description"), string::utf8(b"This is Move on Rooch NFT"));
        display::set_value(nft_display_object, string::utf8(b"image_url"), string::utf8(b"https://rooch.network/logo/rooch_black_combine.svg"));
        let global_obj = object::new_named_object(Global {
            index: 0
        });
        to_shared(global_obj);
    }

    public fun mint(
        btc: &mut Object<UTXO>,
        global_obj: &mut Object<Global>
    ) {
        assert!(!utxo::contains_temp_state<MintTemp>(btc), ErrorAlreadyMint);
        let global = object::borrow_mut(global_obj);
        let nft_obj = object::new(
            NFT{
                index: global.index
            }
        );
        global.index = global.index + 1;
        object::transfer(nft_obj, sender());
    }
}
