import axios from 'axios';
import History from '../history.js';
import {
    AUTH_USER,
    UNAUTH_USER,
    AUTH_ERROR,
    FETCH_FEATURE
} from './types';

const ROOT_URL = 'http://localhost:8080/api';

export const signinUser = ({ username, password }) => {
    console.log('asi')
    return (dispatch) => {
        // submit email/password to the server
        axios.post('http://localhost:8080/api/login', { username, password },{
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
                History.push('/feature');

            }).catch((err) => {
                console.log(err.response)
                // if request is bad...
                // - show an error to the user
                dispatch(authError(err.response.data));
            });
    };
};

export const signupUser = ({ username, email, password,password2,acc_type="basic" }) => {
    return (dispatch) => {
        console.log( username)
        // submit email/password to the server
        axios.post('http://localhost:8080/api/register', { username, email, password,password2,acc_type },{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          })
            .then(response => {
                console.log(response);
                //dispatch({ type: AUTH_USER });
                //localStorage.setItem('token', response.data.token);
                History.push('/');
            })
            .catch(err => {
                console.log(err.response)
                dispatch(authError(err.response.data));
            });
    };
};

export const authError = (error) => {
    return {
        type: AUTH_ERROR,
        payload: error
    };
};

export const signoutUser = () => {
    localStorage.removeItem('token')
    return { type: UNAUTH_USER };
};

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