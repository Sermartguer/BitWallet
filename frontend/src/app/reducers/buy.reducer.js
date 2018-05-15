import {
    GET_ORDERS,
    GET_USER_ORDERS,
    SAVE_ORDER,
    OVERVIEW_ORDER_BALANCE
} from '../actions/types';

export const reducer = (state = {}, action) => {

    switch (action.type) {
        case GET_ORDERS:
            return { ...state, buy: action.payload}
        case GET_USER_ORDERS:
            return {...state, userOrder: action.payload}
        case SAVE_ORDER:
            return {...state, newOrder: action.payload}
        case OVERVIEW_ORDER_BALANCE:
            let currenc = 'orders'+action.currency;
            return { ... state, [currenc]:action.payload}
        default:
            return state;
    }
};