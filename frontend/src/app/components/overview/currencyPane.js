import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import TransactionSection from './transactionsSection';
class CurrencyPane extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            transactions: props.props.transaction,
            currency: props.props.currency,
            amount: props.props.amount,
            prices:props.props.prices,
            
        }
    }
    render() {
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
        return (
            <div className="pane">
                <div className="pane__body">
                    <div className="body__img">
                        <img className="img" src={'http://localhost:8080/static/'+curr+'.svg'}></img>
                    </div>
                    <div className="body__info">
                        <span className="info__title">{curr.toUpperCase()}</span>
                        <div className="info__desc">
                            <span className="info__amount">{parseFloat(this.state.amount)} {this.state.currency}</span>
                            <span className="info__total">{price}{mon}</span>
                        </div>
                    </div>
                </div>
                <TransactionSection transactions={this.state.transactions} />
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(CurrencyPane);
