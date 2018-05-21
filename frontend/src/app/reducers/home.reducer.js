import {
    UPDATE_BTC_PRICE,
    UPDATE_LTC_PRICE
} from '../actions/types';

export const reducer = (state = {}, action) => {
    switch (action.type) {        
        case UPDATE_BTC_PRICE:
            return { ...state, btcprice: action.payload}
        case UPDATE_LTC_PRICE:
            return { ...state, ltcprice: action.payload}
        default:
            return state;
    }
};