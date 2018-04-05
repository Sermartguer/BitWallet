import {
    GET_USER_ADDRESSES
} from '../actions/types';

export const reducer = (state = {}, action) => {

    switch (action.type) {
        case GET_USER_ADDRESSES:
        console.log(action.payload)
            return { ...state, addresses: action.payload}
        default:
            return state;
    }
};