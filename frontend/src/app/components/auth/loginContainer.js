import React, { PureComponent } from 'react';
import { Field, reduxForm } from 'redux-form';
import * as actions from '../../actions';
import { connect } from 'react-redux';

class LoginComponent extends PureComponent {

    handleFormSubmit({ username, password }) {
        this.props.signinUser({ username, password })
    }

    renderError() {
        if (this.props.error) {
            console.log(this.props.error)
            return (
                <div className="alert__danger">
                    <string>Oops! {this.props.error}</string>
                </div>
            );
        }
    }

    render() {
        const { handleSubmit } = this.props;
        return (
            <div className="login">
                <div className="login__modal">
                    <form className="modal" onSubmit={handleSubmit(this.handleFormSubmit.bind(this))}>
                        <span className="login__title">Login</span>
                        <div className="login__question">
                            <span className="question--color">Don't have an account? </span>
                            <a className="question__a" href="/signup">Sign Up</a>
                            <br/>
                            <span className="question--color">Forgot password? </span>
                            <a className="question__a" href="/recover">Recover</a>
                        </div>
                        <div className="input__pattert">
                            <Field className="form__input" name="username" component="input" type="text" placeholder="Username"/>
                        </div>
                        <div className="input__pattert">
                            <Field className="form__input" name="password" component="input" type="password" placeholder="Password"/>
                        </div>
                        {this.renderError()}
                        <div className="form__button__pattern">
                            <button className="form__button" action="submit" >Sign in</button>
                        </div>
                    </form>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => ({...state.auth})


export default reduxForm({
    form: 'signin'
})(connect(mapStateToProps, actions)(LoginComponent));
