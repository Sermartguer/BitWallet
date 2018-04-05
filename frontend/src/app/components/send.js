import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class Send extends PureComponent {
    constructor(props){
        super(props);
        this.onButtonClick = this.onButtonClick.bind(this);
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
                    return <span className="address_info" key={index}>{address.address}</span>
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
            <div className="dash sendView">
                <div className="pane address__pane">
                    <div className="add__address">
                        <span className="add__title">Your BTC Addresses</span>
                        <span className="add__cursor" onClick={this.onButtonClick} id="BTC"><i className="far fa-plus-square"></i> Add Address</span>
                        {bitcoinAddresses}
                    </div>
                </div>
                <div className="pane address__pane">
                    <div className="add__address">
                        <span className="add__title">Your Dogecoin Addresses</span>
                        <span className="add__cursor" onClick={this.onButtonClick} id="DOGE"><i className="far fa-plus-square"></i> Add Address</span>
                        {dogecoinAddresses}
                    </div>
                </div>
                <div className="pane address__pane">
                    <div className="add__address">
                        <span className="add__title">Your LTC Addresses</span>
                        <span className="add__cursor" onClick={this.onButtonClick} id="LTC"><i className="far fa-plus-square"></i> Add Address</span>
                        {litecoinAddresses}
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
