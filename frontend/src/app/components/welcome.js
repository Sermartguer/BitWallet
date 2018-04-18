import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class Welcome extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            currencyDetail: this.props.match.params.id
          };
    }
    componentWillMount(){
        console.log(this.props.match.params.id);
        if(this.props.match.params.id){
            this.props.verifyAccount(this.props.match.params.id);
        }
    }
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
                    <h1 className="plans__title">BitWallet Plans</h1>
                    <div className="plans__cards">
                        <div className="card card__basic card__basic--color">
                            <div className="basic__header">
                                <h1 className="header__title">Basic</h1>
                            </div>
                            <div className="basic__body">
                                <div className="body__item body__item--basic">
                                    Save multiple cryptocurrencies
                                </div>
                                <div className="body__item body__item--basic">
                                    Exchange and purchase
                                </div>
                                <div className="body__item">
                                    Send and recive
                                </div>
                            </div>
                            <div className="basic__footer">
                                <button className="btn btn__basic--color"> Get Plan</button>
                            </div>
                        </div>
                        <div className="card card__business card__business--color">
                            <div className="business__header">
                                <h1 className="header__title">Business</h1>
                            </div>
                            <div className="business__body">
                                <div className="body__item body__item--business">
                                    Save multiple cryptocurrencies
                                </div>
                                <div className="body__item body__item--business">
                                    Exchange and purchase
                                </div>
                                <div className="body__item body__item--business">
                                    Send and recive
                                </div>
                                <div className="body__item body__item--business">
                                    Business tools
                                </div>
                                <div className="body__item">
                                    Minimum fees
                                </div>
                            </div>
                            <div className="business__footer">
                                <button className="btn btn__business btn__business--color">Get Plan</button>
                            </div>
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

export default connect(mapStateToProps, actions)(Welcome);