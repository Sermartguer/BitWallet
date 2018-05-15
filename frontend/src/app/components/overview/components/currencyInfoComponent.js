import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';
import TransactionInfoComponent from './transactionInfoComponent';
class CurrencyInfoComponent extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            transactions: props.props.transaction,
            currency: props.props.currency,
            amount: props.props.amount,
            prices:props.props.prices,
            orderBalance:null
        }
    }
    componentWillMount(){
        this.props.getOrderBalance(this.state.currency);
    }
    componentWillReceiveProps(nextProps) {
        this.setState({
            orderBalance:nextProps["order"+this.state.currency]
        })
    }
    render() {
        console.log(this.state.orderBalance)
        var balanceOrders = null;
        if((this.state.orderBalance != null) && (this.state.orderBalance != "")){
            balanceOrders = this.state.orderBalance;
        }else{
            balanceOrders = 0;
        }
        var curr = null;
        if(this.state.currency === "DOGE"){
            curr = "dogecoin";
        }else if(this.state.currency === "BTC"){
            curr = "bitcoin";
        }else if(this.state.currency === "LTC"){
            curr = "litecoin";
        }
        let price = eval(this.state.amount * this.state.prices[0].price);
        let mon = this.state.prices[0].price_base;
        var transactionUI;
        if(this.state.transactions !== undefined){
            transactionUI = <TransactionInfoComponent transactions={this.state.transactions} />
        }else{
            transactionUI =  <span className="center__trans">No transactions found</span>;
        }
        return (
            <div className="pane">
                <div className="pane__body">
                    <div className="body__img">
                        <img className="img" src={'http://localhost:8080/static/'+curr+'.svg'}></img>
                    </div>
                    <div className="body__info">
                        <span className="info__title">{curr.toUpperCase()}</span>
                        <div className="info__desc">
                            <span className="info__amount">{parseFloat(balanceOrders)} {parseFloat(this.state.amount)-parseFloat(balanceOrders)} {this.state.currency}</span>
                            <span className="info__total">{price}{mon}</span>
                        </div>
                    </div>
                </div>
                    {transactionUI}
            </div>
        );
    }
}

const mapStateToProps = (state) => ({...state.overview})

export default connect(mapStateToProps, actions)(CurrencyInfoComponent);
