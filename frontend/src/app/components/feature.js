import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class Feature extends PureComponent {

    render() {
        return (
            <div>
                <h4>Feature</h4><small>You must be logged in to see the features</small>
                <ul>
                </ul>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Feature);
