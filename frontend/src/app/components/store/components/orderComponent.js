import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';
class OrderComponent extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            amount: props.props.amount,
            currency: props.props.currency,
            price: props.props.price
        }
    }
    render() {
        var curr = null;
        var styleBackGround = null;
        if(this.state.currency === "DOGE"){
            curr = "dogecoin";
            styleBackGround =" doge--color";
        }else if(this.state.currency === "BTC"){
            console.log('BTC')
            curr = "bitcoin";
            styleBackGround =" bitcoin--color";
        }else if(this.state.currency === "LTC"){
            curr = "litecoin";
            styleBackGround =" litecoin--color";
        }
        return (
            <div className={`pane order__pane card__size`}>
                <div className={`header__card ${styleBackGround}`}>
                    <img src={'http://localhost:8080/static/'+curr+'.png'} width="25"></img>
                </div>
                <div className="body__card">
                    <div className="text__card">
                        <span>Aviable:</span>
                        <span>{this.state.amount} {this.state.currency}</span>
                    </div>
                    <div className="text__card">
                        <span>Price:</span> 
                        <span>{this.state.price}$ /{this.state.currency}</span>
                    </div>
                </div>
                <div className="footer__card">
                    <button className={`order__button ${styleBackGround}`}>Buy DOGE</button>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(OrderComponent);
