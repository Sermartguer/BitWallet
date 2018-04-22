import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';

class NewPassword extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            password:'',
            repassword:'',
            id:this.props.match.params.id
          };
          this.handleInputChange = this.handleInputChange.bind(this);
          this.handleSubmit = this.handleSubmit.bind(this);      
        }
    handleSubmit(event) {       
        event.preventDefault();
        this.props.newPassword(this.state)
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
    render() {
        return (
            <div className="login">
                 <div className="login__modal">
                    <span className="login__title recover__title">New password</span>
                    <span>Please enter new password</span>
                    <form className="modal" onSubmit={this.handleSubmit}>
                        <div className="input__pattert">
                            <input className="form__input" name="password" type="password" defaultValue={this.state.email} onChange={this.handleInputChange} placeholder="Password"/>
                        </div>
                        <div className="input__pattert">
                            <input className="form__input" name="repassword" type="password" defaultValue={this.state.email} onChange={this.handleInputChange} placeholder="Repeat Password"/>
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

export default connect(null, actions)(NewPassword);
