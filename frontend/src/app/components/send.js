import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import ReactTable from "react-table";
import Modal from 'react-modal';
import AddressesComponent from './send/tabContent';
import GestBalanceComponent from './send/GestBalance';
class Send extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            isThorIn: true,
            isHulkIn: true,
            isIronmanIn: true,
            isBTC:true,
            isDOGE:true,
            isLTC:true,
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
                    return address.address
                }
            });
            dogecoinAddresses = this.props.addresses.map((address, index)=>{
                if(address.currency === "DOGE"){
                    return address.address
                }
            });
            litecoinAddresses = this.props.addresses.map((address,index)=>{
                if(address.currency === "LTC"){
                    return address.address
                }
            });
        }else{
            bitcoinAddresses = 'Loading...'
            dogecoinAddresses = 'Loading...';
            litecoinAddresses = 'Loading...';
        }
        let bitcoinLabel = null;
        let dogecoinLabel = null;
        let litecoinLabel = null;
        if(this.props.addresses){
            bitcoinLabel = this.props.addresses.map((address, index)=>{
                if(address.currency === "BTC"){
                    return address.label
                }
            });
            dogecoinLabel = this.props.addresses.map((address, index)=>{
                if(address.currency === "DOGE"){
                    return address.label
                }
            });
            litecoinLabel = this.props.addresses.map((address,index)=>{
                if(address.currency === "LTC"){
                    return address.label
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
                <div className="send">
                <Tabs>
                    <div className="send--center">
                        <TabList className="react-tabs__tab-list">
                            {this.state.isBTC && <Tab><img src="http://localhost:8080/static/bitcoin.svg" alt="Bitcoin" height="32" width="32" /> Bitcoin</Tab>}
                            {this.state.isDOGE && <Tab><img src="http://localhost:8080/static/dogecoin.png" alt="Dogecoin" height="32" width="32" /> Dogecoin</Tab>}
                            {this.state.isLTC && <Tab><img src="http://localhost:8080/static/litecoin.svg" alt="Litecoin" height="32" width="32" /> Litecoin</Tab>}
                        </TabList>
                    </div>
                    <div className="send__card">
                    {
                        this.state.isBTC && 
                        <TabPanel className="tab__display">
                            <GestBalanceComponent props={{currency:'BTC'}}/>
                            {
                                this.props.addresses ? <AddressesComponent props={{address:bitcoinAddresses[0],label:bitcoinLabel[0],currency:'BTC'}}/> : 
                                <div className="address__pane">
                                    Loading...
                                </div>
                            }
                        </TabPanel>
                    }
                    {
                        this.state.isDOGE && 
                        <TabPanel className="tab__display">
                            <GestBalanceComponent props={{currency:'DOGE'}}/>
                            {
                                this.props.addresses ? <AddressesComponent props={{address:dogecoinAddresses[1],label:dogecoinLabel[1],currency:'DOGE'}}/> : 
                                <div className="address__pane">
                                    Loading...
                                </div>
                            }
                        </TabPanel>
                    }
                    {
                        this.state.isLTC && 
                        <TabPanel className="tab__display">
                            <GestBalanceComponent props={{currency:'LTC'}}/>
                            {
                                this.props.addresses ? <AddressesComponent props={{address:litecoinAddresses[2],label:litecoinLabel[2],currency:'LTC'}}/> : 
                                <div className="address__pane">
                                    Loading...
                                </div>
                            }
                        </TabPanel>
                    }
                    </div>
                </Tabs>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    console.log(state)
    return { addresses: state.send.addresses }
}

export default connect(mapStateToProps, actions)(Send);
