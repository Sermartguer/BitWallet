import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import Orders from './store/orders';

class Buy extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            currency:'DOGE',
            amount:0,
            price:0,
            newOrders:false
          };
          this.handleInputChange = this.handleInputChange.bind(this);
          this.handleSubmit = this.handleSubmit.bind(this);      

        }
    componentWillMount(){
        this.props.getOrders();
        this.props.getUserOrders();
    }
    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;
    
        this.setState({
          [name]: value
        });
      }
      handleSubmit(event) {       
        event.preventDefault();
        this.props.addNewOrder(this.state);
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
        if(this.props.newOrder !== false){
            this.props.getOrders();
            this.props.getUserOrders();
            this.props.disableOrderNew();
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
                            <form className="modal" onSubmit={this.handleSubmit}>
                                <span className="login__title">Make sale Order</span>
                                <div className="option__sale">
                                    <label>
                                        <input className="form__input" name="currency" type="radio" value="BTC" onChange={this.handleInputChange} />
                                        <img src="http://localhost:8080/static/bitcoin.svg" width="30" heigth="530" />
                                    </label>
                                    <label>
                                        <input className="form__input" name="currency" type="radio" placeholder="Price" value="DOGE" onChange={this.handleInputChange}/>
                                        <img src="http://localhost:8080/static/dogecoin.svg" width="30" heigth="30"/>
                                    </label>
                                    <label>
                                        <input className="form__input" name="currency" type="radio" placeholder="Price" value="LTC" onChange={this.handleInputChange}/>
                                        <img src="http://localhost:8080/static/litecoin.svg" width="30" heigth="30"/>
                                    </label>
                                </div>
                                <div className="input__pattert">
                                    <input className="form__input" name="amount" type="number" placeholder="Amount" onChange={this.handleInputChange}/>
                                </div>
                                <div className="input__pattert">
                                    <input className="form__input" name="price" type="number" placeholder="Price" onChange={this.handleInputChange}/>
                                </div>
                                <div className="form__button__pattern">
                                    <button className="form__button" type="submit">Make sale Order</button>
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
                userOrders:state.buy.userOrder,
                newOrder:state.buy.newOrder
            }
}

export default connect(mapStateToProps, actions)(Buy);
