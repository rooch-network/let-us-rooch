module task2::task2 {
    use moveos_std::object::{to_shared, Object};
    use moveos_std::object;
    struct Counter has key {
        count_value: u64
    }
    fun init() {
        let counter_obj = object::new_named_object(Counter{count_value: 0});
        to_shared(counter_obj)
    }
    entry fun increase(counter_obj: &mut Object<Counter>) {
        let counter = object::borrow_mut(counter_obj);
        counter.count_value = counter.count_value + 1;
    }
}
