import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class Send extends PureComponent {
    render() {
        return (
            <div className="dash overview">
                <div className="pane chart__pane">
                    a
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Send);
