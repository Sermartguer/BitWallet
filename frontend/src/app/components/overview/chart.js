import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
var LineChart = require("react-chartjs").Line;

class Chart extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            transactions: props.props
        }
    }
    render() {
        var chartData = [65, 59, 80, 81, 56, 55, 40];
    
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
