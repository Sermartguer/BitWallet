import React from 'react';
import { Route } from 'react-router-dom';
import App from '../components/app';
import RequireAuth from '../components/auth/require_auth';
import Signin from '../components/auth/signin';
import Signout from '../components/auth/signout';
import Signup from '../components/auth/signup';
import Feature from '../components/feature';
import Welcome from '../components/welcome';
import Dashboard from '../components/dashboard';
import Profile from '../components/profile';
import ProfileMiddleware from '../components/feature';
import Recover from '../components/auth/recover';
import NewPassword from '../components/auth/newPass';
const Routes = () => {
    return (
        <App>
            <Route exact path="/" component={Welcome} />
            <Route exact path="/signin" component={Signin} />
            <Route exact path="/signout" component={Signout} />
            <Route exact path="/signup" component={Signup} />
            <Route exact path="/feature" component={RequireAuth(Feature)} />
            <Route exact path="/dashboard" component={RequireAuth(Dashboard)} />
            <Route exact path="/profile" component={RequireAuth(ProfileMiddleware)} />
            <Route exact path="/verify/:id" component={Welcome} />
            <Route exact path="/recover" component={Recover}/>
            <Route exact path="/newpassword/:id" component={NewPassword}/>
        </App>
    );
};

export default Routes;