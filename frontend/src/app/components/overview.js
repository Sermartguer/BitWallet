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
        this.props.getCoinPrice('DOGE');
        this.props.getCoinPrice('BTC');
        this.props.getCoinPrice('LTC');
        this.props.getUserTransactions();
    }
    renderCurrency() {
        if (this.props.overview) {
            return (
                <CurrencyPane props={true}/>,
                <CurrencyPane props={true}/>
            );
        }
    }
    render() {
        if((this.props.overview !== undefined) && (this.props.DOGE !== undefined) && (this.props.BTC !== undefined) && (this.props.LTC !== undefined) && (this.props.transactions !== undefined)){
            var curr = this.props.overview.map((item, index)=>{
                let curr;
                let trans;
                if(item.currency === 'BTC'){
                    curr = this.props.BTC.data
                    this.props.transactions.map((transaction)=>{
                        if(transaction.currency === 'BTC'){
                            trans = transaction
                        }
                    });
                }else if( item.currency === 'LTC'){
                    curr = this.props.LTC.data
                    this.props.transactions.map((transaction)=>{
                        if(transaction.currency === 'LTC'){
                            trans = transaction
                        }
                    });
                }else if(item.currency === 'DOGE'){
                    curr = this.props.DOGE.data;
                    this.props.transactions.map((transaction)=>{
                        if(transaction.currency === 'DOGE'){
                            trans = transaction
                        }
                    });
                }
                return <CurrencyPane key={index} props={{currency:item.currency, amount:item.amount,prices:curr,transactions:this.props.transactions,transaction:trans}}/>
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
    console.log(state)
    return { 
        overview: state.overview.overview,
        currencyPrice: state.overview.currencyPrice,
        DOGE:state.overview.DOGE,
        BTC: state.overview.BTC,
        LTC: state.overview.LTC,
        transactions: state.overview.transactions
        
    }
}

export default connect(mapStateToProps, actions)(Overview);
