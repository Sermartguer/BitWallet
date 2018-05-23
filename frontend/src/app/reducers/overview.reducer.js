import {
    GET_USER_BASIC,
    GET_COIN_PRICE,
    GET_TRANSACTIONS,
    OVERVIEW_ORDER_STATUS
} from '../actions/types';

export const reducer = (state = {}, action) => {
    switch (action.type) {        
        case GET_USER_BASIC:
            let btcBalance = action.payload[0].amount;
            let dogeBalance = action.payload[1].amount;
            let ltcBalance = action.payload[2].amount;
            return { ...state, BtcBalance: btcBalance,LtcBalance:ltcBalance,DogeBalance:dogeBalance}
        case GET_COIN_PRICE:
            let currency = action.payload.currency
            return { ...state, [currency]: action.payload.data}
        case GET_TRANSACTIONS:
            let dogeTransaction = [];
            let ltcTransaction = [];
            let btcTransaction = [];
            action.payload.forEach(transaction => {
                if(transaction.currency === 'DOGE'){
                    dogeTransaction.push({send:transaction.send_to,amount:transaction.amount})
                }else if(transaction.currency === 'LTC'){
                    ltcTransaction.push({send:transaction.send_to,amount:transaction.amount})
                }else if(transaction.currency === 'BTC'){
                    btcTransaction.push({send:transaction.send_to,amount:transaction.amount})
                }
            });
            return { ...state, transactions:action.payload, transactionDOGE:dogeTransaction,transactionBTC:btcTransaction,transactionLTC:ltcTransaction}
        case OVERVIEW_ORDER_STATUS:
            let currenc = 'order'+action.currency;
            console.log(action)
            return { ... state, [currenc]:action.payload}
        default:
            return state;
    }
};