import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';
class OrderHistoryComponent extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            orderArray: this.props.data
          };
    }
    render() {
        console.log('hello')
        var loginData = this.state.orderArray.map((orderLog)=>{
            let amount = Math.round(orderLog.amount * 100)/100;
            let price = Math.round(orderLog.price * 100)/100;
            return  <tr>
                        <td>{orderLog.action}</td>
                        <td>{orderLog.currency}</td>
                        <td>{amount}</td>
                        <td>{price}</td>
                        <td>{orderLog.time}</td>
                    </tr>
             });
        return (
            <div className="login__panel">  
                <table>
                    <thead>
                        <tr>
                            <td>Action</td>
                            <td>Currency</td>
                            <td>Amount</td>
                            <td>Price</td>
                            <td>Time</td>
                        </tr>
                    </thead>
                    <tbody>
                        {loginData}
                    </tbody>
                </table>
            </div>
        );
    }
}
const mapStateToProps = (state) => {
    return { 

     }
}

export default connect(mapStateToProps, actions)(OrderHistoryComponent);
