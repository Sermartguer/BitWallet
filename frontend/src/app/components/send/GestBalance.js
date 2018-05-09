import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import ReactTable from "react-table";
import Modal from 'react-modal';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import PinInput from 'react-pin-input';

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
class GestBalanceComponent extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            currency:props.props.currency,
            isLocal:true,
            isExternal:true,
            pin:"",
            amount:"",
            to:"",
            sendTo:""
        };
        this.openModal = this.openModal.bind(this);
        this.closeModal = this.closeModal.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleInputChange = this.handleInputChange.bind(this);

    }
    openModal(e) {
        console.log(e.target.name)
        this.setState({modalIsOpen: true,currencyModalActive:e.target.id,sendTo:e.target.name});
        
    }
    closeModal() {
        this.setState({modalIsOpen: false});
    }
    handleSubmit(event) {
        event.preventDefault();
        if(this.state.sendTo === 'local'){
            this.props.sendLocalBalances({"currency":this.state.currency,"amount":this.state.amount,"to":this.state.to,"pin":this.state.pin})
        }else{
            this.props.sendExternalBalances({"currency":this.state.currency,"amount":this.state.amount,"to":this.state.to,"pin":this.state.pin})
        }
        this.closeModal();
        console.log(this.state)
    }
    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;
    
        this.setState({
          [name]: value
        });
      }
      renderError() {
        if (this.props.errorMessage) {
            console.log(this.props.errorMessage)
            return (
                <div className="alert__danger">
                    <string>Error! {this.props.errorMessage}</string>
                </div>
            );
        }
    }
    renderSuccess() {
        if (this.props.successMessage) {
            return (
                <div className="">
                    <string>Success! View transaction<a target="_blank" href={'https://chain.so/tx/DOGETEST/'+this.props.successMessage.txid}>here</a></string>
                </div>
            );
        }
    }
    render() {
        return (
            <div className="send__pane">
                <span>Send {this.state.currency}</span>
                <Tabs>
                    <TabList className="react-tabs__tab-list tab__content">
                        {this.state.isLocal && <Tab>Local</Tab>}
                        {this.state.isExternal && <Tab>External</Tab>}
                    </TabList>
                    <TabPanel className="tab__display tab__send">
                        <div className="tab__cont">
                            <span>BitWallet ID:</span>
                            <input type="text" className="form__input form__input--color" name="to" placeholder="BitWallet ID" onChange={this.handleInputChange}></input>
                        </div>
                        <div className="tab__cont">
                            <span>Amount:</span>
                            <input type="text" className="form__input form__input--color" name="amount" type="number" placeholder="Amount" onChange={this.handleInputChange} step="0.001"></input>
                        </div>
                        <div>
                            {this.renderError()}
                            {this.renderSuccess()}
                        </div>
                        <div>
                            <button className="form__btn" id={this.state.currency} name="local"  onClick={this.openModal}>Send {this.state.currency}</button>
                        </div>
                    </TabPanel>
                    <TabPanel className="tab__display tab__send">
                        <div className="tab__cont">
                            <span>Address:</span>
                            <input type="text" className="form__input form__input--color" name="to" placeholder={this.state.currency+' Address'} onChange={this.handleInputChange}></input>
                        </div>
                        <div className="tab__cont">
                            <span>Amount:</span>
                            <input type="text" className="form__input form__input--color" name="amount" type="number" placeholder="Amount" onChange={this.handleInputChange} step="0.001"></input>
                        </div>
                        <div>
                            {this.renderError()}
                            {this.renderSuccess()}
                        </div>
                        <div>
                            <button className="form__btn" type="button" id={this.state.currency} name="external" onClick={this.openModal}>Send {this.state.currency}</button>
                        </div>
                    </TabPanel>
                </Tabs>
                <Modal
                    isOpen={this.state.modalIsOpen}
                    onAfterOpen={this.afterOpenModal}
                    onRequestClose={this.closeModal}
                    style={customStyles}
                    contentLabel="Example Modal">
                        <span>Total</span>
                        <form onSubmit={this.handleSubmit}>
                            total
                            <PinInput 
                                length={4}
                                onChange={(value, index) => { }} 
                                type="numeric" 
                                style={{padding: '10px'}}  
                                inputStyle={{borderColor: 'red'}}
                                inputFocusStyle={{borderColor: 'blue'}}
                                onComplete={(value, index) => {
                                    this.setState({pin:value})
                                }}
                            />
                            <button className="form__btn" type="submit">Create Order</button>
                        </form>
                    </Modal>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    console.log(state)
    return { 
        trans: state.overview.transactions,
        errorMessage:state.send.error,
        successMessage:state.send.success
     }
}

export default connect(mapStateToProps, actions)(GestBalanceComponent);
