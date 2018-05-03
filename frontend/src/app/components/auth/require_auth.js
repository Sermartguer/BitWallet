import React, { Component } from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

export default function (ComposedComponent) {
    class Authentication extends Component {
        static contextTypes = {
            router: PropTypes.object
        }

        componentWillMount() {
            if (this.props.authenticated === undefined) {
                console.log('undefined')
                this.props.history.push('/signin');
            }
        }

        componentWillUpdate(nextProps) {
            if (this.props.authenticated === undefined) {
                this.props.history.push('/signin');
            }
        }

        render() {
            return <ComposedComponent {...this.props} />
        }
    }

    function mapStateToProps(state) {
        return { authenticated: state.auth.authenticated };
    }

    return connect(mapStateToProps)(Authentication);
}
