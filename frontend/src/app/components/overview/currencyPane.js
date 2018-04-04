import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import TransactionSection from './transactionsSection';
class CurrencyPane extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            transactions: props.props,
            currency: props.props.currency,
            amount: props.props.amount
        }
    }
    render() {
        return (
            <div className="pane">
                <div className="pane__body">
                    <div className="body__img">
                        <img className="img" src="http://dogecoin.com/imgs/dogecoin-300.png"></img>
                    </div>
                    <div className="body__info">
                        <span className="info__title">Monedaasdad</span>
                        <div className="info__desc">
                            <span className="info__amount">{parseFloat(this.state.amount)} {this.state.currency}</span>
                            <span className="info__total">5.30$</span>
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
