import {
    AUTH_USER,
    UNAUTH_USER,
    AUTH_ERROR,
    GET_PROFILE_DATA
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
        default:
            return state;
    }
};