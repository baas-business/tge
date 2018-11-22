var fs = require('fs');


var content = JSON.parse(fs.readFileSync('../build/contracts/UserRegistry.json'));


fs.writeFileSync('abi.json', JSON.stringify(content.abi), 'utf8');
fs.writeFileSync('bin', content.bytecode);
