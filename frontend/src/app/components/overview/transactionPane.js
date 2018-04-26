import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import ReactTable from "react-table";
class TransactionPane extends PureComponent {
    render() {
       if(this.props.trans !== undefined){
            var data = [];
            this.props.trans.map((transaction)=>{
                let object = {}
                if(Math.sign(transaction.amount) === 1){
                    object.action = 'Recive';
                }else{
                    object.action = 'Send';
                }
                object.moneda = transaction.currency;
                object.amount = Math.round(transaction.amount * 100)/100;
                object.to = transaction.send_to;
                let d = new Date(transaction.trans_time)
                var curr_date = d.getDate();
                var curr_month = d.getMonth() + 1;
                var curr_year = d.getFullYear();
                object.date = curr_date + "/" + curr_month + "/" + curr_year;
                data.push(object);
            });
        }else{
        }
        
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
    return { 
        trans: state.overview.transactions
     }
}

export default connect(mapStateToProps, actions)(TransactionPane);
