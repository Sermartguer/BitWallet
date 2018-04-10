import {
    GET_ORDERS
} from '../actions/types';

export const reducer = (state = {}, action) => {

    switch (action.type) {
        case GET_ORDERS:
        console.log(action.payload)
            return { ...state, buy: action.payload}
        default:
            return state;
    }
};