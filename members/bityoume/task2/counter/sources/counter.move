module counter::counter {

   use moveos_std::account;

   struct Counter has key {
      value:u64,
   }

   fun init() {
      let signer = moveos_std::signer::module_signer<Counter>();
      account::move_resource_to(&signer, Counter { value: 0 });
   }

   public fun increase_() {
      let counter = account::borrow_mut_resource<Counter>(@counter);
      counter.value = counter.value + 1;
   }

   public entry fun increase() {
      Self::increase_()
   }

   public fun value(): u64 {
      let counter = account::borrow_resource<Counter>(@counter);
      counter.value
   }
}
