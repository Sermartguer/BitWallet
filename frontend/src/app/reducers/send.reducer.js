import {
    GET_USER_ADDRESSES,
    SEND_LOCAL_ERROR,
    SEND_LOCAL_SUCCESS,
    SEND_EXTERNAL_SUCCESS,
    SEND_EXTERNAL_ERROR
} from '../actions/types';

export const reducer = (state = {}, action) => {

    switch (action.type) {
        case GET_USER_ADDRESSES:
        console.log(action.payload)
            return { ...state, addresses: action.payload}
        case SEND_LOCAL_SUCCESS:
            console.log(action.payload)
            return { ...state, success: action.payload, error:""}
        case SEND_LOCAL_ERROR:
            return {...state, error: action.payload}
        case SEND_EXTERNAL_SUCCESS:
            return { ...state, success: action.payload, error:""}
        case SEND_EXTERNAL_ERROR:
            return {...state, error: action.payload}
        default:
            return state;
    }
};