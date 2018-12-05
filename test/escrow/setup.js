
var BaasToken = artifacts.require("./BaasToken.sol");
var BaasEscrow = artifacts.require("./BaasEscrow.sol");

module.exports = {
    setup: async function(accounts)  {
        const incentivesAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];

        const baasToken = await BaasToken.deployed();
        const baasEscrow = await BaasEscrow.deployed();
        const escrowAddress = baasEscrow.address;
        let setup = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);
        setup = await baasEscrow.setup(new web3.BigNumber(2000));
    }
};