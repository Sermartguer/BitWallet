import {
    GET_USER_BASIC
} from '../actions/types';

export const reducer = (state = {}, action) => {

    switch (action.type) {
        case GET_USER_BASIC:
            return { ...state, homePageFeatures: action.payload}
        default:
            return state;
    }
};