import axios from 'axios';
import History from '../history.js';
import {
    AUTH_USER,
    UNAUTH_USER,
    AUTH_ERROR,
} from './types';

export const signinUser = (param) => {
    console.log(param);
    let response = {"username":param.username,"password":param.password,"ip":'123'}
    return (dispatch) => {
        // submit email/password to the server
        axios.post('http://localhost:8080/api/login', response,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          })
            .then(response => {

                // if request is good...
                // - update state to indicate user is authenticated
                console.log(response.data);
                localStorage.setItem('token', response.data.AuthToken);
                dispatch({ token:response.data.AuthToken,type: AUTH_USER });
                // - save the jwt token
                

                // - redirect to the route '/feature'
                History.push('/');

            }).catch((err) => {
                console.log(err.response)
                // if request is bad...
                // - show an error to the user
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
            }).catch((err) => {
                console.log(err)
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
                //dispatch({type: GET_PROFILE_DATA, payload:response.data[0]});
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

