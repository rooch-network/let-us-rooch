module btc_locker::btc_nft {

    use std::string::{Self, String};
    use btc_locker::collection;
    use btc_locker::holder_coin::HDC;
    use moveos_std::display;
    use moveos_std::object::{Self, Object};
    use moveos_std::object::ObjectID;
    use rooch_framework::coin_store;
    use moveos_std::account;
    use rooch_framework::account_coin_store;

    #[test_only]
    use std::option;

    const ErrorCreatorNotMatch: u64 = 1;

    struct NFT has key,store {
        name: String,
        collection: ObjectID,
        creator: address,
    }

    struct Global has key, store {
        coin_store: object::Object<coin_store::CoinStore<HDC>>
    }

    fun init(){
        let nft_display_object = display::display<NFT>();
        display::set_value(nft_display_object, string::utf8(b"name"), string::utf8(b"{value.name}"));
        display::set_value(nft_display_object, string::utf8(b"owner"), string::utf8(b"{owner}"));
        display::set_value(nft_display_object, string::utf8(b"creator"), string::utf8(b"{value.creator}"));
        display::set_value(nft_display_object, string::utf8(b"uri"), string::utf8(b"https://base_url/{value.collection}/{id}"));

        let coin_store = coin_store::create_coin_store<HDC>();
        let signer = moveos_std::signer::module_signer<NFT>();
        account::move_resource_to(&signer, Global{
            coin_store,
        });
    }


    /// Mint a new NFT,
    public fun mint(
        s: &signer,
        collection_obj: &mut Object<collection::Collection>,
        name: String,
    ): Object<NFT> {
    
        let  mint_gas = 1000;
        let g = account::borrow_mut_resource<Global>(@btc_locker);
        let user_coin = account_coin_store::withdraw<HDC>(s,mint_gas);
        coin_store::deposit<HDC>(&mut g.coin_store, user_coin);

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
    public entry fun mint_entry(collection_obj: &mut Object<collection::Collection>, name: String,signer: &signer) {
        let sender = moveos_std::tx_context::sender();
        let nft_obj = mint(signer,collection_obj, name);
        object::transfer(nft_obj, sender);
    }

     /// Update the base uri of the NFT
    /// the Collection is shared object, so we need to check the creator of collection, only the creator of collection can update the base uri
    public entry fun update_base_uri(collection_obj: &Object<collection::Collection>, new_base_uri: String){
        let sender_address = moveos_std::tx_context::sender();
        let collection = object::borrow(collection_obj);
        assert!(collection::creator(collection) == sender_address, ErrorCreatorNotMatch);
        let nft_display_obj = display::display<NFT>();
        string::append(&mut new_base_uri, string::utf8(b"{ collection }/{ metadata.id }"));
        display::set_value(nft_display_obj, string::utf8(b"uri"), new_base_uri);
    }

}