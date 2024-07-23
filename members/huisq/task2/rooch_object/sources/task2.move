module rooch_object::rooch_object {
    use moveos_std::object::{Self, Object, ObjectID};
    
    struct NewEvent has copy, drop{
        id: ObjectID,
        value: u64
    }

    struct CookieBox has key, store {
        value: u64 //number of cookies in it
    }

    public fun print(value: u64){
        let output = std::string::utf8(b"Created new cookie box with ");
        std::string::append(&mut output, moveos_std::string_utils::to_string_u64(value));
        std::string::append(&mut output, std::string::utf8(b" cookies in it."));
        std::debug::print(&output);
    }

    public fun new(value: u64): Object<CookieBox>{
        let obj = object::new(CookieBox{value});
        let id = object::id(&obj);
        moveos_std::event::emit(NewEvent{id, value});
        obj
    }

    public entry fun create_box(value: u64){
        let obj = new(value);
        object::transfer(obj, moveos_std::tx_context::sender());
        print(value);
    }

    #[test]
    fun test_one(){
        create_box(10);
    }
}