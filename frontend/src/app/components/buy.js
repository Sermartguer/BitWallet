import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import Orders from './store/orders';
class Buy extends PureComponent {
    render() {
        return (
            <div>
                <div className="dash dashboard__title">
                    <span>Buy Orders</span>
                </div>
                <div className="dash dashboard__order">
                    <Orders />
                    <Orders />
                    <Orders />
                    <Orders />
                    <Orders />
                    <Orders />
                    <Orders />
                    <Orders />
                    <Orders />
                    <Orders />
                    <Orders />
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Buy);
