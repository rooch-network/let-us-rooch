import {Args, getRoochNodeUrl, RoochClient,} from '@roochnetwork/rooch-sdk'

const contractAddr = import.meta.env.VITE_CONTRACT_ADDR + "::";
const coinType = `${contractAddr}EGGS`

// create a client connected to testnet
export const client = new RoochClient({
    url: getRoochNodeUrl('testnet'),
})

export const getBalance = async (address: string) => {
    let balance = "0";
    (await client.getBalances({owner: address,})).data.forEach((e) => {
        if (e.coin_type === coinType) {
            balance = e.balance
        }
    });
    return balance;
}

export const getPoolCapacity = async () => {
    return (await client.executeViewFunction(
        {
            target: `${contractAddr}get_pool_capacity`
        }
    )).return_values![0].decoded_value.toString()
}

export const getLastBlockHeight = async () => {
    return Number((await client.executeViewFunction(
        {
            target: `${contractAddr}get_last_block_height`
        }
    )).return_values![0].decoded_value)
}

export const checkInscriptionsState = async (ins: string[]): Promise<any> => {
    const data = await client.executeViewFunction(
        {
            args: [Args.vec('objectId', ins)],
            target: `${contractAddr}check_inscription_status`
        }
    )

    return data.return_values![0].decoded_value
}

