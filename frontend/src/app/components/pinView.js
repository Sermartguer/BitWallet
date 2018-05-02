import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
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
    render() {
        return (
            <div className="dash">
            <PinInput 
            length={4}
            onChange={(value, index) => { }} 
            type="numeric" 
            style={{padding: '10px'}}  
            inputStyle={{borderColor: 'red'}}
            inputFocusStyle={{borderColor: 'blue'}}
            onComplete={(value, index) => {
                this.setState({pin:value})
                this.sendPin()
            }}
            />
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { }
}

export default connect(mapStateToProps, actions)(PinView);
