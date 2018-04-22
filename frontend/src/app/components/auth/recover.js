import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';

class Recover extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            email:''
          };
          this.handleInputChange = this.handleInputChange.bind(this);
          this.handleSubmit = this.handleSubmit.bind(this);      
        }
    handleSubmit(event) {       
        event.preventDefault();
        console.log(this.state)
        this.props.recoverPassword(this.state);
    }
    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;
    
        this.setState({
          [name]: value
        });
      }
    render() {
        return (
            <div className="login">
                 <div className="login__modal">
                    <span className="login__title recover__title">Recover password</span>
                    <span>Enter email to recover your password</span>
                    <form className="modal" onSubmit={this.handleSubmit}>
                        <div className="input__pattert">
                            <input className="form__input" name="email"  type="text" defaultValue={this.state.email} onChange={this.handleInputChange} placeholder="Email"/>
                        </div>
                        <div className="form__button__pattern">
                            <button className="form__button" action="submit" >Sign in</button>
                        </div>
                    </form>
                </div>
            </div>
        )
    }
}

export default connect(null, actions)(Recover);
