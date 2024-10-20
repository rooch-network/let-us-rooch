.PHONY: build publish inscription open object

egg = colored_egg10

build:
	rooch move build  --named-addresses colored_egg=default

publish:
	rooch move publish  --named-addresses colored_egg=default

inscription:
	rooch rpc request --method \
 	btc_queryInscriptions \
	--params '[{"owner":"tb1p588cfjvg4d6hchnuslzx6zhchzq3ht4csqmfeafhvqn5g680c74s89day6"}, null, "50", true]'

open:
	rooch move run --function default::$(egg)::open \
	--args object_id:0xd23c49fb9a742624498390a5b59c90e968c11d52b6437278a0e2eb9b82824103af601544b8188115132807fd5b5459a8333d2069031776d473f072b568a13f92 \
	--args object_id:0x62b99ecdf459afd575a565da457d78d11cf4489d5d8c3d79f1334e490359202a

switch test:
	rooch env switch --alias test

object:
	rooch rpc request --method rooch_queryObjectStates \
	--params '[{"object_type":"0xdb2e764a11715cc8cd056adf867053ba022a5d06116bfcde751b8b96a7e5f978::$(egg)::CoinInfoHolder"}, null, "2", {"descending": true,"showDisplay":false}]'

is_open:
	rooch move view --function default::$(egg)::is_open --args object_id:0xd23c49fb9a742624498390a5b59c90e968c11d52b6437278a0e2eb9b82824103f633be9c1690053ed7c540d00acce3181c8056967e51ec8158011957dc84f798

is_opens:
	rooch move view --function default::$(egg)::is_opens \
 	--args 'vector<object_id>:0xd23c49fb9a742624498390a5b59c90e968c11d52b6437278a0e2eb9b8282410397d6bb2a6c46e9d4915281389aa8b8247abcb6d74758f12722abdc457fd84d34'

get_pool_capacity:
	rooch move view --function default::$(egg)::get_pool_capacity

get_publish_height:
	rooch move view --function default::$(egg)::get_publish_height

get_last_block_height:
	rooch move view --function default::$(egg)::get_last_block_height

get_coin_num:
	rooch move view --function default::$(egg)::get_coin_num \
	--args object_id:0xd23c49fb9a742624498390a5b59c90e968c11d52b6437278a0e2eb9b82824103f633be9c1690053ed7c540d00acce3181c8056967e51ec8158011957dc84f798

check_inscription_status:
	rooch move view --function default::$(egg)::check_inscription_status \
 	--args 'vector<object_id>:0xd23c49fb9a742624498390a5b59c90e968c11d52b6437278a0e2eb9b82824103f633be9c1690053ed7c540d00acce3181c8056967e51ec8158011957dc84f798,0xd23c49fb9a742624498390a5b59c90e968c11d52b6437278a0e2eb9b82824103f633be9c1690053ed7c540d00acce3181c8056967e51ec8158011957dc84f798'
