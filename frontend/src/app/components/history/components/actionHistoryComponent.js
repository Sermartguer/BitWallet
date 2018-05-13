import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';
class ActionHistoryComponent extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            actionArray: this.props.data
          };
    }
    render() {
        var actionData = this.state.actionArray.map((actionLog)=>{
            let amount = Math.round(actionLog.amount * 100)/100;
            console.log(amount)
            return  <tr>
                        <td>{actionLog.action}</td>
                        <td>{actionLog.address_local}</td>
                        <td>{amount}</td>
                        <td>{actionLog.currency}</td>
                        <td>{actionLog.time}</td>
                    </tr>
             });
        return (
            <div className="login__panel">  
                <table>
                    <thead>
                        <tr>
                            <td>Send</td>
                            <td>To</td>
                            <td>Amount</td>
                            <td>Currency</td>
                            <td>Time</td>
                        </tr>
                    </thead>
                    <tbody>
                        {actionData}
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

export default connect(mapStateToProps, actions)(ActionHistoryComponent);
