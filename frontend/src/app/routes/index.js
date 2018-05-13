import React from 'react';
import { Route } from 'react-router-dom';
import App from '../components/app';
import RequireAuth from '../components/auth/require_auth';
import LoginComponent from '../components/auth/loginContainer';
import RegisterComponent from '../components/auth/registerContainer';
import LogoutComponent from '../components/auth/logoutContainer';
import RecoverComponent from '../components/auth/recoverContainer';
import NewPasswordComponent from '../components/auth/newPassContainer';
import HomeContainer from '../components/home/homeContainer';
import ProfileContainer from '../components/auth/profileContainer';
import PinViewContainer from '../components/auth/pinViewContainer';

import Feature from '../components/feature';
import Dashboard from '../components/dashboard';
import ProfileMiddleware from '../components/feature';
const Routes = () => {
    return (
        <App>
            <Route exact path="/" component={HomeContainer} />
            <Route exact path="/signin" component={LoginComponent} />
            <Route exact path="/signout" component={LogoutComponent} />
            <Route exact path="/signup" component={RegisterComponent} />
            <Route exact path="/feature" component={RequireAuth(Feature)} />
            <Route exact path="/dashboard" component={RequireAuth(Dashboard)} />
            <Route exact path="/profile" component={RequireAuth(ProfileMiddleware)} />
            <Route exact path="/verify/:id" component={PinViewContainer} />
            <Route exact path="/recover" component={PinViewContainer}/>
            <Route exact path="/newpassword/:id" component={NewPasswordComponent}/>
        </App>
    );
};

export default Routes;