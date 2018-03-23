import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
class Chart extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            transactions: props.props
        }
    }
    render() {
        return (
            <div className="pane chart__pane">

            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Chart);
