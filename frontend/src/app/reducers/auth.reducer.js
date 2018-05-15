import {
    AUTH_USER,
    UNAUTH_USER,
    AUTH_ERROR,
    GET_PROFILE_DATA,
    VERIFY_ERROR,
    VERIFY_SUCCESS,
    UPDATE_RECOVER_FIELD,
    UPDATE_PASSWORD
} from '../actions/types';
var jwtDecode = require('jwt-decode');

export const reducer = (state = {}, action) => {
    switch (action.type) {
        case AUTH_USER:
        try {
            var decoded = jwtDecode(action.token);
            return { ...state, error: '', authenticated: true, username:decoded.sub }
        } catch (error) {
            return { ...state, error: '', authenticated: true }
        }    
        case UNAUTH_USER:
            return { ...state, authenticated: false }
        case AUTH_ERROR:
            return { ...state, error: action.payload }
        case GET_PROFILE_DATA:
            return {...state, error: '', profileData: action.payload}
        case VERIFY_ERROR:
            return { ...state, error:action.payload}
        case VERIFY_SUCCESS:
            return { ...state, success:action.payload}
        case UPDATE_RECOVER_FIELD:
            return { ...state, [action.key]:action.value}
        case UPDATE_PASSWORD:
        default:
            return state;
    }
};