const Binance = require('node-binance-api');
const apiKey = "2k82QayDqeL36edeO3kkUPFTQrqGyICzPYRmw2cC3BHsSjteaIE0k2dcuGoyY4dU"
const secretKey = "tpwei5t7NOdvYJx1qhKlQ3VqZincxuo6NVg4jdGmo2kMrmhkwcvYoDrdeWp1mpJz"

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

async function marketSell() {
    console.info( await binance.futuresMarketSell( 'ONEUSDT', 100 ) );
}

async function test() {
    let position_data = await binance.futuresPositionRisk(), markets = Object.keys( position_data );
    for ( let market of markets ) {
      let obj = position_data[market], size = Number( obj.positionAmt );
      if ( size == 0 ) continue;
      console.info( `${leverage}x\t${market}\t${obj.unRealizedProfit}` );
    }
}
marketSell()

setTimeout(function(){test()}, 1000)

