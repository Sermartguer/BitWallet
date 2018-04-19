import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class Profile extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            currencyDetail: this.props.match.params.id
          };
    }
    handleFormSubmit(e) {
        e.preventDefault()
        let fistname = e.target.fistname.value;
        let surname = e.target.surname.value;
        let send = {fistname:fistname,surname:surname}
        this.props.updateProfile(send);
    }
    render() {
        return (
            <div className="profile">
                <div className="profile__box">
                    <div className="box__header">
                        <img src="http://localhost:8080/static/avatar.png" className="header__avatar"></img>
                        <span className="header__username">Sermartguer</span>
                    </div>
                    <div className="box__body">
                        <form className="body__form" onSubmit={this.handleFormSubmit.bind(this)}>
                            <div className="input__pattert">
                                <input className="form__input" type="text" name="username" placeholder="Username" disabled/>
                            </div>
                            <div className="input__pattert">
                                <input className="form__input" type="text" name="email" placeholder="Email" disabled/>
                            </div>
                            <div className="input__pattert">
                                <input className="form__input" type="text" name="fistname" placeholder="Name" />
                            </div>
                            <div className="input__pattert">
                                <input className="form__input" type="text" name="surname" placeholder="Surname" />
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
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Profile);