import {
    GET_ORDERS,
    GET_USER_ORDERS,
    SAVE_ORDER
} from '../actions/types';

export const reducer = (state = {}, action) => {

    switch (action.type) {
        case GET_ORDERS:
        console.log(action.payload)
            return { ...state, buy: action.payload}
        case GET_USER_ORDERS:
            return {...state, userOrder: action.payload}
        case SAVE_ORDER:
            return {...state, newOrder: action.payload}
        default:
            return state;
    }
};