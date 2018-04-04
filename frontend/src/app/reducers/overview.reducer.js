import {
    GET_USER_BASIC
} from '../actions/types';

export const reducer = (state = {}, action) => {

    switch (action.type) {
        case GET_USER_BASIC:
        console.log(action.payload)
            return { ...state, overview: action.payload}
        default:
            return state;
    }
};