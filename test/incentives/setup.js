
var BaasToken = artifacts.require("./BaasToken.sol");
var BaasInc = artifacts.require("./BaasIncentives.sol");

module.exports = {
    setup: async function(accounts)  {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];

        const baasToken = await BaasToken.deployed();
        const baasInc = await BaasInc.deployed();
        const incentivesAddress = baasInc.address;
        let setup = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);
        setup = await baasInc.setup();
    }
};