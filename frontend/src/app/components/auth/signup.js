import React, { PureComponent } from 'react';
import { Field, reduxForm } from 'redux-form';
import * as actions from '../../actions';
import { connect } from 'react-redux';
import AccType from './accType';
class Signup extends PureComponent {
    handleFormSubmit(formProps) {
        console.log(formProps)
        this.props.signupUser(formProps)
    }

    renderField = ({ input, label, type, meta: { touched, error },className="form__input" }) => (
        <div style={{textAlign:'center'}}>
            <input className={className} {...input} placeholder={label} type={type} />
            {touched && error && <span className="alert__danger">{error}</span>}
        </div>
    );

    renderError() {
        if (this.props.errorMessage) {
            return (
                <div className="alert__danger">
                    <string>Oops! {this.props.errorMessage}</string>
                </div>
            );
        }
    }

    render() {
        const { handleSubmit, submitting } = this.props;
        console.log(this.props)
        return (
            <div className="login">
                <div className="login__modal">

                <form className="modal" onSubmit={handleSubmit(this.handleFormSubmit.bind(this))}>
                        <span className="login__title">Sign Up</span>
                        <div className="input__pattert">
                            <Field name="username" component={this.renderField} type="text" label="Username"/>
                        </div>
                        <div className="input__pattert">
                            <Field name="email" component={this.renderField} type="text" label="Email"/>
                        </div>
                        <div className="input__pattert">
                            <Field name="password" component={this.renderField} type="password" label="Password"/>
                        </div>
                        <div className="input__pattert">
                            <Field name="password2" component={this.renderField} type="password" label="Repeat Password"/>
                        </div>
                        <div className="accounts">
                            <div className="account__section">
                                <i className="far fa-user fa-3x"></i>
                                
                                <label><div>Basic</div><Field type="radio" className="size__radio" component={this.renderField} name="account" value="Basic"/></label>
                            </div>
                            <div className="account__section">
                                <i className="far fa-gem fa-3x"></i>
                                <div></div>
                                <label><div>Business</div><Field type="radio" className="size__radio" component={this.renderField} name="account" value="Business"/></label>
                            </div>
                        </div>
                        {this.renderError()}
                        <div className="form__button__pattern">
                            <button className="form__button" type="submit" disabled={submitting}>Sign up</button>
                        </div>
                    </form>
                </div>
            </div>
        );
    }
}

const validate = values => {
    const errors = {};

    if (!values.username){
        errors.username = 'Please enter an username';
    }else if(values.username.length < 3){
        
        errors.username = 'Username must be more than 2 characters';
    }
    if (!values.email) {
        errors.email = 'Please enter an email';
    } else if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)) {
        errors.email = 'Invalid email address';
    }

    if (!values.password) {
        errors.password = 'Please enter an password';
    }else if((values.password.length < 5) || (values.password.length > 10)){
        errors.password = 'Password must be min: 5 characters, max: 10 characters';
    }

    if (!values.password2) {
        errors.password = 'Please enter an password confirmation';
    }else if((values.password2.length < 5) || (values.password2.length > 10)){
        errors.password = 'Password must be min: 5 characters, max: 10 characters';
    }

    if (values.password !== values.password2) {
        errors.password = 'Password must match';
    }

    return errors;
};

const mapStateToProps = (state) => {
    return { errorMessage: state.auth.error }
};

export default reduxForm({
    form: 'signin',
    validate
})(connect(mapStateToProps, actions)(Signup));
