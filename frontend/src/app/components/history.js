import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class History extends PureComponent {
    render() {
        return (
            <div className="dash">
            History
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(History);
