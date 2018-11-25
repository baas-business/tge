var BaasToken = artifacts.require("./BaasToken.sol");
var utils = require('./utils.js');
var BigNumber = require('bignumber.js');

const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";

contract('BaasToken', function (accounts) {
    it("should initialize with 0 total supply and isInitialized false", async () => {

        const baasToken = await BaasToken.deployed();

        let totalSupply = await baasToken.totalSupply();
        assert.equal(totalSupply, 0, "totalSupply should be 0");
        let isInitialized = await baasToken.isInitialized();
        assert.equal(isInitialized, false, "isInitialized should be false");
        let owner = await baasToken.owner();
        assert.equal(owner, contractDeployer, "owner should be contract deployer");
        let transferResult = await baasToken.transfer(accounts[0], 200);
        assert.equal(transferResult.logs.length, 0, "transfer should not create logs as it will return")
    });

    it("should setup correctly", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];
        const incentivesAddress = accounts[3];

        const baasToken = await BaasToken.deployed();

        let isInitialized = await baasToken.isInitialized();
        assert.equal(isInitialized, false, "isInitialized should be false");

        let init = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);

        // check flag
        isInitialized = await baasToken.isInitialized();
        assert.equal(isInitialized, true, "isInitialized should be true after setup");

        // check total supply
        let totalSupply = await baasToken.totalSupply();
        utils.compareBigNumber(new BigNumber('100e18'), totalSupply, "Total Supply");

        // check pots
        let pots = await baasToken.pots();
        pots.forEach((a,i) => {
            assert.equal(accounts[i], a, "account should be equal")
        });

        // check balances
        let escrowSupply = await baasToken.balanceOf(escrowAddress);
        utils.compareBigNumber(new BigNumber('60e18'), escrowSupply, "Escrow Supply");
        let ppSupply = await baasToken.balanceOf(ppAddress);
        utils.compareBigNumber(new BigNumber('20e18'), ppSupply, "PP Supply");
        let founderSupply = await baasToken.balanceOf(founderAddress);
        utils.compareBigNumber(new BigNumber('10e18'), founderSupply, "Founder Supply")
        let incentivesSupply = await baasToken.balanceOf(incentivesAddress);
        utils.compareBigNumber(new BigNumber('10e18'), incentivesSupply, "Incentives Supply")

    });
});
