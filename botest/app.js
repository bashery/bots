//import Binance from 'binance-api-node'


const APIKEY = '2k82QayDqeL36edeO3kkUPFTQrqGyICzPYRmw2cC3BHsSjteaIE0k2dcuGoyY4dU'
const SECRITKEY = 'tpwei5t7NOdvYJx1qhKlQ3VqZincxuo6NVg4jdGmo2kMrmhkwcvYoDrdeWp1mpJz'


const Binance = require('binance-api-node').default

const client = Binance({
    apiKey:APIKEY,
    apiSecret:SECRITKEY
})


async function tradeFee() {
    console.log(await client.tradeFee())
}
//tradeFee()


// extchange info
async function getExchangeInfo()  {
    let info = await client.exchangeInfo()
    console.log(info)
}

//getExchangeInfo()

async function book()  {
    let booked = await client.book({ symbol: 'ONEUSDT' })
    console.log(booked)
}
//book()


async function futuresPing() {
    let pong = await client.futuresPing()
    console.log(pong)
}
//futuresPing()

async function futuresTime() { 
    let pong = await client.futuresTime()
    console.log(pong)
}

//futuresTime()

async function futuresAccontBalance() { 
    let pong = await client.futuresAccountBalance()
    console.log(pong)
}
//futuresAccontBalance()
async function tradesHistory() {
    let result = await client.tradesHistory({ symbol: 'BTCUSDT'})


    console.log(result)
    console.log("tih")
    
}
// tradesHistory()


async function myTrades() {
    console.log(
         await client.myTrades({
             symbol: 'ONEUSDT',
        }),
    )
}
//myTrades()

async function depositHist() {
    console.log(await client.depositHistory())
}
// depositHist()



async function openOrders() {
    console.log(
      await client.openOrders({
          symbol: 'ONEUSDT',
      }),
    )
}
//openOrders()


async function futuresTrades() {
    console.log(await client.futuresTrades({ symbol: 'ONEUSDT' }))
}
futuresTrades()

async function futuresAllBookTickers() {
    console.log(await client.futuresAllBookTickers())
}
//futuresAllBookTickers()

async function futuresAllBookTickers() {
    console.log(await client.futuresMarkPrice())
}
//futuresAllBookTickers()
//

async function futuresAllForceOrders() {
    console.log(await client.futuresAllForceOrders())
}
//futuresAllForceOrders()
//console.log(client.getInfo())
//

//client.ws.futuresTicker('ONEUSDT', ticker => {
  //console.log(ticker)
//})

async function openOrders() {
console.log(
  await client.openOrders({
      symbol: 'ONEUSDT',
  }),
)
}
//openOrders()
// changeMargin()

async function order() {
   console.log(
  await client.order({
    symbol: 'ONEUSDT',
    side: 'SELL',
    quantity: '12',
  
  }),
) 
}
//order()

async function marketPrice() {
    console.log(await client.futuresMarkPrice('ONEUSDT'))
}

//marketPrice()
