module mintnft::collection {
    use std::option::{Self, Option};
    use std::string::{Self, String};

    use moveos_std::display;
    use moveos_std::event;
    use moveos_std::object;
    use moveos_std::object::ObjectID;
    use moveos_std::tx_context;

    friend mintnft::nft;

    const COLLECTION_NAME: vector<u8> = b"Snow NFT Collection";
    const COLLECTION_DESCRIPTION: vector<u8> = b"This is Snow NFT Collection";
    const COLLECTION_MAX_SUPPLY: u64 = 1000000;

    const ErrorCollectionMaximumSupply: u64 = 1;

    struct Collection has key {
        name: String,
        creator: address,
        supply: Supply,
        description: String,
    }

    struct Supply has store {
        current: u64,
        maximum: Option<u64>,
    }

    struct CreateCollectionEvent has drop, copy {
        object_id: ObjectID,
        name: String,
        creator: address,
        maximum: Option<u64>,
        description: String,
    }

    fun init() {
        let collection_display_obj = display::display<Collection>();
        display::set_value(
            collection_display_obj, string::utf8(b"name"), string::utf8(b"{name}")
        );
        display::set_value(
            collection_display_obj,
            string::utf8(b"description"),
            string::utf8(b"{description}"),
        );
        display::set_value(
            collection_display_obj, string::utf8(b"creator"), string::utf8(b"{creator}")
        );
        display::set_value(
            collection_display_obj, string::utf8(b"supply"), string::utf8(b"{supply}")
        );
    }

    public fun create_collection(): ObjectID {
        let collection = Collection {
            name: string::utf8(COLLECTION_NAME),
            creator: tx_context::sender(),
            supply: Supply { current: 0, maximum: option::some(COLLECTION_MAX_SUPPLY), },
            description: string::utf8(COLLECTION_DESCRIPTION),
        };

        let collection_obj = object::new(collection);
        let collection_id = object::id(&collection_obj);
        event::emit(
            CreateCollectionEvent {
                object_id: collection_id,
                name: string::utf8(COLLECTION_NAME),
                creator: tx_context::sender(),
                maximum: option::some(COLLECTION_MAX_SUPPLY),
                description: string::utf8(COLLECTION_DESCRIPTION),
            },
        );
        object::to_shared(collection_obj);
        collection_id
    }

    public entry fun create_collection_entry() {
        create_collection();
    }

    public(friend) fun increment_supply(collection: &mut Collection): Option<u64> {
        collection.supply.current = collection.supply.current + 1;
        if (option::is_some(&collection.supply.maximum)) {
            assert!(
                collection.supply.current <= *option::borrow(&collection.supply.maximum),
                ErrorCollectionMaximumSupply,
            );
            option::some(collection.supply.current)
        } else {
            option::none<u64>()
        }
    }

    public(friend) fun decrement_supply(collection: &mut Collection): Option<u64> {
        collection.supply.current = collection.supply.current - 1;
        if (option::is_some(&collection.supply.maximum)) {
            option::some(collection.supply.current)
        } else {
            option::none<u64>()
        }
    }

    // view
    public fun name(collection: &Collection): String {
        collection.name
    }

    public fun creator(collection: &Collection): address {
        collection.creator
    }

    public fun current_supply(collection: &Collection): u64 {
        collection.supply.current
    }

    public fun maximum_supply(collection: &Collection): Option<u64> {
        collection.supply.maximum
    }
}
