const Binance = require('node-binance-api');

const binance = new Binance().options({
  APIKEY: apiKey,
  APISECRET: secretKey
});

async function prcs() {
    console.info( await binance.futuresPrices() );
}
//prcs()

async function balance() {
    console.info( await binance.futuresBalance() );
}
//balance()
let resp
async function marketSell() {
  resp = await binance.futuresMarketSell( 'ONEUSDT', 100 ) 
    console.info(resp);
}
//marketSell()

async function close() {
    if ( resp.side == 'LONG' ) order = await binance.futuresMarketSell( 'ONEUSDT', 100, {reduceOnly: true} )
    else order = await binance.futuresMarketBuy( 'ONEUSDT', 100, {reduceOnly: true} )
    console.log(order)
}
//setTimeout(function(){close()}, 5000)

// =========================================
// ======== SPOT API ======================
async function spotPrice() {
    console.log("start spot")
    let ticker = await binance.prices();
    console.info(`spot Price : ${ticker.ONEUSDT}`);
}
spotPrice()
