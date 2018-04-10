import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import Orders from './store/orders';

class Buy extends PureComponent {
    componentWillMount(){
        this.props.getOrders();
        this.props.getUserOrders();
    }
    render() {
        if(this.props.buy){
            var orders = this.props.buy.map((order,index)=>{
                            let amount = Math.round(order.amount * 100)/100;
                            let price = Math.round(order.price * 100)/100;
                            return <Orders props={{amount:amount,currency:order.currency,price:price}}/>
                        });
        }
        if(this.props.userOrders){
            var userOrders = this.props.userOrders.map((order,index)=>{
                let amount = Math.round(order.amount * 100)/100;
                let price = Math.round(order.price * 100)/100;
                return <Orders props={{amount:amount,currency:order.currency,price:price}}/>
            });
        }
        return (
            <div>
                <div className="dash dashboard__title">
                    <span>Buy Orders</span>
                </div>
                <div className="dash dashboard__order">
                    {orders}
                </div>
                <div className="dash dashboard__sell">
                    <div className="saleOrders">
                        <div className="dashboard__title">
                            <span>Your Sale Orders</span>
                        </div>
                        <div className="dash dashboard__sell">
                        <div className="dash dashboard__order">
                            {userOrders}
                        </div>
                </div>
                    </div>
                    <div className="sale">
                        <div className="dashboard__title">
                            <span>Sell</span>
                        </div>
                        <div className="login__modal">
                            <form className="modal">
                                <span className="login__title">Make sale Order</span>
                                <div className="">
                                    <input className="form__input" name="currency" type="radio" placeholder="Price"/>
                                    <input className="form__input" name="currency" type="radio" placeholder="Price"/>
                                    <input className="form__input" name="currency" type="radio" placeholder="Price"/>
                                </div>
                                <div className="input__pattert">
                                    <input className="form__input" name="username" type="text" placeholder="Amount"/>
                                </div>
                                <div className="input__pattert">
                                    <input className="form__input" name="password" type="text" placeholder="Price"/>
                                </div>
                                <div className="form__button__pattern">
                                    <button className="form__button" action="submit" >Sign in</button>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    console.log(state)
    return { 
                buy: state.buy.buy,
                userOrders:state.buy.userOrder
            }
}

export default connect(mapStateToProps, actions)(Buy);
