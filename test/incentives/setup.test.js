var BaasToken = artifacts.require("./BaasToken.sol");
var BaasInc = artifacts.require("./BaasIncentives.sol");

const BigNumber = web3.BigNumber;
const should = require('chai')
    .use(require('chai-bignumber')(BigNumber))
    .should();

const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";

contract('BaasIncentives', function (accounts) {
    it("should initialize correctly", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];

        const baasToken = await BaasToken.deployed();
        const baasInc = await BaasInc.deployed();
        const incentivesAddress = baasInc.address;


        let tokenAddress = await baasInc.tokenAddress();
        assert.equal(tokenAddress, baasToken.address, "baasROI should have baasToken address registered!");

        // Setup token
        let setup = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);
        setup = await baasInc.setup();


        (await baasInc.incentivesProvided()).should.be.bignumber.equal(new BigNumber("0"));
        (await baasInc.incentivesLeft()).should.be.bignumber.equal(new BigNumber("10e24"));

    });

});
