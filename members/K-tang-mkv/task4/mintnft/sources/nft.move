module mintnft::nft {
    use std::string::{Self, String};

    use moveos_std::display;
    use moveos_std::object::ObjectID;
    use moveos_std::object::{Self, Object};
    use moveos_std::signer;
    use mintnft::coin;
    use mintnft::collection;

    const REQUIRED_MINT_FEE_COIN_AMOUNT: u256 = 1000;

    const ErrorCreatorNotMatch: u64 = 1;

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
    }

    public fun mint(
        account: &signer,
        collection_obj: &mut Object<collection::Collection>,
        name: String,
    ): Object<NFT> {
        let account_address = signer::address_of(account);
        coin::deposit_to_treasury(account, REQUIRED_MINT_FEE_COIN_AMOUNT);

        let collection_id = object::id(collection_obj);
        let collection = object::borrow_mut(collection_obj);
        collection::increment_supply(collection);
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

    public fun burn(
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
    public entry fun mint_entry(sender: &signer, collection_obj: &mut Object<collection::Collection>, name: String) {
        let sender_address = moveos_std::tx_context::sender();
        let nft_obj = mint(sender, collection_obj, name);
        object::transfer(nft_obj, sender_address);
    }
}
