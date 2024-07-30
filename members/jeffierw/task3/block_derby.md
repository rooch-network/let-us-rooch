## Week 3 
### 任务：构思一个 Bitcoin 生态的应用或者游戏，可以利用 Rooch 提供的特性，写成文章或者 Github Issue

#### 应用名字：Bitcoin Block Derby
- 利用Rooch特性的竞猜新生成区块特征的游戏

#### 利用的Rooch特性：
使用`bitcoin::get_block_by_height`获取目标区块信息

结合比特币区块的特性和简单的赛马游戏机制
#### 游戏概念：
玩家可以在即将到来的比特币区块上下注，预测该区块的某些特征。这些特征被比喻为不同的"赛马"，每个赛马代表区块的一个特定属性。
#### 游戏机制：

    1.赛道设置：

    设置5个"赛道"，每个赛道代表区块的一个特征：

    区块哈希的最后一位（16种可能，分为高、中、低）
    交易数量（分为多、中、少）
    区块大小（分为大、中、小）
    挖矿难度变化（上升、不变、下降）
    特殊事件（如包含大额交易、特殊OP_RETURN数据等）


    2.下注：

    玩家可以在下一个区块生成前对一个或多个"赛道"进行下注。
    每个赛道可以选择一个或多个选项。


    3.结算：

    当目标区块生成后，系统自动结算所有下注。
    玩家猜中的赛道将获得奖励，奖励倍数根据选项的难度而定。


    4.奖励池：

    所有下注金额的一部分进入奖励池。
    完全猜中所有赛道的玩家可以瓜分奖励池。


    5.连续下注：

    玩家可以选择将他们的下注"延续"到下一个区块，获得额外奖励。


    6.社交元素：

    玩家可以创建或加入"赛马俱乐部"，分享他们的策略和胜利

```move
    struct Race has key, store {
        target_height: u64,
        bets: vector<Bet>,
        reward_pool: u64,
    }

    struct Bet has store {
        bettor: address,
        predictions: vector<u8>,
        amount: u64,
    }

    struct DerbyStats has key, store {
        total_races: u64,
        total_bets: u64,
        total_rewards: u64,
    }
```

```
public fun create_race(creator: &signer, target_height: u64)

public fun place_bet(bettor: &signer, race_id: ObjectID, predictions: vector<u8>, amount: u64): Race 

public fun settle_race(settler: &signer, race_id: ObjectID): Race, DerbyStats

fun calculate_results(header: &Header): vector<u8>

fun calculate_reward(bet: &Bet, results: &vector<u8>): u64
```