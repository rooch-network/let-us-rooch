## 场景说明：

对于艺术家，设计师可以在rooch链上发布自己的设计作品。
对有BTC的用户，在mint艺术家的作品，类似盲盒。记录写到链上。
进而形成资产共识，产生价值。

## 操作demo：

1. 将合约上链，rooch move publish。
2. 艺术家或者设计师将作品发布到链上
   rooch move run --function 0x71fc15189cdba4addaed7c865bb3478c2ef0979dc2dc4971c6574648f2518717::nft::create_nft_entry --sender-account 0x71fc15189cdba4addaed7c865bb3478c2ef0979dc2dc4971c6574648f2518717 --args 'string:nft1' --args 'string:/Users/roc/Code/rooch/let-us-rooch/members/heisenbergx626/demo/nft_imgs/1.jpeg'
   或者是存在IPFS上。
3. 用户自己去mint出自己的NFT
   rooch move run --function 0x71fc15189cdba4addaed7c865bb3478c2ef0979dc2dc4971c6574648f2518717::nft::mint_entry --sender-account 0x1c996f29a21020a9c67a0c50410e21e98b0855b75c974544e8eb128b83a37e5c
   类似开盲盒，不一定是哪个作品。