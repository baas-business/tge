var fs = require('fs');


extract_abi_and_bin("BaasToken");
extract_abi_and_bin("BaasEscrow");
extract_abi_and_bin("BaasFounder");
extract_abi_and_bin("BaasIncentive");
extract_abi_and_bin("BaasPP");
extract_abi_and_bin("BaasROI");

function extract_abi_and_bin(contractName) {
    let content = JSON.parse(fs.readFileSync('../build/contracts/'+contractName+'.json'));


    fs.writeFileSync(contractName + '_abi.json', JSON.stringify(content.abi), 'utf8');
    fs.writeFileSync(contractName + '_bin', content.bytecode);
}