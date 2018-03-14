import React, { Component } from 'react';
import Header from './header';
import Footer from './footer';
class App extends Component {
  render() {
    return (
      <div>
        <Header />
        <div className="cont">
              {this.props.children}
        </div>
        <Footer />
      </div>
    );
  }
}

export default App;
