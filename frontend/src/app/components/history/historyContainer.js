import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import LoginHistoryComponent from './components/loginHistoryComponent';
import ActionHistoryComponent from './components/actionHistoryComponent';
import OrderHistoryComponent from './components/orderHistoryComponent';
class History extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            loginArray: "",
            actionArray:"",
            orderArray:"",
          };
    }
    componentWillMount(){
        this.props.getLoginHistory();
        this.props.getActionHistory();
        this.props.getOrderHistory();
    }
    componentWillReceiveProps(nextProps) {
        console.log(nextProps.action)
        this.setState({
            loginArray:nextProps.login,
            actionArray:nextProps.action,
            orderArray:nextProps.order,
        });
    }
    render() {
        console.log(this.state)
        if((this.state.loginArray != "") && (this.state.loginArray != undefined)){
            var loginData = <LoginHistoryComponent key={0} data={this.state.loginArray}/>
        }else{
            var loginData = "Loading login history..."
        }
        if((this.state.actionArray != "") && (this.state.actionArray != undefined)){
            var actionData = <ActionHistoryComponent key={1} data={this.state.actionArray}/>
        }else{
            var actionData = "Loading action history..."
        }
        if((this.state.orderArray != "") && (this.state.orderArray != undefined)){
            var orderData = <OrderHistoryComponent key={2} data={this.state.orderArray}/>
        }else{
            var orderData = "Loading action history..."
        }
        return (
            <div className="dash">
            {loginData}
            {actionData}
            {orderData}
            </div>
        );
    }
}

const mapStateToProps = (state) => ({...state.history})

export default connect(mapStateToProps, actions)(History);
