import React, { PureComponent } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../actions';
class Orders extends PureComponent {
    constructor(props){
        super(props)
        this.state = {
            type: props.type
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
                            <span className="">Amount: 25 DOGE</span>
                        </div>
                        <div>
                            <span>Price: 0.015$ </span>
                        </div>
                    </div>
                </div>
                <div>
                    <button className="form__button">Buy DOGE</button>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return { features: state.features.homePageFeatures }
}

export default connect(mapStateToProps, actions)(Orders);
