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
        
          };
    }
    componentWillMount(){
        this.props.getUserBasic();
    }
    renderCurrency() {
        if (this.props.overview) {
            console.log(this.props.overview)
            return (
                    <CurrencyPane props={true}/>,
                    <CurrencyPane props={true}/>
            );
        }
    }
    render() {
        console.log(this.props.overview)
        if(this.props.overview !== undefined){
            console.log('a')
            var curr = this.props.overview.map((item, index)=>{
                console.log(item)
                return <CurrencyPane key={index} props={{currency:item.currency, amount:item.amount}}/>
            })
        }else{
            var curr = 'Loading...'
        }
        return (
            <div className="dash overview">
                <div className="overview__detail">
                    {curr}
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
    return { overview: state.overview.overview }
}

export default connect(mapStateToProps, actions)(Overview);
