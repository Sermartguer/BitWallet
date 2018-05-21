import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';
import Modal from 'react-modal';
const customStyles = {
    content : {
      top                   : '50%',
      left                  : '50%',
      right                 : 'auto',
      bottom                : 'auto',
      marginRight           : '-50%',
      transform             : 'translate(-50%, -50%)',
      width                 : '500px',
      overflow:'hidden'
    }
  };
  Modal.setAppElement('#modal')

class OrderComponent extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            amount: props.props.amount,
            currency: props.props.currency,
            price: props.props.price,
            modalIsOpen: false,
            id:props.props.id,
            amount_buy:0,
            total_pay:0,
            currency_to:props.props.currency_to
        }
        this.handleInputChange = this.handleInputChange.bind(this);
        this.openModal = this.openModal.bind(this);
        this.closeModal = this.closeModal.bind(this);
        this.submitPay = this.submitPay.bind(this)
    }
    openModal() {
        let a = document.getElementsByClassName('dash');
        a[0].classList.add("blur")
        this.setState({modalIsOpen: true});
      }
    closeModal() {
        let a = document.getElementsByClassName('dash');
        a[0].classList.remove("blur")
        this.setState({modalIsOpen: false});
    }
    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;
        if(name === 'amount_buy'){
            console.log(this.state.price)
            let total =parseFloat(this.state.price).toFixed(10)*value;            
            this.setState({
                total_pay:total,
                [name]: value
            });
        }else{
            this.setState({
                [name]: value
            });
        }
        
    }
    submitPay(){
        if(this.state.amount_buy > this.state.amount){
            document.getElementById('error').innerHTML= 'Error the amount to buy is higher that balance aviable'
        }else{
            let object = {"restAmount":""+eval(this.state.amount-this.state.amount_buy)+"","payTo":this.state.id,"currencyO":this.state.currency,"currencyD":this.state.currency_to,"totalToPay":this.state.amount_buy,"totalToRecive":""+this.state.total_pay+""}
            this.props.payOrder(object)
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
                        <span>{this.state.price} BTC /{this.state.currency}</span>
                    </div>
                </div>
                <div className="footer__card">
                    <button className={`order__button ${styleBackGround}`} onClick={this.openModal}>Buy DOGE</button>
                </div>
                <Modal
                    isOpen={this.state.modalIsOpen}
                    onAfterOpen={this.afterOpenModal}
                    onRequestClose={this.closeModal}
                    style={customStyles}
                    contentLabel="Example Modal">
                    <div className="modal__pattern">
                        <div className="modal__header">Amount to buy</div>
                        <div className="input__pattert">
                            <input className="form__input" name="amount_buy" type="number" placeholder="Amount to buy" step="0.001" min="0" max={this.state.amount} onChange={this.handleInputChange}/>
                        </div>
                        <div className="modal__pricing modal__preview">
                            <span>Price by {this.state.currency}: {this.state.price} {this.state.currency_to}</span>
                        </div>
                        <div className="modal__pricing modal--size">
                            <span className="modal__total modal__total--color">Total Pay {this.state.total_pay}</span>
                        </div>
                        <span id="error"></span>
                        <button className="form__button" onClick={this.submitPay}>Buy DOGE</button>
                    </div>
                    </Modal>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(OrderComponent);
