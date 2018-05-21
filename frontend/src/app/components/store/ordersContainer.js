import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import OrderComponent from './components/orderComponent';
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

class OrdersContainer extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            currency:'DOGE',
            amount:0,
            price:0,
            currency_to:'',
            newOrders:false,
            modalIsOpen: false,
            maxCURR:0,
            minCURR:3,
            actualBalance:0,
            total:0
          };
          this.handleInputChange = this.handleInputChange.bind(this);
          this.handleSubmit = this.handleSubmit.bind(this);
          this.openModal = this.openModal.bind(this);
          this.closeModal = this.closeModal.bind(this);

        }
    componentWillMount(){
        this.props.getOrders();
        this.props.getUserOrders();
        this.props.getOrderBalanceOrders('DOGE');
        this.props.getOrderBalanceOrders('BTC');
        this.props.getOrderBalanceOrders('LTC');

    }
    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;
        let currAmount;
        let minCurrAmount;
        let totalVali = null;
        var currencyActual = null;
        var orderValidation = 0;

        if(name === 'currency'){
            let that = this;
            this.props.validationCur.forEach(element => {
                if(element.currency === value){                    
                    if(this.props['orders'+element.currency] != ""){
                        orderValidation = this.props['orders'+element.currency]
                    }else{
                        orderValidation = 0;
                    }
                    document.getElementById('curr').innerHTML = (parseFloat(element.amount)-parseFloat(orderValidation)) + ' '+element.currency + ' available';
                    currAmount = element.amount;
                    if(value==="DOGE"){
                        minCurrAmount = "3";
                    }else if(value==="LTC"){
                        minCurrAmount="0.006"
                    }else if(value==="BTC"){
                        minCurrAmount="0.001060"
                    }
                }
            });
            let finalVali = (parseFloat(currAmount)-parseFloat(orderValidation));
            if(value==="DOGE"){
                minCurrAmount = "3";
            }else if(value==="LTC"){
                minCurrAmount="0.006"
            }else if(value==="BTC"){
                minCurrAmount="0.001060"
            }
            this.setState({
                actualBalance:finalVali,
                maxCURR:currAmount,
                minCURR:minCurrAmount,
            })
        }
        let balance = 0;
        if((name === 'amount') && (value != 0)){
            balance = value-3
        }
        this.setState({
            [name]: value,
            total:  balance
        });
      }
    handleSubmit(event) {
        event.preventDefault();
        console.log(this.state.actualBalance);
        console.log(this.state.amount);
        if(this.state.amount > this.state.actualBalance){
            alert("Can't do this");
        }else{
            this.props.addNewOrder(this.state);
        }
        
        this.closeModal();
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
    render() {
        if(this.props.buy){
            var orders = this.props.buy.map((order,index)=>{
                if(order.amount > 0){
                    let amount = Math.round(order.amount * 100)/100;
                    let price = Math.round(order.price * 10000000000)/10000000000;
                    return <OrderComponent props={{amount:amount,currency:order.currency,price:price,id:order.id,currency_to:order.currency_to}}/>
                }
            });
        }
        if(this.props.userOrders){
            var userOrders = this.props.userOrders.map((order,index)=>{
                if(order.amount > 0){
                    let amount = Math.round(order.amount * 100)/100;
                    let price = Math.round(order.price * 100)/100;
                    return <OrderComponent props={{amount:amount,currency:order.currency,price:price,id:order.id,currency_to:order.currency_to}}/>
                }
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
                <div className="dash dashboard__order buy__resposive">
                    {orders}
                </div>
                <div className="dash dashboard__sell responsive__user">
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
                    <div>
                        <h1>Make your order</h1>
                        <button onClick={this.openModal}>Open Modal</button>
                    </div>
                    <Modal
                    isOpen={this.state.modalIsOpen}
                    onAfterOpen={this.afterOpenModal}
                    onRequestClose={this.closeModal}
                    style={customStyles}
                    contentLabel="Example Modal">
                    <div className="modal__pattern">
                        <div className="modal__header">
                            <span>Create order</span>
                        </div>
                        <div className="modal__body">
                            <form onSubmit={this.handleSubmit}>
                                <div className="option__sale">
                                    <label>
                                        <input className="form__input" name="currency" type="radio" value="BTC" onChange={this.handleInputChange} />
                                        <img src="http://localhost:8080/static/bitcoin.svg" width="30" />
                                    </label>
                                    <label>
                                        <input className="form__input" name="currency" type="radio" placeholder="Price" value="DOGE" onChange={this.handleInputChange}/>
                                        <img src="http://localhost:8080/static/dogecoin.svg" width="30"/>
                                    </label>
                                    <label>
                                        <input className="form__input" name="currency" type="radio" placeholder="Price" value="LTC" onChange={this.handleInputChange}/>
                                        <img src="http://localhost:8080/static/litecoin.svg" width="30"/>
                                    </label>
                                </div>
                                <span id="curr" className="align__aviable">Choose one currency</span>
                                <div className="input__pattert">
                                    <input className="form__input" name="amount" type="number" placeholder="Amount" min={this.state.minCURR} max={this.state.maxCURR} step="0.001" onChange={this.handleInputChange}/>
                                </div>
                                <div className="input__pattert">
                                    <input className="form__input" name="price" type="number" placeholder="Price" step="0.001" min="0" onChange={this.handleInputChange}/>
                                </div>
                                <div>
                                    <input className="form__input" name="currency_to" type="text" placeholder="Currency" onChange={this.handleInputChange}/>
                                </div>
                                <div>
                                    <div className="modal__pricing">Total to sell: <span className="modal__total modal__total--color">{this.state.total}</span> {this.state.currency}</div>
                                </div>
                                <button className="form__button" type="submit">Create Order</button>
                            </form>
                        </div>
                    </div>
                    </Modal>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { 
                buy: state.buy.buy,
                userOrders:state.buy.userOrder,
                newOrder:state.buy.newOrder,
                validationCur: state.overview.overview,
                ordersDOGE: state.buy.ordersDOGE,
                ordersLTC: state.buy.ordersLTC,
                ordersBTC: state.buy.ordersBTC,
            }
}

export default connect(mapStateToProps, actions)(OrdersContainer);
