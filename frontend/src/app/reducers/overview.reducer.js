import {
    GET_USER_BASIC,
    GET_COIN_PRICE,
    GET_TRANSACTIONS
} from '../actions/types';

export const reducer = (state = {}, action) => {
    switch (action.type) {        
        case GET_USER_BASIC:
            return { ...state, overview: action.payload}
        case GET_COIN_PRICE:
            let currency = action.payload.currency
            return { ...state, [currency]: action.payload}
        case GET_TRANSACTIONS:
            return { ...state, transactions:action.payload}
        default:
            return state;
    }
};