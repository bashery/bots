const Binance = require('node-binance-api');
const env = require('./env.js')

console.log(env.apiKey)
console.log(env.secretKey)

const binance = new Binance().options({
  APIKEY: env.apiKey,
  APISECRET: env.secretKey
});

async function futuresMarketSell(symbol) {
  let resp = await binance.futuresMarketSell( symbol, 100 ) 
    console.info(resp);
}
//marketSell()

async function close(symbol) {
    if ( resp.side == 'LONG' ) order = await binance.futuresMarketSell( 'ONEUSDT', 100, {reduceOnly: true} )
    else order = await binance.futuresMarketBuy( symbol, 100, {reduceOnly: true} )
    console.log(order)
}

// ======== SPOT API ======================

function marketSell(symbol, size) {
    binance.marketSell(symbol, size, function(resp) {
        console.info(resp.body)
    });
}

function marketBuy(symbol, size) {
    binance.marketSell(symbol, size, function(resp) {
        console.info(resp.body)
    });
}




var tradePrice
var futurePrice
var isOpen = "" 

binance.websockets.trades(['ONEUSDT'], function(trades) {
    //let {e:eventType, E:eventTime, s:symbol, p:price, q:quantity, m:maker, a:tradeId} = trades;
    tradePrice =  trades.p
    compar()
//  console.info("spot  : ", trades.p);
});

//streamTrade()
// stream futures

function streamFutures(symbol) {
    binance.futuresMarkPriceStream(symbol, function(data) {
        // {eventType,eventTime, symbol, markPrice, indexPrice, fundingRate,fundingTime} = data;
        futurePrice = data.markPrice
        speed = '@100ms'
        compar()
        //console.log(data)
   }); 
}
streamFutures('ONEUSDT')

function compar() {
    //console.log(futurePrice, tradePrice)
    if ((futurePrice/1000) * 3 <= futurePrice-tradePrice && isOpen !== "opened") {
        isOpen = "opened"
        console.log("opened with : ", futurePrice, tradePrice)
    }

    if (futurePrice <= tradePrice && isOpen === "opened" ) {
        isOpen = "closed"
        console.log("close:", futurePrice, tradePrice)
    }
}


