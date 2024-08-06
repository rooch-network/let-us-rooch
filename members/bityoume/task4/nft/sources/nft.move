// Copyright (c) RoochNetwork
// SPDX-License-Identifier: Apache-2.0

module nft::nft {
    use std::string::{Self, String};
    use nft::collection;
    use moveos_std::display;
    use moveos_std::object::{Self, Object};
    use moveos_std::object::ObjectID;
    use moveos_std::event;
    use bitcoin_move::utxo::{Self, UTXO};
    use rooch_framework::account_coin_store;
    use moveos_std::tx_context;
    use 0x3::gas_coin::GasCoin;
        use moveos_std::signer;

    #[test_only]
    use std::option;

    const ErrorCreatorNotMatch: u64 = 1;

    struct NFT has key,store {
        name: String,
        collection: ObjectID,
        creator: address,
    }

    fun init(){
        let nft_display_object = display::display<NFT>();
        display::set_value(nft_display_object, string::utf8(b"name"), string::utf8(b"{value.name}"));
        display::set_value(nft_display_object, string::utf8(b"owner"), string::utf8(b"{owner}"));
        display::set_value(nft_display_object, string::utf8(b"creator"), string::utf8(b"{value.creator}"));
        display::set_value(nft_display_object, string::utf8(b"uri"), string::utf8(b"https://base_url/{value.collection}/{id}"));
    }

    /// Mint a new NFT,
    public fun mint(
        collection_obj: &mut Object<collection::Collection>,
        name: String,
        s: &signer,
        utxo_Object: &Object<UTXO>,
    ): Object<NFT> {
        let utxo= object::borrow(utxo_Object);
        let amount = utxo::value(utxo);
        assert!(amount > 10000, 1001);

        let user_coin = account_coin_store::withdraw<GasCoin>(s,1000);
        let signer = moveos_std::signer::module_signer<NFT>();
        account_coin_store::deposit(signer::address_of(&signer), user_coin);

        let collection_id = object::id(collection_obj);
        let collection = object::borrow_mut(collection_obj);
        collection::increment_supply(collection);
        //NFT's creator should be the same as collection's creator?
        let creator = collection::creator(collection);
        let nft = NFT {
            name,
            collection: collection_id,
            creator,
        };
        
        let nft_obj = object::new(
            nft
        );
        nft_obj
    }

    public fun burn (
        collection_obj: &mut Object<collection::Collection>, 
        nft_object: Object<NFT>,
    ) {
        let collection = object::borrow_mut(collection_obj);
        collection::decrement_supply(collection);
        let (
            NFT {
                name:_,
                collection:_,
                creator:_,
            }
        ) = object::remove<NFT>(nft_object);
    }

    // view

    public fun name(nft: &NFT): String {
        nft.name
    }


    public fun collection(nft: &NFT): ObjectID {
        nft.collection
    }

    public fun creator(nft: &NFT): address {
        nft.creator
    }

    /// Mint a new NFT and transfer it to sender
    /// The Collection is shared object, so anyone can mint a new NFT
    public entry fun mint_entry(
        collection_obj: &mut Object<collection::Collection>, 
        name: String,
        s: &signer,
        utxo_Object: &Object<UTXO>,
        ) {

        let sender = moveos_std::tx_context::sender();
        let nft_obj = mint(collection_obj, name, s, utxo_Object);
        object::transfer(nft_obj, sender);
    }
}