import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import ReactTable from "react-table";
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

class Send extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            isThorIn: true,
            isHulkIn: true,
            isIronmanIn: true,
            modalIsOpen: false,
            currencyModalActive: 'Nothing'
        };
        this.onButtonClick = this.onButtonClick.bind(this);
        this.handleCheckClicked = this.handleCheckClicked.bind(this);
        this.openModal = this.openModal.bind(this);
        this.closeModal = this.closeModal.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);


    }
    componentWillMount(){
        this.props.getUserAddresses();
    }
    onButtonClick(e) {
        e.preventDefault();
        this.props.addNewAddress(e.target.id)

      }
      handleCheckClicked(e) {
        this.setState({
          [e.target.name]: e.target.checked,
        });
      }
    openModal(e) {
        this.setState({modalIsOpen: true,currencyModalActive:e.target.id});
        console.log(e.target.id)
    }
    closeModal() {
        this.setState({modalIsOpen: false});
    }
    handleSubmit(event) {
        event.preventDefault();
        console.log(event.target.label.value)
        console.log(this.state.currencyModalActive)
        this.props.addNewAddress({currency: this.state.currencyModalActive, label:event.target.label.value})
        //this.props.addNewOrder(this.state);
        this.closeModal();
    }
      render() {
        let bitcoinAddresses = null;
        let dogecoinAddresses = null;
        let litecoinAddresses = null;
        if(this.props.addresses){
            bitcoinAddresses = this.props.addresses.map((address, index)=>{
                if(address.currency === "BTC"){
                    return <span key={index}>{address.address} {address.label}</span>
                }
            });
            dogecoinAddresses = this.props.addresses.map((address, index)=>{
                if(address.currency === "DOGE"){
                    return <span key={index}>{address.address} {address.label}</span>
                }
            });
            litecoinAddresses = this.props.addresses.map((address,index)=>{
                if(address.currency === "LTC"){
                    return <span key={index}>{address.address} {address.label}</span>
                }
            });
        }else{
            bitcoinAddresses = 'Loading...'
            dogecoinAddresses = 'Loading...';
            litecoinAddresses = 'Loading...';
        }
        if(litecoinAddresses[0] === null){
            litecoinAddresses = <span onClick={this.onButtonClick} id="LTC">No Litecoin Addresses yet, click to create...</span>
        }
        if(dogecoinAddresses[0] === null){
            dogecoinAddresses = <span onClick={this.onButtonClick} id="DOGE">No Dogecoin Addresses yet, click to create...</span>
        }
        if(bitcoinAddresses[0] === null){
            bitcoinAddresses = <span onClick={this.onButtonClick} id="BTC">No Bitcoin Addresses yet, click to create...</span>
        }
        return (
            <div className="dash dash__send">
                <div className="dash__send--size">
                    <div className="pane pane__currency">
                        <Tabs>
                            <TabList className="react-tabs__tab-list">
                                {this.state.isThorIn && <Tab><img src="http://localhost:8080/static/bitcoin.svg" alt="Bitcoin" height="32" width="32" /></Tab>}
                                {this.state.isHulkIn && <Tab><img src="http://localhost:8080/static/dogecoin.svg" alt="Hulk" height="32" width="32" /></Tab>}
                                {this.state.isIronmanIn && <Tab><img src="http://localhost:8080/static/litecoin.svg" alt="Ironman" height="32" width="32" /></Tab>}
                            </TabList>
                            {this.state.isThorIn && 
                            <TabPanel>
                                <div className="address__pane">
                                    <div className="add__address">
                                        <span className="add__title">Your BTC Addresses</span>
                                        <span className="add__cursor" onClick={this.openModal} id="BTC"><i className="far fa-plus-square"></i> Add Address</span>
                                        {bitcoinAddresses}
                                    </div>
                                </div>
                            </TabPanel>}
                            {this.state.isHulkIn && 
                            <TabPanel>
                                <div className="address__pane">
                                    <div className="add__address">
                                        <span className="add__title">Your Dogecoin Addresses</span>
                                        <span className="add__cursor" onClick={this.openModal} id="DOGE"><i className="far fa-plus-square"></i> Add Address</span>
                                        {dogecoinAddresses}
                                    </div>
                                </div>
                            </TabPanel>}
                            {this.state.isIronmanIn && 
                            <TabPanel>
                                <div className="address__pane">
                                    <div className="add__address">
                                        <span className="add__title">Your LTC Addresses</span>
                                        <span className="add__cursor" onClick={this.openModal} id="LTC"><i className="far fa-plus-square"></i> Add Address</span>
                                        {litecoinAddresses}
                                    </div>
                                </div>
                            </TabPanel>}
                        </Tabs>
                    </div>
                </div>
                <Modal
                    isOpen={this.state.modalIsOpen}
                    onAfterOpen={this.afterOpenModal}
                    onRequestClose={this.closeModal}
                    style={customStyles}
                    contentLabel="Example Modal">
                        <span>Pon nombre a tu direccion</span>
                        <form onSubmit={this.handleSubmit}>
                            <input type="text" name="label"></input>
                            <button className="form__button" type="submit">Create Order</button>
                        </form>
                        
                    </Modal>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    console.log(state)
    return { addresses: state.send.addresses }
}

export default connect(mapStateToProps, actions)(Send);
