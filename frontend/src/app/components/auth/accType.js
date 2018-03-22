import React, { PureComponent } from 'react';
import { Field, reduxForm } from 'redux-form';
import * as actions from '../../actions';
import { connect } from 'react-redux';

class AccType extends PureComponent {

    handleFormSubmit(formProps) {
        this.props.signupUser(formProps)
    }

    render() {
        const { handleSubmit, submitting } = this.props;

        return (
            <div>
                <span>asdasd</span>
            </div>
        );
    }
}


const mapStateToProps = (state) => {
    return { errorMessage: state.auth.error }
};

export default reduxForm({
    form: 'signin'
})(connect(mapStateToProps, actions)(AccType));
