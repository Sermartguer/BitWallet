import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../actions';
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
class AddressesComponent extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            address:props.props.address,
            label:props.props.label,
            modalIsOpen: false,
            currencyModalActive: 'Nothing',
            currency:props.props.currency
        };
        this.openModal = this.openModal.bind(this);
        this.closeModal = this.closeModal.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
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
        return (
                <div className="address__pane">
                    <div className="add__address">
                        <span className="add__title">Your Dogecoin Addresses</span>
                        <span className="add__cursor" onClick={this.openModal} id={this.state.currency}><i id={this.state.currency} className="far fa-plus-square"></i> Add Address</span>
                        Your dogecoin public address:
                        <span className="body__pin">{this.state.address}</span>
                        Your BitWallet criptocurrency id:
                        <span className="body__pin">{this.state.label}</span>
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
    return { 
        trans: state.overview.transactions
     }
}

export default connect(mapStateToProps, actions)(AddressesComponent);
