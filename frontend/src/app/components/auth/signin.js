import React, { PureComponent } from 'react';
import { Field, reduxForm } from 'redux-form';
import * as actions from '../../actions';
import { connect } from 'react-redux';

class Signin extends PureComponent {

    handleFormSubmit({ username, password }) {
        this.props.signinUser({ username, password })
    }

    renderError() {
        if (this.props.errorMessage) {
            console.log(this.props.errorMessage)
            return (
                <div className="alert__danger">
                    <string>Oops! {this.props.errorMessage}</string>
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

const mapStateToProps = (state) => {
    console.log(state.auth )
    return { errorMessage: state.auth.error }
};

export default reduxForm({
    form: 'signin'
})(connect(mapStateToProps, actions)(Signin));
