import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';
class LoginHistoryComponent extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            loginArray: this.props.data
          };
    }
    render() {
        console.log('hello')
        var loginData = this.state.loginArray.map((loginLog)=>{
            let status = "";
            if(loginLog.success === "1"){
                status = "Success"
            }else{
                status = "Error"
            }
            return  <tr>
                        <td>{loginLog.ip}</td>
                        <td>{loginLog.time}</td>
                        <td>{status}</td>
                    </tr>
             });
        return (
            <div className="login__panel">  
                <table>
                    <thead>
                        <tr>
                            <td>Ip Address</td>
                            <td>Time</td>
                            <td>Status</td>
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

export default connect(mapStateToProps, actions)(LoginHistoryComponent);
