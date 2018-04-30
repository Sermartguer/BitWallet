import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import Orders from './store/orders';
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

class Buy extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            currency:'DOGE',
            amount:0,
            price:0,
            newOrders:false,
            modalIsOpen: false,
            maxCURR:0
          };
          this.handleInputChange = this.handleInputChange.bind(this);
          this.handleSubmit = this.handleSubmit.bind(this);
          this.openModal = this.openModal.bind(this);
          this.closeModal = this.closeModal.bind(this);

        }
    componentWillMount(){
        this.props.getOrders();
        this.props.getUserOrders();
    }
    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;
        let currAmount;
        if(name === 'currency'){
            let that = this;
            console.log(this.props.validationCur)
            this.props.validationCur.forEach(element => {
                if(element.currency === value){
                    document.getElementById('curr').innerHTML = element.amount + ' '+element.currency + ' available';
                    currAmount = element.amount
                }
            });
        }
        this.setState({
            maxCURR:currAmount,
            [name]: value
        });
      }
    handleSubmit(event) {
        event.preventDefault();
        this.props.addNewOrder(this.state);
        this.closeModal();
    }
    openModal() {
        this.setState({modalIsOpen: true});
      }
      closeModal() {
        this.setState({modalIsOpen: false});
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
        console.log(this)
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
                                    <input className="form__input" name="amount" type="number" placeholder="Amount" min="0" max={this.state.maxCURR} step="0.001" onChange={this.handleInputChange}/>
                                </div>
                                <div className="input__pattert">
                                    <input className="form__input" name="price" type="number" placeholder="Price" min="0" onChange={this.handleInputChange}/>
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
                validationCur: state.overview.overview
            }
}

export default connect(mapStateToProps, actions)(Buy);
