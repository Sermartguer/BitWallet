import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';
import OverviewContainer from './overview/overviewContainer';
import OrdersContainer from './store/ordersContainer';
import SendContainer from './send/sendContainer';
import History from './history/historyContainer';
class Dashboard extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            component: OverviewContainer,
            currencyDetail: ''
          };
        this.onButtonClick = this.onButtonClick.bind(this);
    }
    onButtonClick(e) {
        console.log(e.target.id)
        let component;
        if(e.target.id === 'Overview'){
            component = OverviewContainer
        }else if(e.target.id === 'Buy'){
            component = OrdersContainer
        }else if(e.target.id === 'Send'){
            component = SendContainer
        }else if(e.target.id === 'History'){
            component = History
        }
        this.setState({
            component: component,
        });
      }
    render() {
        return (
            <div className="dash">
                <div className="dash__items">
                    <section className="items__dashboard">
                    <br/>

                        <span className="alert__verify"><i className="fa fa-info-circle" aria-hidden="true"></i> Please verify your account</span>
                        <div className="item__dashboard" onClick={this.onButtonClick} id="Overview"><i className="fas fa-chart-line space--icons"></i>Overview</div>
                        <div className="item__dashboard" onClick={this.onButtonClick} id="Buy"><i className="far fa-handshake space--icons"></i>Exchange</div>
                        <div className="item__dashboard" onClick={this.onButtonClick} id="Send"><i className="fab fa-telegram-plane space--icons"></i>Send</div>
                        <div className="item__dashboard" onClick={this.onButtonClick} id="History"><i className="fas fa-history space--icons"></i>History</div>
                    </section>
                </div>
                <section className="dash__board">
                
                    
                    <this.state.component key={1}/> 
                </section>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Dashboard);
