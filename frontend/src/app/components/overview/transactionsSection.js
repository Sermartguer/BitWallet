import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';

class TransactionSection extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            transactions: props.transactions
        }
    }
    render() {
        return (
            <div>
                {this.state.transactions ? <div className="transaction">
                    <span className="transaction__mount">-350 DOGE</span>
                    <span className="transaction__to">to:DaoisdhashdbHJSIHGAFsdfsdFDFE</span>
                </div> : <div className="transaction">
                    <span className="transaction__mount">No transactions found</span>
                </div>}
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(TransactionSection);
