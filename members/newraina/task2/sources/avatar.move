module avatar::avatar {
    use std::signer;
    use std::string;
    use std::string::String;

    use moveos_std::event;
    use moveos_std::object;
    use moveos_std::object::ObjectID;

    struct Avatar has key, drop, store {
        url: String,
    }

    struct AvatarCreatedEvent has copy, store, drop {
        id: ObjectID,
    }

    fun init(owner: &signer) {
        create_default(owner);
    }

    public fun create(owner: &signer, url: String): ObjectID {
        let avatar = object::new(Avatar { url });
        let owner_address = signer::address_of(owner);
        let id = object::id(&avatar);

        let avatar_created_event = AvatarCreatedEvent { id };
        event::emit(avatar_created_event);

        object::transfer(avatar, owner_address);
        id
    }

    public entry fun create_default(owner: &signer) {
        create(owner, string::utf8(b"https://example.com/avatar.png"));
    }
}
