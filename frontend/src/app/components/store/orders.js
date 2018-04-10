import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
class Orders extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            amount: props.props.amount,
            currency: props.props.currency,
            price: props.props.price
        }
    }
    render() {
        return (
            <div className="pane order__pane">
                <div className="order">
                    <div>
                        <img src="http://localhost:8080/static/avatar.png" width="100" height="100"></img>
                    </div>
                    <div className="order__info">
                        <div>
                            <span className="">Amount: {this.state.amount} {this.state.currency}</span>
                        </div>
                        <div>
                            <span>Price: {this.state.price}$ /{this.state.currency}</span>
                        </div>
                    </div>
                </div>
                <div className="order__buy">
                    <button className="order__button">Buy DOGE</button>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Orders);
