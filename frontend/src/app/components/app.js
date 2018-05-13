import React, { Component } from 'react';
import Header from './core/headerComponent';
import Footer from './core/footerComponent';
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
