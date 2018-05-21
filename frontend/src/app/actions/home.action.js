import { UPDATE_BTC_PRICE,UPDATE_LTC_PRICE } from '../actions/types';
export const enableBtcSocket = ()=>{
    return (dispatch) => {
        var ws = new WebSocket("wss://api.bitfinex.com/ws");
        ws.onopen = function() {
            ws.send(JSON.stringify({
                "event": "subscribe",
                "channel": "ticker",
                "pair": "BTCUSD"
            }));
        };
        ws.onmessage = function(msg) {
            var response = JSON.parse(msg.data);
            var hb = response[1];
            if (hb != "hb") {
                console.log(response[7])
                dispatch({
                    type: UPDATE_BTC_PRICE,
                    payload: response[7]
                 });
            }
        }
    };
}
export const enableLtcSocket = ()=>{
    return (dispatch) => {
        var ws = new WebSocket("wss://api.bitfinex.com/ws");
        ws.onopen = function() {
            ws.send(JSON.stringify({
                "event": "subscribe",
                "channel": "ticker",
                "pair": "LTCUSD"
            }));
        };
        ws.onmessage = function(msg) {
            var response = JSON.parse(msg.data);
            var hb = response[1];
            if (hb != "hb") {
                console.log(response[7])
                dispatch({
                    type: UPDATE_LTC_PRICE,
                    payload: response[7]
                 });
            }
        }
    };
}
export const enableDogeSocket = ()=>{
    return (dispatch) => {
        var ws = new WebSocket("wss://api.bitfinex.com/ws");
        ws.onopen = function() {
            ws.send(JSON.stringify({
                "event": "subscribe",
                "channel": "ticker",
                "pair": "DOGEUSD"
            }));
        };
        ws.onmessage = function(msg) {
            var response = JSON.parse(msg.data);
            var hb = response[1];
            if (hb != "hb") {
                console.log(response[7])
            }
        }
    };
}