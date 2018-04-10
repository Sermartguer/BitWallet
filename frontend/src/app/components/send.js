import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import ReactTable from "react-table";

class Send extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            isThorIn: true,
            isHulkIn: true,
            isIronmanIn: true
        };
        this.onButtonClick = this.onButtonClick.bind(this);
        this.handleCheckClicked = this.handleCheckClicked.bind(this);

    }
    componentWillMount(){
        this.props.getUserAddresses();
    }
    renderCurrency() {
        if (this.props.addresses) {
            console.log(his.props.addresses)
            return (
                    <span>a</span>
            );
        }
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
    render() {
        let bitcoinAddresses = null;
        let dogecoinAddresses = null;
        let litecoinAddresses = null;
        if(this.props.addresses){
                        bitcoinAddresses = this.props.addresses.sort()
                        .reverse().map((address, index)=>{
                if(address.currency === "BTC"){
                    return <span key={index}>{address.address}</span>
                }else{
                    return null;
                }
            });
            dogecoinAddresses = this.props.addresses.sort()
            .reverse().map((address, index)=>{
                if(address.currency === "DOGE"){
                    return <span>{address.address}</span>
                }else{
                    return null;
                }
            });
            litecoinAddresses = this.props.addresses.sort()
            .reverse().map((address,index)=>{
                if(address.currency === "LTC"){
                    return <span key={index}>{address.address}</span>
                }else{
                    return null;
                }
            });
        }else{
            bitcoinAddresses = 'Loading...'
            dogecoinAddresses = [];
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
        console.log({dogecoinAddresses})
        return (
            <div className="dash dash__send">
            <div className="dash__send--size">
                <div className="pane pane__currency">
                    <Tabs>
                        <TabList className="react-tabs__tab-list">
                            {this.state.isThorIn && <Tab><img src="https://cdn.worldvectorlogo.com/logos/bitcoin.svg" alt="Bitcoin" height="32" width="32" /></Tab>}
                            {this.state.isHulkIn && <Tab><img src="https://cdn.worldvectorlogo.com/logos/dogecoin.svg" alt="Hulk" height="32" width="32" /></Tab>}
                            {this.state.isIronmanIn && <Tab><img src="https://cdn.worldvectorlogo.com/logos/litecoin.svg" alt="Ironman" height="32" width="32" /></Tab>}
                        </TabList>
                        {this.state.isThorIn && 
                        <TabPanel>
                            <div className="address__pane">
                                <div className="add__address">
                                    <span className="add__title">Your BTC Addresses</span>
                                    <span className="add__cursor" onClick={this.onButtonClick} id="BTC"><i className="far fa-plus-square"></i> Add Address</span>
                                    {bitcoinAddresses}
                                </div>
                            </div>
                        </TabPanel>}
                        {this.state.isHulkIn && 
                        <TabPanel>
                            <div className="address__pane">
                                <div className="add__address">
                                    <span className="add__title">Your Dogecoin Addresses</span>
                                    <span className="add__cursor" onClick={this.onButtonClick} id="DOGE"><i className="far fa-plus-square"></i> Add Address</span>
                                    {dogecoinAddresses}
                                </div>
                            </div>
                        </TabPanel>}
                        {this.state.isIronmanIn && 
                        <TabPanel>
                            <div className="address__pane">
                                <div className="add__address">
                                    <span className="add__title">Your LTC Addresses</span>
                                    <span className="add__cursor" onClick={this.onButtonClick} id="LTC"><i className="far fa-plus-square"></i> Add Address</span>
                                    {litecoinAddresses}
                                </div>
                            </div>
                        </TabPanel>}
                    </Tabs>
                </div>
            </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { addresses: state.send.addresses }
}

export default connect(mapStateToProps, actions)(Send);
