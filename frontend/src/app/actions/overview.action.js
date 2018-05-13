import axios from 'axios';
import History from '../history.js';
import {
    GET_USER_BASIC,
    GET_COIN_PRICE,
    GET_TRANSACTIONS,
} from './types';

const ROOT_URL = 'http://localhost:8080/api';

export const getUserBasic = () =>{
    let token = {token: localStorage.getItem('token')}
    console.log(token)
    return (dispatch) => {
        // submit email/password to the server
        axios.post('http://localhost:8080/api/getUserData',  token ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                dispatch({type: GET_USER_BASIC, payload:response.data});
            }).catch((err) => {
                console.log(err)
            });
    };
}
export const getCoinPrice = (currency) => {
    return (dispatch) => {
        axios.post('http://localhost:8080/api/getCurrencyPrice', {currency:currency} ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data.data)
            let res = {currency:response.data.data.network,data: response.data.data.prices}
            dispatch({type: GET_COIN_PRICE, payload:res})
        }).catch((err) => {
        });
    };
}
export const getUserTransactions = () => {
    let token = {token:localStorage.getItem('token')};

    return (dispatch) => {
        axios.post('http://localhost:8080/api/getUserTrans', token ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data)
            dispatch({type: GET_TRANSACTIONS, payload:response.data})
        }).catch((err) => {
        });
    };
}