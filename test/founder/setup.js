
var BaasToken = artifacts.require("./BaasToken.sol");
var BaasFounder = artifacts.require("./BaasFounder.sol");

module.exports = {
    setup: async function(accounts)  {
        const incentivesAddress = accounts[0];
        const ppAddress = accounts[1];
        const escrowAddress = accounts[2];

        const baasToken = await BaasToken.deployed();
        const baasFounder = await BaasFounder.deployed();
        const founderAddress = baasFounder.address;

        let setup = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);
        setup = await baasFounder.setup(new web3.BigNumber(2000));
    }
};