import axios from 'axios';
import History from '../history.js';
import {
    GET_USER_ADDRESSES,
    SEND_LOCAL_SUCCESS,
    SEND_EXTERNAL_SUCCESS,
    SEND_LOCAL_ERROR,
    SEND_EXTERNAL_ERROR
} from './types';

const ROOT_URL = 'http://localhost:8080/api';
export const getUserAddresses = () =>{
    let token = {token: localStorage.getItem('token')}
    return (dispatch) => {
        // submit email/password to the server
        axios.post('http://localhost:8080/api/getAddresses',  token ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                dispatch({type: GET_USER_ADDRESSES, payload:response.data});
            }).catch((err) => {
                console.log(err)
            });
    };   
}
export const addNewAddress = (object) => {
    let response = {token: localStorage.getItem('token'),currency:object.currency, label: object.label}
    return (dispatch) => {
        // submit email/password to the server
        axios.post('http://localhost:8080/api/getNewAddress', response ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                dispatch({type: GET_USER_ADDRESSES, payload:response.data});
            }).catch((err) => {
                console.log(err)
            });
    };
}
export const sendLocalBalances = (data)=>{
    data.token = localStorage.getItem('token');
    return (dispatch) => {
        console.log(data)
        axios.post('http://localhost:8080/api/sendLocal', data ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data)
            dispatch({type: SEND_LOCAL_SUCCESS, payload:response.data})
        }).catch((err) => {
            dispatch(sendErrorLocalBalances(err.response))
            console.log(err.response)
        });
    };
}
export const sendExternalBalances = (data)=>{
    data.token = localStorage.getItem('token');
    return (dispatch) => {
        console.log(data)
        axios.post('http://localhost:8080/api/sendExternal', data ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data)
            dispatch({type: SEND_EXTERNAL_SUCCESS, payload:response.data})
        }).catch((err) => {
            dispatch(sendErrorExternalBalances(err.response))
            console.log(err.response)
        });
    };
}

export const sendErrorLocalBalances = (error) => {
    return {
        type: SEND_LOCAL_ERROR,
        payload: error.data
    };
};
export const sendErrorExternalBalances = (error) => {
    return {
        type: SEND_EXTERNAL_ERROR,
        payload: error.data
    };
};
