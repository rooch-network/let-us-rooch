module bitcoin_monsters::monsters {

    use std::vector;
    use moveos_std::object::{Self, Object};
    use moveos_std::timestamp;
    use bitcoin_move::ord::{Self, Inscription};

    const ErrorNotMonster: u64 = 0;
    const ErrorAlreadyTrained: u64 = 1;
    const ErrorNotTrained: u64 = 2;
    const ErrorTrainTooFrequently: u64 = 3;
    const ErrorMonsterExhausted: u64 = 4;
    const ErrorAlreadycreated: u64 = 5;

    const TRAIN_INTERVAL: u64 = 60 * 60 * 24; // 1 day
    const MAX_VARIETIES: u64 = 5;
    const MAX_HEALTH: u8 = 100;
    const MAX_LEVEL: u64 = 100;
    const Monster_egg_Rarity: u8 = 1;

    /// Basic attributes of pets
    struct Monster has store {
        variety: u64,
        level: u64,
        experience: u64,
        health: u8,
        last_training_time: u64,
        wins: u64,
        losses: u64,
        achievement: u64,
    }
    
    /// User operations on pets
    struct Actions has store, copy, drop {
        creation_time: u64,
        training_time: vector<u64>,
        task_completion: vector<u64>,
    }

    fun init() {}

    ///mint new monsters
    public entry fun mint_monster(monster_egg_Rarity: &mut Object<Inscription>) {
        let inscription = object::borrow(monster_egg_Rarity);
        ensure_monster_inscription(inscription);
        assert!(!ord::contains_permanent_state<Monster>(monster_egg_Rarity), ErrorAlreadycreated);
        let monster = Monster {
            variety: 5,
            level: 1,
            experience: 0,
            health: 100,
            last_training_time: timestamp::now_seconds(),
            wins: 0,
            losses: 0,
            achievement:0,
        };

        ord::add_permanent_state(monster_egg_Rarity, monster);

        let actions = Actions {
            creation_time: timestamp::now_seconds(),
            training_time: vector::empty(),
            task_completion: vector::empty(),
        };
        ord::add_temp_state(monster_egg_Rarity, actions);

    }

    /// train monster
public entry fun train_monster(monster: &mut Object<Inscription>) {
    let monster_data = ord::borrow_mut_permanent_state<Monster>(monster);
    let now = timestamp::now_seconds();
    assert!(now - monster_data.last_training_time >= TRAIN_INTERVAL, ErrorTrainTooFrequently);
    
    monster_data.experience = monster_data.experience + calculate_experience_gain(); 
    monster_data.level = calculate_level(monster_data.experience);
    monster_data.health = calculate_health(monster_data.health, now - monster_data.last_training_time);
    if (monster_data.health == 0) {
        return
    };
    monster_data.last_training_time = now;

    let actions = ord::borrow_mut_temp_state<Actions>(monster);
    vector::push_back(&mut actions.training_time, now);
}

/// battle
public entry fun battle(monster: &mut Object<Inscription>, opponent: &mut Object<Inscription>) {
    let monster_data = ord::borrow_mut_permanent_state<Monster>(monster);
    let opponent_data = ord::borrow_mut_permanent_state<Monster>(opponent);

    assert!(monster_data.health > 0 && opponent_data.health > 0, ErrorMonsterExhausted);

    let _now = timestamp::now_seconds();
   
    // Simulate the battle outcome (simple logic for demonstration)
    let outcome = simulate_battle(monster_data, opponent_data);
    if (outcome == 1) {
    monster_data.wins = monster_data.wins + 1;
    opponent_data.losses = opponent_data.losses + 1;
    } else if (outcome == 2) {
    monster_data.losses = monster_data.losses + 1;
    opponent_data.wins = opponent_data.wins + 1;
    };

    let _actions = ord::borrow_mut_temp_state<Actions>(monster);
}

public fun do_task(monster: &mut Object<Inscription>, task_id: u64): vector<u64> {
    let monster_data = ord::borrow_mut_permanent_state<Monster>(monster);
    assert!(monster_data.health > 0, ErrorMonsterExhausted);

    let reward = complete_task(task_id);

    let actions = ord::borrow_mut_temp_state<Actions>(monster);
    vector::push_back(&mut actions.task_completion, task_id);

    reward
}

public fun is_monster(_inscription: &Inscription): bool {
    // TODO: Parse the Inscription content and check if it is a valid Monster
    true
}

public fun calculate_experience_gain(): u64 {
    // TODO: Implement logic to calculate experience gain from training or battles
    100
}

public fun calculate_level(experience: u64): u64 {
    // calculate level
    experience / 100 + 1
}

public fun calculate_health(current_health: u8, _time_since_last_action: u64): u8 {
    // health recover
    let recovery_rate: u8 = 5; // 5 per hour
    let hours_passed: u8 = ((_time_since_last_action / 3600) as u8); // transfer time to hour
    // recovery health
    let new_health = current_health + hours_passed * recovery_rate;

    // limit to 100
    if (new_health > 100) {
        100
    } else {
        new_health
    }
}

public fun get_monster_permanent_state(monster: &Object<Inscription>): &Monster {
    ord::borrow_permanent_state<Monster>(monster)
}

public fun get_monster_temp_state(monster: &Object<Inscription>): &Actions {
    ord::borrow_temp_state<Actions>(monster)
}

fun ensure_monster_inscription(inscription: &Inscription) {
    assert!(is_monster(inscription), ErrorNotMonster);
}

fun simulate_battle(_monster: &Monster, _opponent: &Monster) : u8 {
    1 // 1 for monster win, 2 for opponent win, 0 for draw
}

fun complete_task(_task_id: u64) : vector<u64> {
    // finish tasks to get rewards
   let rewards: vector<u64> = vector::empty();

    //rewards rules
    if (_task_id == 1) {
        vector::push_back(&mut rewards, 100);
    } else if (_task_id == 2) {
        vector::push_back(&mut rewards, 200);
    } else {
       
    };

    rewards
}

#[test_only]
use std::option;

#[test_only]
use rooch_framework::genesis;

#[test]
fun test() {
    genesis::init_for_test();
    let inscription_obj = ord::new_inscription_object_for_test(
        @0x3232423,
        0,
        0,
        vector::empty(),
        option::none(),
        option::none(),
        vector::empty(),
        option::none(),
        vector::empty(),
        option::none(),
    );

    mint_monster(&mut inscription_obj);

    let  i = 0u8;
    loop {
        timestamp::fast_forward_seconds_for_test(TRAIN_INTERVAL);
        train_monster(&mut inscription_obj);
        i = i + 1;
        if (i == 10) break;
    };

    let rewards = do_task(&mut inscription_obj, 1);
    assert!(vector::length(&rewards) > 0, 1);

    let monster = ord::remove_permanent_state<Monster>(&mut inscription_obj);
    ord::destroy_permanent_area(&mut inscription_obj);

    // Assert that the pet attributes and operation results are consistent with expectations
    assert!(monster.level > 0, 1);
}
}