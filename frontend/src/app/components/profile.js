import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class Profile extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            username: this.props.profileData.username,
            email: this.props.profileData.email,
            firstname: this.props.profileData.firstname,
            surname: this.props.profileData.surname,
            mobile_pin:this.props.profileData.mobile_pin
          };
        this.handleChange = this.handleChange.bind(this);
    }
    handleFormSubmit(e) {
        e.preventDefault()
        let fistname = e.target.fistname.value;
        let surname = e.target.surname.value;
        let send = {fistname:fistname,surname:surname}
        this.props.updateProfile(send);
    }
    handleChange(event) {
        let name = event.target.name;
        this.setState({[name]: event.target.value});
    }
    render() {
        console.log(this.props)
        return (
            <div className="profile">
                <div className="profile__box">
                    <div className="box__header">
                        <img src="http://localhost:8080/static/avatar.png" className="header__avatar"></img>
                        <span className="header__username">{this.state.username}</span>
                        
                    </div>
                    <div className="box__body">
                        <form className="body__form" onSubmit={this.handleFormSubmit.bind(this)}>
                            <div className="input__pattert">
                                <input className="form__input" type="text" name="username" placeholder="Username" disabled value={this.state.username}/>
                            </div>
                            <div className="input__pattert">
                                <input className="form__input" type="text" name="email" placeholder="Email" disabled value={this.state.email}/>
                            </div>
                            <div className="input__pattert">
                                <input className="form__input" type="text" name="fistname" placeholder="Name" defaultValue={this.state.firstname} onChange={this.handleChange}/>
                            </div>
                            <div className="input__pattert">
                                <input className="form__input" type="text" name="surname" placeholder="Surname" defaultValue={this.state.surname} onChange={this.handleChange}/>
                            </div>
                            <div>
                                {this.state.mobile_pin}
                            </div>
                            <div className="form__button__pattern">
                                <button className="form__button" action="submit" >Update Profile</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    console.log(state)
    return { datos: state.auth.profileData }
}

export default connect(mapStateToProps, actions)(Profile);