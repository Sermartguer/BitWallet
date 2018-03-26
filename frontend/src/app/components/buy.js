import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import Orders from './store/orders';

class Buy extends PureComponent {
    render() {
        return (
            <div>
                <div className="dash dashboard__title">
                    <span>Buy Orders</span>
                </div>
                <div className="dash dashboard__order">
                    <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                    <Orders props={{amount:0.15,currency:'BTC',price:20}}/>
                    <Orders props={{amount:0.02,currency:'LTC',price:13.5}}/>
                    <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                    <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                    <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                    <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                    <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                    <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                    <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                </div>
                <div className="dash dashboard__sell">
                    <div className="saleOrders">
                        <div className="dashboard__title">
                            <span>Your Sale Orders</span>
                        </div>
                        <div className="dash dashboard__sell">
                        <div className="dash dashboard__order">
                            <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                            <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                            <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                            <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
                            <Orders props={{amount:25,currency:'DOGE',price:0.015}}/>
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
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Buy);
