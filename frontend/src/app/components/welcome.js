import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class Welcome extends PureComponent {
    render() {
        return (
            <div className="welcome">
                <div className="welcome__price price price--color">
                    <h1 className="price--title">
                        BitWallet, your criptocurrency wallet
                    </h1>
                    <div className="price__button">
                       <button className="btn btn__started btn__started--color btn__started--font"> Get Started</button>
                    </div>
                    <div className="price__curr currInfo">
                        <div className="currInfo__panel currInfo__panel--color panel">
                            <div className="panel__header">
                                <span className="panel__title panel__title--color">
                                    <img src="http://localhost:8080/static/bitcoin.svg" width="30" className="panel__title__img--size"></img>
                                    Bitcoin
                                </span>
                                <span className="panel__update panel__update--color">
                                    Last Update: 4 days
                                </span>
                            </div>
                            <div className="panel__body body">
                                <span className="body__pattern body__pattern--size body__pattern--color">1 BTC</span>
                                <i className="fas fa-exchange-alt body__pattern--color"></i>
                                <span className="body__pattern body__pattern--size body__pattern--color">123456 €</span>
                            </div>
                            <div className="panel__footer">
                                <div className="panel__button">
                                    <button className="btn btn__more btn__more--color btn__more--font"> View More</button>
                                </div>
                            </div>
                        </div>
                        <div className="currInfo__panel currInfo__panel--color panel">
                            <div className="panel__header">
                                <span className="panel__title panel__title--color">
                                    <img src="http://localhost:8080/static/dogecoin.svg" width="30" className="panel__title__img--size"></img>
                                    Dogecoin
                                </span>
                                <span className="panel__update panel__update--color">
                                    Last Update: 4 days
                                </span>
                            </div>
                            <div className="panel__body body">
                                <span className="body__pattern body__pattern--size body__pattern--color">1 DOGE</span>
                                <i className="fas fa-exchange-alt body__pattern--color"></i>
                                <span className="body__pattern body__pattern--size body__pattern--color">123456 €</span>
                            </div>
                            <div className="panel__footer">
                                <div className="panel__button">
                                    <button className="btn btn__more btn__more--color btn__more--font"> View More</button>
                                </div>
                            </div>
                        </div>
                        <div className="currInfo__panel currInfo__panel--color panel">
                            <div className="panel__header">
                                <span className="panel__title panel__title--color">
                                    <img src="http://localhost:8080/static/litecoin.svg" width="30" className="panel__title__img--size"></img>
                                    Litecoin
                                </span>
                                <span className="panel__update panel__update--color">
                                    Last Update: 4 days
                                </span>
                            </div>
                            <div className="panel__body body">
                                <span className="body__pattern body__pattern--size body__pattern--color">1 LTC</span>
                                <i className="fas fa-exchange-alt body__pattern--color"></i>
                                <span className="body__pattern body__pattern--size body__pattern--color">123456 €</span>
                            </div>
                            <div className="panel__footer">
                                <div className="panel__button">
                                    <button className="btn btn__more btn__more--color btn__more--font"> View More</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="welcome__info">
                    <div className="item__top">
                        <span className="item__title item__title--color">Try our cryptocurrency system</span>
                        <span className="item__subtitle item__subtitle--color">Here are a few reasons why you should choose BitWallet</span>
                    </div>
                    <div className="info">
                        <div className="info__item">
                            <i className="fas fa-key info__icon info__icon--size info__icon--color"></i>
                            <span className="info__pattern">Save</span>
                            <span className="info__desc">We accept different popular cryptocurrencies like Bitcoin, Dogecoin or Litecoin, store your coins in our secure online wallet.</span>
                        </div>  
                        <div className="info__item">
                            <i className="fas fa-handshake info__icon info__icon--size info__icon--color"></i>
                            <span className="info__pattern">Exchange</span>
                            <span className="info__desc">Exchange between different available cryptocurrencies, with our system of sale and purchase of cryptocurrencies with minimum costs</span>
                        </div>
                        <div className="info__item">
                            <i className="fas fa-shopping-cart info__icon info__icon--size info__icon--color"></i>
                            <span className="info__pattern">Send and Recive</span>
                            <span className="info__desc">Send and receive cryptocurrencies, generate addresses and send to an address with QR code</span>
                        </div>
                    </div>
                    
                </div>
                <div className="welcome__plans">
                    <span className="info__pattern">Are you a company?</span>
                    <span>Register with our business plan and get improvements, such as:</span>
                    <ul className="">
                        <li className="plan__center">Less taxes</li>
                        <li className="plan__center">Vendor tools</li>
                        <li className="plan__center">Vccount report at all times</li>
                    </ul>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Welcome);