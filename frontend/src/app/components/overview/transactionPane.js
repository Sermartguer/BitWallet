import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import ReactTable from "react-table";
class TransactionPane extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            transactions: props.props
        }
        
    }
    render() {
        const data = [{
            action:'Recive',
            moneda: 'DOGE',
            amount: '26 DOGE',
            to: 'DNJDAHShduhasdaiBSDHSA',
            date: '27/04/18'
          },{
            action:'Send',
            moneda: 'BTC',
            amount: '0.04 BTC',
            to: 'DNJDAHShduhasdaiBSDHSA',
            date: '26/04/18'
          },
          {
            action:'Send',
            moneda: 'LTC',
            amount: '1.15 LTC',
            to: 'DNJDAHShduhasdaiBSDHSA',
            date: '26/04/18'
          }]
        
          const columns = [
            {
                Header: 'Action',
                accessor: 'action'
            },
            {
                Header: 'Moneda',
                accessor: 'moneda'
            }, {
                Header: 'Amount',
                accessor: 'amount',
            },{
                Header : 'To',
                accessor: 'to'
            },{
                Header: 'Date',
                accessor: 'date'
            }
        ]
        return (
            <div className="pane transaction__pane">
                <span className="pane__title">Transactions</span>
                <ReactTable
                data={data}
                columns={columns}
                defaultPageSize={5} className="-striped -highlight full__pane"/>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(TransactionPane);
