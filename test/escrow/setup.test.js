var BaasToken = artifacts.require("./BaasToken.sol");
var BaasEscrow = artifacts.require("./BaasEscrow.sol");
var setup = require("./setup").setup;

const BigNumber = web3.BigNumber;
const should = require('chai')
    .use(require('chai-bignumber')(BigNumber))
    .should();

const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";

contract('BaasEscrow', function (accounts) {
    it("should initialize correctly", async () => {
        const baasToken = await BaasToken.deployed();
        const baasEscrow = await BaasEscrow.deployed();

        await setup(accounts);


        let tokenAddress = await baasEscrow.tokenAddress();
        assert.equal(tokenAddress, baasToken.address, "baasROI should have baasToken address registered!");




    });

});
