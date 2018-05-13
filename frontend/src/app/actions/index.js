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
    let response = {token: localStorage.getItem('token'),amount:orderData.amount,price:orderData.price,currency:orderData.currency}
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
export const setPin = (pin) =>{
    return (dispatch) => {
        console.log(pin)
        /*axios.post('http://localhost:8080/api/getUserTrans', token ,{
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
        }).then(response => {
            console.log(response.data)
            dispatch({type: GET_TRANSACTIONS, payload:response.data})
        }).catch((err) => {
        });
    };*/
    }
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