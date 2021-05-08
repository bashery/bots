
const { exec } = require("child_process");

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
    console.log(`stdout: ${stdout}`);
});
} 

execute('espeak "hello, the work is done"')
