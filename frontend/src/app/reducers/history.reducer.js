import {
    LOGIN_HISTORY_SUCCESS,
    ACTION_HISTORY_SUCCESS,
    ORDER_HISTORY_SUCCESS
} from '../actions/types';

export const reducer = (state = {}, action) => {
    switch (action.type) {        
        case LOGIN_HISTORY_SUCCESS:
            return { ...state, login: action.payload}
        case ACTION_HISTORY_SUCCESS:
            return { ...state, action:action.payload}
        case ORDER_HISTORY_SUCCESS:
            return { ...state, order:action.payload}
        default:
            return state;
    }
};