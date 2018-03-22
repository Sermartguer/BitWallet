import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import CurrencyPane from './overview/currencyPane';
class Overview extends PureComponent {
    render() {
        return (
            <div className="dash overview">
                <div className="overview__detail">
                    <CurrencyPane props={true}/>
                    <CurrencyPane props={false}/>
                    <CurrencyPane props={false}/>
                </div>
                <div>a</div>
                <div>a</div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Overview);
