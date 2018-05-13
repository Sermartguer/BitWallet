import axios from 'axios';
import History from '../history.js';
import {
    LOGIN_HISTORY_SUCCESS,
    ACTION_HISTORY_SUCCESS,
    ORDER_HISTORY_SUCCESS
} from './types';

const ROOT_URL = 'http://localhost:8080/api';

export const getLoginHistory = ()=>{
    let data = {}
    data.token = localStorage.getItem('token');
    return (dispatch) => {
        console.log(data)
        axios.post('http://localhost:8080/api/loginHistory', data ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data)
            dispatch({type: LOGIN_HISTORY_SUCCESS, payload:response.data})
        }).catch((err) => {
            //dispatch(sendErrorExternalBalances(err.response))
            console.log(err.response)
        });
    };
}
export const getActionHistory = ()=>{
    let data = {}
    data.token = localStorage.getItem('token');
    return (dispatch) => {
        console.log(data)
        axios.post('http://localhost:8080/api/actionHistory', data ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data)
            dispatch({type: ACTION_HISTORY_SUCCESS, payload:response.data})
        }).catch((err) => {
            //dispatch(sendErrorExternalBalances(err.response))
            console.log(err.response)
        });
    };
}
export const getOrderHistory = ()=>{
    let data = {}
    data.token = localStorage.getItem('token');
    return (dispatch) => {
        console.log(data)
        axios.post('http://localhost:8080/api/orderHistory', data ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data)
            dispatch({type: ORDER_HISTORY_SUCCESS, payload:response.data})
        }).catch((err) => {
            //dispatch(sendErrorExternalBalances(err.response))
            console.log(err.response)
        });
    };
}