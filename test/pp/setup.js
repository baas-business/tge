
var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");

module.exports = {
    setup: async function(accounts)  {
        const incentivesAddress = accounts[0];
        const founderAddress = accounts[1];
        const escrowAddress = accounts[2];

        const baasToken = await BaasToken.deployed();
        const baasPP = await BaasPP.deployed();
        const ppAddress = baasPP.address;

        let setup = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);
    }
};