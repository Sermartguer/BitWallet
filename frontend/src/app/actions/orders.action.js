import axios from 'axios';
import History from '../history.js';
import {
    GET_ORDERS,
    GET_USER_ORDERS,
    SAVE_ORDER,
    OVERVIEW_ORDER_BALANCE
} from './types';

const ROOT_URL = 'http://localhost:8080/api';

export const getOrders = () => {
    let token = {token: localStorage.getItem('token')}
    return (dispatch) => {
        // submit email/password to the server
        axios.post('http://localhost:8080/api/getOrders', token ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
              console.log(response)
                dispatch({type: GET_ORDERS, payload:response.data});
            }).catch((err) => {
                console.log(err)
            });
    };
}
export const getUserOrders = () => {
    let token = {token: localStorage.getItem('token')}
    return (dispatch) => {
        // submit email/password to the server
        axios.post('http://localhost:8080/api/getUserOrders', token ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
              console.log(response)
                dispatch({type: GET_USER_ORDERS, payload:response.data});
            }).catch((err) => {
                console.log(err)
            });
    };
}

export const addNewOrder = (orderData) => {
    let response = {token: localStorage.getItem('token'),amount:orderData.amount,price:orderData.price,currency:orderData.currency,currency_to:orderData.currency_to}
    return (dispatch) => {
        // submit email/password to the server
        axios.post('http://localhost:8080/api/saveOrder', response ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                dispatch({type: SAVE_ORDER, payload:response.data});
            }).catch((err) => {
                console.log(err)
            });
    };
}
export const disableOrderNew = () => {
    return (dispatch) => {
        dispatch({type: SAVE_ORDER, payload:false});
    };
}
export const getOrderBalanceOrders= (currency) => {
    let data = {}
    data.token = localStorage.getItem('token');
    data.currency = currency;
    return (dispatch) => {
        axios.post('http://localhost:8080/api/getOrderBalance', data ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            //console.log(response.data)
            dispatch({type: OVERVIEW_ORDER_BALANCE, payload:response.data, currency:currency})
        }).catch((err) => {
        });
    };
}
export const payOrder = (dataPay) =>{
    dataPay.token = localStorage.getItem('token');
    return (dispatch) => {
        axios.post('http://localhost:8080/api/payOrder', dataPay ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data)
            //dispatch({type: OVERVIEW_ORDER_BALANCE, payload:response.data, currency:currency})
        }).catch((err) => {
        });
    };
}