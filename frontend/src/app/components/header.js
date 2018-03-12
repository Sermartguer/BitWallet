import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

class Header extends PureComponent {
    
    renderLinks() {
        if (this.props.authenticated) {
            return [
            <li key="logo" className="col--6 col--md--5 col--sm--4 col--xs--12 col--ps--12 navbar__item">
                <span className="navbar__item title--color">BitWallet</span>
            </li>,
            <li key="home" className="col--1 col--md--1 col--sm--1 col--xs--12 col--ps--12 navbar__item">
                <Link className="navbar__item--color" to='/'>Home</Link>
            </li>,
            <li  key="dashboard" className="col--1 col--md--1 col--sm--1 col--xs--12 col--ps--12 navbar__item">
                <Link className="navbar__item--color" to='/dashboard'>Dashboard</Link>
            </li>,
            <li key="username" className="col--1 col--md--1 col--sm--1 col--xs--12 col--ps--12 navbar__item">
                <Link className="navbar__item--color" to="#">{this.props.username}</Link>
            </li>,
            <li key="logout" className="col--1 col--md--1 col--sm--1 col--xs--12 col--ps--12 navbar__item">
                <Link className="navbar__item--color" to="/signout" >Logout</Link>
            </li>
            ];
        } else {
            return [
                <li key="logo" className="col--6 col--md--5 col--sm--4 col--xs--12 col--ps--12 navbar__item">
                    <span className="navbar__item title--color">BitWallet</span>
                </li>,
                <li  key="home" className="col--1 col--md--1 col--sm--1 col--xs--12 col--ps--12 navbar__item">
                    <Link className="navbar__item--color" to='/'>Home</Link>
                </li>,
                <li key="about" className="col--1 col--md--1 col--sm--1 col--xs--12 col--ps--12 navbar__item">
                    <Link className="navbar__item--color" to='/about'>About</Link>
                </li>,
                <li key="login" className="col--1 col--md--1 col--sm--1 col--xs--12 col--ps--12 navbar__item">
                    <Link className="navbar__item--color" to="/signin">Login</Link>
                </li>
            ];
        }
    }

    render() {
        console.log(this.props)
        return (
            <header>
                <nav>
                <ul className="row center center-md center-sm center-xs center-ps navbar">
                    {this.renderLinks()}
                </ul>
                </nav>
            </header>
        );
    }
}

const mapStateToProps = (state) => {
    return { authenticated: state.auth.authenticated,username: state.auth.username }
};

export default connect(mapStateToProps)(Header);
