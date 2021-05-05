const WebSocket = require('ws');

const ws = new WebSocket('wss://stream.binance.com:9443/ws/oneusdt@trade');

ws.on('message', function incoming(json) {
    let data = JSON.parse(json);
    console.log(data.p);
});
