module nft_demo::nft{

    use std::string::{Self,String};
    use std::vector;

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
    const EUNAUTHORITY:u64=993;
    const EMINTEND:u64=994;

    const MINT_GAS:u256 = 1000;

    struct NFT has key,store {
        name: String,
        creator: address,
        uri: String,
    }


    struct Collection has key{
        name: String,
        nfts: vector<ObjectID>,
    }

    struct MintEvent has drop,copy {
        name: String,
        sender :address,
        contract_owner :address,
        uri: String,
        ntf_obj: ObjectID,
    }

    struct MintInternalEvent has drop,copy{
        exists: bool,
        gas_balance: u256,
    }

    struct CreateNFTEvent has drop,copy {
        name: String,
        nft_obj_id:ObjectID,
        creator: address,
        uri:String,
    }


    fun init(owner: &signer){

        let ntf_display_obj = display::display<NFT>();
        display::set_value(ntf_display_obj,string::utf8(b"name"),string::utf8(b"{value.name}"));
        display::set_value(ntf_display_obj,string::utf8(b"owner"),string::utf8(b"{owner}"));
        display::set_value(ntf_display_obj,string::utf8(b"creator"),string::utf8(b"{value.creator}"));
        display::set_value(ntf_display_obj,string::utf8(b"uri"),string::utf8(b"https://base_url/{id}"));

        let cls = Collection{
            name:string::utf8(b"nfts"),
            nfts:vector::empty(),
        };

        account::move_resource_to(owner,cls);

    }

    fun mint_internal(s: &signer):Object<NFT>{

        let module_sign = moveos_std::signer::module_signer<NFT>();
        let module_sign_address = signer::address_of(&module_sign);

        // check account has utxo object , only has utxo can mint objc
        // assert!(account::exists_resource<UTXO>(signer::address_of(s)),ENOTEXIST_UTXO);

        // check sufficient gas coin
        assert!(rooch_framework::gas_coin::balance(signer::address_of(s)) > MINT_GAS,EINSUFFICIENT_GAS_COIN);

        // check is enough nft
        let cls = account::borrow_mut_resource<Collection>(module_sign_address); 
        assert!(vector::length(&cls.nfts)>0,EMINTEND);

        // transfer gas coin , from signer to contract owner
        rooch_framework::account_coin_store::transfer<GasCoin>(s,@nft_demo,MINT_GAS);

        let mint_obj_id:ObjectID = vector::pop_back(&mut cls.nfts);

        event::emit(MintInternalEvent{
            exists:account::exists_resource<UTXO>(signer::address_of(s)),
            gas_balance:rooch_framework::gas_coin::balance(signer::address_of(s)),
        });

        
        let nft_obj= object::take_object<NFT>(&module_sign, mint_obj_id);
        nft_obj
    }

    public entry fun create_nft_entry(s: &signer,name: String, uri: String){

        let module_sign = moveos_std::signer::module_signer<NFT>();
        let module_sign_address = signer::address_of(&module_sign);
        // create nft must be contract owner
        assert!(signer::address_of(s)==module_sign_address,EUNAUTHORITY);

        let nft = NFT{name:name,creator:signer::address_of(s),uri:uri};
        let nft_obj = object::new(nft);
        let id = object::id(&nft_obj);
        object::transfer(nft_obj,module_sign_address);
        
        let cls = account::borrow_mut_resource<Collection>(module_sign_address); 
        vector::push_back(&mut cls.nfts,id);

        event::emit(CreateNFTEvent{
            name:name,
            nft_obj_id:id,
            creator:module_sign_address,
            uri:uri,
        });

    }


    public entry fun mint_entry(s: &signer){

        let nft_obj = mint_internal(s); 
        let name = object::borrow<NFT>(&nft_obj).name;
        let uri =  object::borrow<NFT>(&nft_obj).uri;
        event::emit(MintEvent{
            name:name,
            sender:signer::address_of(s),
            ntf_obj:object::id(&nft_obj),
            uri:uri,
            contract_owner:@nft_demo,
        });

        object::transfer(nft_obj,signer::address_of(s));       
    }

    // // view 
    // public fun get_uri_by_nftid(id: &Object<NFT>):String{

    //     let module_sign = moveos_std::signer::module_signer<NFT>();
    //     let module_sign_address = signer::address_of(&module_sign);

    //     let cls = account::borrow_mut_resource<Collection>(module_sign_address); 

    // }
}