import axios from 'axios';
import History from '../history.js';
import {
    AUTH_USER,
    UNAUTH_USER,
    AUTH_ERROR,
    FETCH_FEATURE,
    GET_USER_BASIC,
    GET_USER_ADDRESSES,
    GET_ORDERS,
    GET_USER_ORDERS,
    SAVE_ORDER,
    GET_PROFILE_DATA,
    GET_COIN_PRICE,
    GET_TRANSACTIONS,
    SEND_LOCAL_ERROR,
    SEND_LOCAL_SUCCESS,
    SEND_EXTERNAL_SUCCESS,
    SEND_EXTERNAL_ERROR,
    LOGIN_HISTORY_SUCCESS,
    ACTION_HISTORY_SUCCESS,
    ORDER_HISTORY_SUCCESS
} from './types';
const ROOT_URL = 'http://localhost:8080/api';
export const fetchFeature = () => {
    return (dispatch) => {
        axios.get(ROOT_URL, {
            headers: { authorization: localStorage.getItem('token'),
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json',
               }
        })
        .then(response =>{
            dispatch({
                type: FETCH_FEATURE,
                payload: response.data
             });
        });
    };
};
export const getProfileData = () =>{
    let token = {token:localStorage.getItem('token')};
    return (dispatch) => {
        axios.post('http://localhost:8080/api/getAccountProfile', token ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                console.log(response)
                dispatch({type: GET_PROFILE_DATA, payload:response.data[0]});
            }).catch((err) => {
                console.log(err)
            });
    };
}
export * from './orders.action'
export * from './send.action'
export * from './history.action'
export * from './auth.action'
export * from './overview.action'