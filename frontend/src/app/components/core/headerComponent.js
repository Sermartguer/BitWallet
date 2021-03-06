import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Grid, Row, Col } from 'react-flexbox-grid';

class Header extends PureComponent {
    
    renderLinks() {
        if (this.props.authenticated) {
            return [
            <div key="logo" className="">
                <li key="logo" className="navbar__item">
                    <Link to='/' className="navbar__item title--color"><img src="http://localhost:8080/static/bitwallet_logo5.png" width="100"></img></Link>
                </li>
            </div>,
            <div key="menu" className="navbar__items">
                <li key="dashboard" className=" navbar__item">
                    <Link className="navbar__item--color" to='/dashboard'>Dashboard</Link>
                </li>
                <li key="username" className="navbar__item">
                    <Link className="navbar__item--color" to="/profile">{this.props.username}</Link>
                </li>
                <li key="logout" className="navbar__item">
                    <Link className="navbar__item--color" to="/signout" >Logout</Link>
                </li>
            </div>
            ];
        } else {
            return [
                <div key="logo" className="">
                    <li key="logo" className="navbar__item">
                        <Link to='/' className="navbar__item title--color">BitWallet</Link>
                    </li>
                </div>,
                <div key="menu" className="navbar__items">
                    <li key="home" className=" navbar__item">
                        <Link className="navbar__item--color" to='/'>Home</Link>
                    </li>
                    <li  key="about" className="navbar__item">
                        <Link className="navbar__item--color" to='/about'>About</Link>
                    </li>
                    <li key="login" className="navbar__item">
                        <Link className="navbar__item--color" to="/signin">Login</Link>
                    </li>
                </div>
            ];
        }
    }

    render() {
        console.log(this.props)
        return (
            <header>
                <nav>
                <Row className="navbar">
                    {this.renderLinks()}
                </Row>
                </nav>
            </header>
        );
    }
}

const mapStateToProps = (state) => {
    return { authenticated: state.auth.authenticated,username: state.auth.username }
};

export default connect(mapStateToProps)(Header);
