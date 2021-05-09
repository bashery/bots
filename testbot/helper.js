
const { exec } = require("child_process");



async function isoCrose(symbol, isoCro) {
    console.info( await binance.futuresMarginType( symbol, isoCro ) ); // OR 'CROSSED'
} 
//isoCrose('ONEUSDT', 'CROSSED')

// clear screan
function clear() {
    setTimeout( function() {console.log('\033c');clear()}, 1000)
}

function execute(cmd) {
exec(cmd, (error, stdout, stderr) => {
        if (error) {
            console.log(`error: ${error.message}`);
            return;
        }
        if (stderr) {
            console.log(`stderr: ${stderr}`);
            return;
        }
        if(stdout) {
            console.log(`stdout: ${stdout}`);
            return
            
        }
    });
} 

module.exports = {execute, isoCrose}
//execute('espeak "hello, the work is done"')

