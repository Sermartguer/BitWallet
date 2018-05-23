import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';
import moment from 'moment'
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
                status = <i style={{color:"#7BCC70"}}className="fas fa-check-circle"></i>
            }else{
                status = <i style={{color:"#d94c4c"}} className="fas fa-exclamation-circle"></i>
            }
            let date = new Date(loginLog.time)
            console.log(date.getUTCDate())
            let time = moment([date.getFullYear(),date.getMonth()+1,date.getDate()]).fromNow(true)
            console.log(loginLog.time)
            return  <tr>
                        <td>{loginLog.ip}</td>
                        <td>{time} ago</td>
                        <td style={{textAlign:"center"}}>{status}</td>
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
