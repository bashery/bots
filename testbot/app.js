const Binance = require('node-binance-api');



const binance = new Binance().options({
  APIKEY: apiKey,
  APISECRET: secretKey
});

async function futuresMarketSell() {
  let resp = await binance.futuresMarketSell( 'ONEUSDT', 100 ) 
    console.info(resp);
}
//marketSell()

async function close() {
    if ( resp.side == 'LONG' ) order = await binance.futuresMarketSell( 'ONEUSDT', 100, {reduceOnly: true} )
    else order = await binance.futuresMarketBuy( 'ONEUSDT', 100, {reduceOnly: true} )
    console.log(order)
}

// ======== SPOT API ======================

function marketSell() {
    binance.marketSell("ONEUSDT", 89.9, function(resp) {
        console.info(resp.body)
    });
}

binance.websockets.trades(['ONEUSDT'], (trades) => {
      //let {e:eventType, E:eventTime, s:symbol, p:price, q:quantity, m:maker, a:tradeId} = trades;
      console.info("spot  : ", trades.p);
});

// stream futures
function streamFutures(symbol) {
    binance.futuresMarkPriceStream( symbol, function(data) {
        // {eventType,eventTime, symbol, markPrice, indexPrice, fundingRate,fundingTime} = data;
        console.log("future: ", data.markPrice), speed = '@100ms'
    } );
}
streamFutures('oneusdt')


