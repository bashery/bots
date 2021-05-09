const Binance = require('node-binance-api');
const helper = require('./helper.js');
const env = require('./env.js')

//console.log(env.apiKey); console.log(env.secretKey)

const binance = new Binance().options({
    APIKEY: env.hamzaKey,
  APISECRET:env.hamzaSecret
});


var resp
async function futuresMarketSell(symbol, size) {
  resp = await binance.futuresMarketSell( symbol, size ) 
  console.info(resp);
}

//futuresMarketSell('ONEUSDT', 120)

async function close(symbol, size) {
    if ( resp.side == 'LONG' ) order = await binance.futuresMarketSell( 'ONEUSDT', size, {reduceOnly: true} )
    else order = await binance.futuresMarketBuy( symbol, size, {reduceOnly: true} )
    console.log(order)
}
//setTimeout(function(){close('ONEUSDT', 120)}, 4000)

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

marketBuy('ONEUSDT', 180)
setTimeout(function(){
    marketSell('ONEUSDT',180)
}, 4000)


var tradePrice
var futurePrice
var isOpen = "" 

/* binance.websockets.trades(['ONEUSDT'], function(trades) {
    //let {e:eventType, E:eventTime, s:symbol, p:price, q:quantity, m:maker, a:tradeId} = trades;
    tradePrice =  trades.p
    compar()
//  console.info("spot  : ", trades.p);
});
*/
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
//streamFutures('ONEUSDT')

function compar() {
    //console.log(futurePrice, tradePrice)
    if ((futurePrice/1000) * 2 <= futurePrice-tradePrice && isOpen !== "opened") {
        isOpen = "opened"
        console.log("opened with : ", futurePrice, tradePrice)
        helper.execute('espeak "hello, the deal is opened with future price:"'+futurePrice+", and trade price:"+tradePrice)

    }

    if (futurePrice <= tradePrice && isOpen === "opened" ) {
        isOpen = "closed"
        console.log("close:", futurePrice, tradePrice)
        helper.execute('espeak "hello, the deal is closed with future price:"'+futurePrice+", and trade price:"+tradePrice)
    }
}


