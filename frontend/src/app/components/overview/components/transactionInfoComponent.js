import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';

class TransactionInfoComponent extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            transactions: props.transactions,
            currency: props.currency
        }
    }
    render() {
        let trans;
        if(this.state.transactions != null){
            if(this.state.transactions.length < 1){
                trans = <div className="transaction">
                            <span className="transaction__mount">No transactions found</span>
                        </div>
            }else{
                let transaction = this.state.transactions[this.state.transactions.length-1];
                trans = <div className="transaction">
                            <span className="transaction__mount">{Math.round(transaction.amount * 100)/100 } {this.state.currency}</span>
                            <span className="transaction__to">to:{transaction.send}</span>
                        </div> 
            }
        }
        return (
            <div>
               {trans}
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(TransactionInfoComponent);
