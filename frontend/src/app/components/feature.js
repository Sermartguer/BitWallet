import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import Profile from './profile';
class ProfileMiddleware extends PureComponent {
    componentWillMount(){
        this.props.getProfileData();
    }
    render() {
        if(this.props.datos){
            console.log(this.props.datos)
            var profile = <Profile profileData={{username:this.props.datos.username,email:this.props.datos.email,firstname:this.props.datos.firstname,surname:this.props.datos.surname,mobile_pin:this.props.datos.mobile_hash}}></Profile>
        }else{
            var profile = 'Loading...'
        }
        return (
            <div>
                {profile}
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { datos: state.auth.profileData }
}

export default connect(mapStateToProps, actions)(ProfileMiddleware);
