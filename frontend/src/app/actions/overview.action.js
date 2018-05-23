import axios from 'axios';
import History from '../history.js';
import {
    GET_USER_BASIC,
    GET_COIN_PRICE,
    GET_TRANSACTIONS,
    OVERVIEW_ORDER_STATUS
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
    let curr;
    if(currency === 'BTC'){
        curr = 'bitcoin'
    }else if(currency === 'DOGE'){
        curr = 'dogecoin'
    }else if(currency === 'LTC'){
        curr = 'litecoin'
    }
    return (dispatch) => {
        axios.get('https://api.coinmarketcap.com/v1/ticker/'+curr+'/',{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            dispatch({type: GET_COIN_PRICE, payload:{currency:currency+'price',data:response.data[0].price_usd}})
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
export const getOrderBalance= (currency) => {
    let data = {}
    data.token = localStorage.getItem('token');
    data.currency = currency;
    return (dispatch) => {
        axios.post('http://localhost:8080/api/getOrderBalance', data ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            //console.log(response.data)
            dispatch({type: OVERVIEW_ORDER_STATUS, payload:response.data, currency:currency})
        }).catch((err) => {
        });
    };
}