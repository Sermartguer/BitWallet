import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import CurrencyPane from './overview/currencyPane';
import Chart from './overview/chart';
import TransactionPane from './overview/transactionPane';
class Overview extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            currencyDetail: ''
          };
       // this.onButtonClick = this.onButtonClick.bind(this);
    }
    componentWillMount(){ 
        this.state.currencyDetail = this.props.getUserBasic();
    }
    render() {
        console.log(this.state.currencyDetail)
        let currency = this.state.currencyDetail.map(item=>{
            return <CurrencyPane props={item}/>
        });
        return (
            <div className="dash overview">
                <div className="overview__detail">
                    {currency}
                </div>
                <div>
                    <Chart />
                </div>
                <div>
                    <TransactionPane />
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    console.log(state)
    return { overview: state.overview }
}

export default connect(mapStateToProps, actions)(Overview);
