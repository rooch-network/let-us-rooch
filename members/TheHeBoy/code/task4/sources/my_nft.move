// Copyright (c) RoochNetwork
// SPDX-License-Identifier: Apache-2.0

module task4::task4{
    use std::signer::address_of;
    use std::string::{Self, String};
    use moveos_std::display;
    use moveos_std::object::{Self, Object};
    use moveos_std::event;
    use bitcoin_move::utxo::{Self, UTXO};
    use rooch_framework::account_coin_store;
    use moveos_std::tx_context;
    use 0x3::gas_coin::GasCoin;

    struct NFT has key,store {
        name: String,
        creator: address,
    }

    struct BtcInsufficient has copy, drop {
        amount: u64
    }

    struct Global has key, store {
        owner:address
    }

    fun init(){
        let nft_display_object = display::display<NFT>();
        display::set_value(nft_display_object, string::utf8(b"name"), string::utf8(b"{value.name}"));
        display::set_value(nft_display_object, string::utf8(b"owner"), string::utf8(b"{owner}"));
        display::set_value(nft_display_object, string::utf8(b"creator"), string::utf8(b"{value.creator}"));
        display::set_value(nft_display_object, string::utf8(b"uri"), string::utf8(b"https://base_url/{id}"));
    }

    /// Mint a new NFT,
    public entry fun mint(
        s: &signer,
        name: String,
        utxo_Object: &Object<UTXO>,
    ) {
        let utxo= object::borrow(utxo_Object);
        let amount = utxo::value(utxo);
        if (amount < 10000){
            event::emit<BtcInsufficient>(BtcInsufficient {
                amount,
            });
            return
        };

        let user_coin = account_coin_store::withdraw<GasCoin>(s,1000);
        let signer = moveos_std::signer::module_signer<NFT>();
        account_coin_store::deposit(address_of(&signer), user_coin);

        let creator = tx_context::sender();
        let nft = NFT {
            name,
            creator,
        };

        let nft_obj = object::new(
            nft,
        );

        object::transfer(nft_obj, creator);
    }
}