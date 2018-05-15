import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import PinInput from 'react-pin-input';
class PinView extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            param: this.props.match.params.id,
            pin: ''
          };
    }
    sendPin(){
        console.log('he')
        this.props.verifyAccount(this.state);
    }
    renderError() {
        if (this.props.error) {
            return (
                <div className="alert__danger">
                    <string>{this.props.error}</string>
                </div>
            );
        }
    }
    renderSuccess() {
        if (this.props.success) {
            return (
                <div className="success__text">
                    <string>{this.props.success} <a href="/">Click to return</a></string>
                </div>
            );
        }
    }
    render() {
        return (
            <div className="pin">
                <div className="pin__view">
                    <span className="pin__view__title pin__view__title--color">Verify account</span>
                    <span className="pin__view__subtitle pin__view__subtitle--color">Insert your account pin, this pin will be used when you withdraw funds, create addresses ...</span>
                    <PinInput 
                    length={4}
                    onChange={(value, index) => { }} 
                    type="numeric" 
                    style={{padding: '10px',display: 'flex',justifyContent: 'center',}}  
                    inputStyle={{borderColor: 'red'}}
                    inputFocusStyle={{borderColor: 'blue'}}
                    onComplete={(value, index) => {
                        this.setState({pin:value})
                        this.sendPin()
                    }}
                    />
                    {this.renderError()}
                    {this.renderSuccess()}
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => ({...state.auth})

export default connect(mapStateToProps, actions)(PinView);
