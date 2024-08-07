module nft_demo::nft{

    use std::string::{Self,String};
    use moveos_std::signer;
    use moveos_std::event;
    use moveos_std::object::{Self, Object};
    use moveos_std::object::ObjectID;
    use moveos_std::display;
    use moveos_std::account;
    use bitcoin_move::utxo::UTXO;
    use rooch_framework::gas_coin::GasCoin;
    
    const ENOTEXIST_UTXO:u64 = 991;
    const EINSUFFICIENT_GAS_COIN:u64=992;
    const MINT_GAS:u256 = 1000;

    struct NFT has key,store {
        name: String,
        creator: address,
    }

    struct MintEvent has drop,copy {
        name: String,
        sender :address,
        contract_owner :address,
        ntf_obj: ObjectID,
    }

    struct MintInternalEvent has drop,copy{
        exists: bool,
        gas_balance: u256,
    }

    fun init(){
        let ntf_display_obj = display::display<NFT>();
        display::set_value(ntf_display_obj,string::utf8(b"name"),string::utf8(b"{value.name}"));
        display::set_value(ntf_display_obj,string::utf8(b"owner"),string::utf8(b"{owner}"));
        display::set_value(ntf_display_obj,string::utf8(b"creator"),string::utf8(b"{value.creator}"));
        display::set_value(ntf_display_obj,string::utf8(b"uri"),string::utf8(b"https://base_url/{id}"));
    }

    fun mint_internal(s: &signer,name: String):Object<NFT>{

        // check account has utxo object , only has utxo can mint objc
        assert!(account::exists_resource<UTXO>(signer::address_of(s)),ENOTEXIST_UTXO);

        // check sufficient gas coin
        assert!(rooch_framework::gas_coin::balance(signer::address_of(s)) > MINT_GAS,EINSUFFICIENT_GAS_COIN);

        // transfer gas coin , from signer to contract owner
        rooch_framework::account_coin_store::transfer<GasCoin>(s,@nft_demo,MINT_GAS);

        event::emit(MintInternalEvent{
            exists:account::exists_resource<UTXO>(signer::address_of(s)),
            gas_balance:rooch_framework::gas_coin::balance(signer::address_of(s)),
        });

        let nft = NFT{name:name,creator:signer::address_of(s)};
        let nft_obj = object::new(nft);
        nft_obj
    }

    public entry fun mint_entry(s: &signer,name: String){

        let nft_obj =        mint_internal(s,name); 

        object::transfer(nft_obj,signer::address_of(s));

        event::emit(MintEvent{
            name:name,
            sender:signer::address_of(s),
            ntf_obj:object::id(&nft_obj),
            contract_owner:@nft_demo,
        });
    }
}