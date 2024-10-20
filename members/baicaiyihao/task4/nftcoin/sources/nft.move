// Copyright (c) RoochNetwork
// SPDX-License-Identifier: Apache-2.0

module nftcoin::nft {
    use std::string::{Self, String};
    use moveos_std::signer;
    use nftcoin::collection;
    use moveos_std::display;
    use moveos_std::object::{Self, Object};
    use moveos_std::object::ObjectID;
    
    use rooch_framework::coin;
    use rooch_framework::coin_store::{Self, CoinStore};
    use rooch_framework::account_coin_store;
    #[test_only]
    use std::option;

    const ErrorCreatorNotMatch: u64 = 1;
    const ErrorValidateBalance: u64 = 2;

    const TOTAL_SUPPLY: u256 = 210_000_000_000u256;
    const DECIMALS: u8 = 1u8;

    // The `FSC` CoinType has `key` and `store` ability.
    // So `FSC` coin is public.
    struct FSC has key, store {}

    // construct the `FSC` coin and make it a global object that stored in `Treasury`.
    struct Treasury has key {
        coin_store: Object<CoinStore<FSC>>
    }

    struct NFT has key, store {
        name: String,
        collection: ObjectID,
        creator: address,
    }

    fun init() {
        let nft_display_object = display::display<NFT>();
        display::set_value(nft_display_object, string::utf8(b"name"), string::utf8(b"{value.name}"));
        display::set_value(nft_display_object, string::utf8(b"owner"), string::utf8(b"{owner}"));
        display::set_value(nft_display_object, string::utf8(b"creator"), string::utf8(b"{value.creator}"));
        display::set_value(nft_display_object, string::utf8(b"uri"), string::utf8(b"https://base_url/{value.collection}/{id}"));
        
        let coin_info_obj = coin::register_extend<FSC>(
            string::utf8(b"Fixed Supply Coin"),
            string::utf8(b"FSC"),
            DECIMALS,
        );
        
        // Mint the total supply of coins, and store it to the treasury
        let coin = coin::mint_extend<FSC>(&mut coin_info_obj, TOTAL_SUPPLY);
        
        // Frozen the CoinInfo object, so that no more coins can be minted
        object::to_frozen(coin_info_obj);
        
        let coin_store_obj = coin_store::create_coin_store<FSC>();
        coin_store::deposit(&mut coin_store_obj, coin);
        
        let treasury_obj = object::new_named_object(Treasury { coin_store: coin_store_obj });
        
        // Make the treasury object to shared, so anyone can get mutable Treasury object
        object::to_shared(treasury_obj);
    }

    /// Provide a faucet to give out coins to users
    /// In a real world scenario, the coins should be given out in the application business logic.
    public entry fun faucet(account: &signer, treasury_obj: &mut Object<Treasury>) {
        let account_addr = signer::address_of(account);
        let treasury = object::borrow_mut(treasury_obj);
        let coin = coin_store::withdraw(&mut treasury.coin_store, 10000);
        account_coin_store::deposit(account_addr, coin);
    }

    /// Mint a new NFT
    public fun mint(
        s: &signer,
        collection_obj: &mut Object<collection::Collection>,
        name: String,
        treasury_obj: &mut Object<Treasury>
    ): Object<NFT> {
        let mint_gas = 1000;
        let account_addr = signer::address_of(s);
        let user_coin = account_coin_store::withdraw<FSC>(s, mint_gas);
        
        // Deposit the withdrawn coin back to the treasury
        let treasury = object::borrow_mut(treasury_obj);
        coin_store::deposit(&mut treasury.coin_store, user_coin);
        
        let collection_id = object::id(collection_obj);
        let collection = object::borrow_mut(collection_obj);
        collection::increment_supply(collection);
        
        let creator = collection::creator(collection);
        let nft = NFT {
            name,
            collection: collection_id,
            creator,
        };
        
        // Create an Object<NFT> type
        let nft_obj = object::new(nft);

        // Return the Object<NFT>
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
                name: _,
                collection: _,
                creator: _,
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
        account: &signer,
        collection_obj: &mut Object<collection::Collection>,
        name: String,
        treasury_obj: &mut Object<Treasury>
    ) {
        let nft_obj = mint(account, collection_obj, name, treasury_obj);
        object::transfer(nft_obj, signer::address_of(account));
    }
}
