import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
import CurrencyInfoComponent from './components/currencyInfoComponent';
import ChartComponent from './components/chartComponent';
import TransactionComponent from './components/transactionComponent';
class OverviewContainer extends PureComponent {
    constructor(props){
        super(props);
        this.state = {
            btcBalance:null,
            ltcBalance:null,
            dogeBalance:null
          };
    }
    componentWillMount(){
        this.props.getUserBasic();
        this.props.getUserTransactions();
    }
    renderCurrency() {
        if (this.props.overview) {
            return (
                <CurrencyInfoComponent props={true}/>,
                <CurrencyInfoComponent props={true}/>
            );
        }
    }
    componentWillReceiveProps(nextProps){
        if((nextProps.DogeBalance !=undefined) && (nextProps.BtcBalance !=undefined) && (nextProps.LtcBalance !=undefined)){
            this.setState({
                dogeBalance:nextProps.DogeBalance,
                btcBalance:nextProps.BtcBalance,
                ltcBalance:nextProps.LtcBalance
            })
        }
    }
    render() {
        if(this.state.dogeBalance != null){
            var dogeCoin = <CurrencyInfoComponent key={0} props={{currency:'DOGE', amount:this.state.dogeBalance, prices:1, transactions:0, transaction:0}}/>
        }else{
            var dogeCoin = 'Loading'
        }
        if(this.state.btcBalance != null){
            var bitCoin = <CurrencyInfoComponent key={1} props={{currency:'BTC', amount:this.state.btcBalance, prices:2, transactions:0,transaction:0}}/>
        }else{
            var bitCoin = 'Loading'
        }
        if(this.state.ltcBalance != null){
            var liteCoin = <CurrencyInfoComponent key={2} props={{currency:'LTC', amount:this.state.ltcBalance, prices:2, transactions:0,transaction:0}}/>
        }else{
            var liteCoin = 'Loading'
        }
        return (
            <div className="dash overview">
                <div className="overview__detail">
                    {bitCoin}
                    {dogeCoin}
                    {liteCoin}
                </div>
                <div>
                    <ChartComponent />
                </div>
                <div>
                    <TransactionComponent />
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => ({...state.overview})

export default connect(mapStateToProps, actions)(OverviewContainer);
