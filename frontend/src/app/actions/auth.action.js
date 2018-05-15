import axios from 'axios';
import History from '../history.js';
import {
    AUTH_USER,
    UNAUTH_USER,
    AUTH_ERROR,
    VERIFY_ERROR,
    VERIFY_SUCCESS
} from './types';

export const signinUser = (param) => {
    console.log(param);
    let response = {"username":param.username,"password":param.password,"ip":'123'}
    return (dispatch) => {
        axios.post('http://localhost:8080/api/login', response,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          })
            .then(response => {
                localStorage.setItem('token', response.data.AuthToken);
                dispatch({ token:response.data.AuthToken,type: AUTH_USER });
                History.push('/');
            }).catch((err) => {
                console.log(err.response)
                dispatch(authError(err.response.data));
            });
    };
};
export const newPassword = (form) =>{
    console.log(form)
    return (dispatch) => {
        axios.post('http://localhost:8080/api/newPassword', form ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                History.push('/');
                console.log(response)
            }).catch((err) => {
                console.log(err)
            });
    };
}
export const verifyAccount = (verifyPin) => {
    console.log(verifyPin)
    return (dispatch) => {
        axios.post('http://localhost:8080/api/verifyAccount', verifyPin ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                console.log(response)
                dispatch({payload:'Success', type:VERIFY_SUCCESS})
            }).catch((err) => {
                console.log(err.response.data)
                dispatch({payload:err.response.data, type:VERIFY_ERROR})
            });
    };
}
export const updateProfile = (formValues) =>{
    console.log(formValues);
    formValues.token = localStorage.getItem('token')
    return (dispatch) => {
        axios.post('http://localhost:8080/api/updateProfile', formValues ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                console.log(response)
            }).catch((err) => {
                console.log(err)
            });
    };
}
export const recoverPassword = (email) =>{
    console.log(email)
    return (dispatch) => {
        axios.post('http://localhost:8080/api/recoverPassword', email ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }).then(response => {
                console.log(response)
                History.push('/');
            }).catch((err) => {
                console.log(err)
            });
    };
}
export const signupUser = ({ username, email, password,password2,account }) => {
    return (dispatch) => {
        console.log( username)
        // submit email/password to the server
        let acc_type = account;
        axios.post('http://localhost:8080/api/register', { username, email, password,password2,acc_type },{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          })
            .then(response => {
                localStorage.setItem('pin',false);
                History.push('/');
            })
            .catch(err => {
                console.log(err.response);
                dispatch(authError(err.response.data));
            });
    };
};
export const signoutUser = () => {
    localStorage.removeItem('token')
    return { type: UNAUTH_USER };
};
export const authError = (error) => {
    return {
        type: AUTH_ERROR,
        payload: error
    };
};
const ROOT_URL = 'http://localhost:8080/api';

