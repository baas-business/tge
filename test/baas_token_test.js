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
        let transferResult = await baasToken.transfer(accounts[0], 200).catch(e => false);
        assert.equal(false, transferResult, "transfer should fail")
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
        utils.compareBigNumber(new BigNumber('100e24'), totalSupply, "Total Supply");

        // check pots
        let pots = await baasToken.pots();
        pots.forEach((a, i) => {
            assert.equal(accounts[i], a, "account should be equal")
        });

        // check balances
        let escrowSupply = await baasToken.balanceOf(escrowAddress);
        utils.compareBigNumber(new BigNumber('60e24'), escrowSupply, "Escrow Supply");
        let ppSupply = await baasToken.balanceOf(ppAddress);
        utils.compareBigNumber(new BigNumber('20e24'), ppSupply, "PP Supply");
        let founderSupply = await baasToken.balanceOf(founderAddress);
        utils.compareBigNumber(new BigNumber('10e24'), founderSupply, "Founder Supply");
        let incentivesSupply = await baasToken.balanceOf(incentivesAddress);
        utils.compareBigNumber(new BigNumber('10e24'), incentivesSupply, "Incentives Supply");

        let tokenHolders = await baasToken.tokenHolderSnapShot();
        assert.equal(escrowAddress, tokenHolders[0], "escrow address should be token holders");
        assert.equal(incentivesAddress, tokenHolders[3], "incentives address should be token holders");
        assert.equal(ppAddress, tokenHolders[1], "pp address should be token holders");
        assert.equal(founderAddress, tokenHolders[2], "founder address should be token holders");

        let tokenHolderCount = await baasToken.tokenHolderCount();
        assert.equal(4, tokenHolderCount, "token holders count wrong");

        assert.equal(true, await baasToken.isTokenHolder(escrowAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(incentivesAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(founderAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(ppAddress), "should be token holders");
        assert.equal(false, await baasToken.isTokenHolder(accounts[9]), "should not be token holders");
    });

    it("should deal with token holders", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];
        const incentivesAddress = accounts[3];
        const receiver = accounts[4];
        const receiver2 = accounts[5];

        const baasToken = await BaasToken.deployed();

        let amount = web3.toBigNumber("60e24");
        let receipt = await baasToken.transfer(receiver, amount, {from: escrowAddress});

        let tokenHolders = await baasToken.tokenHolderSnapShot();

        // address moved
        assert.equal(incentivesAddress, tokenHolders[0], "incentives address should be token holders");
        assert.equal(ppAddress, tokenHolders[1], "pp address should be token holders");
        assert.equal(founderAddress, tokenHolders[2], "founder address should be token holders");
        assert.equal(receiver, tokenHolders[3], "receiver address should be token holders");

        assert.equal(4, await baasToken.tokenHolderCount(), "token holders count wrong");

        assert.equal(false, await baasToken.isTokenHolder(escrowAddress), "should not be token holders");
        assert.equal(true, await baasToken.isTokenHolder(incentivesAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(founderAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(ppAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(receiver), "should be token holders");
        assert.equal(false, await baasToken.isTokenHolder(receiver2), "should not be token holders");

        /*
            Smaller amount
         */

        receipt = await baasToken.transfer(receiver2, web3.toBigNumber("10e24"), {from: receiver});
        assert.equal(5, await baasToken.tokenHolderCount(), "token holders count wrong");
        assert.equal(false, await baasToken.isTokenHolder(escrowAddress), "should not be token holders");
        assert.equal(true, await baasToken.isTokenHolder(incentivesAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(founderAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(ppAddress), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(receiver), "should be token holders");
        assert.equal(true, await baasToken.isTokenHolder(receiver2), "should be token holders");

    });


    it("should pause correctly", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];
        const incentivesAddress = accounts[3];
        const receiver = accounts[4];
        const receiver2 = accounts[5];

        const baasToken = await BaasToken.deployed();


        assert.equal(false, await baasToken.paused(), "should not be paused");

        // pause
        let pauseReceipt = await baasToken.setPause(true);
        assert.equal("0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2", pauseReceipt.receipt.logs[0].topics[0], "should emit correct log");
        assert.equal("0x0000000000000000000000000000000000000000000000000000000000000001", pauseReceipt.receipt.logs[0].data, "should emit correct log");

        assert.equal(true, await baasToken.paused(), "should be paused");

        let receipt = await baasToken.transfer(receiver, web3.toBigNumber("1e24"), {from: ppAddress}).catch(e => false);
        assert.equal(false, receipt, "transfer should not be possible");


        receipt = await baasToken.approve(receiver, web3.toBigNumber("1e24"), {from: ppAddress}).catch(e => false);
        //receipt = await baasToken.transferFrom(receiver, ppAddress, web3.toBigNumber("1e24"), {from: ppAddress}).catch(e => false);
        //assert.equal(false, receipt, "transferfrom should fail");

        // unpause
        pauseReceipt = await baasToken.setPause(false);
        assert.equal("0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2", pauseReceipt.receipt.logs[0].topics[0], "should emit correct log");
        assert.equal("0x0000000000000000000000000000000000000000000000000000000000000000", pauseReceipt.receipt.logs[0].data, "should emit correct log");

        assert.equal(false, await baasToken.paused(), "should not be paused");

        receipt = await baasToken.transfer(receiver, web3.toBigNumber("1e10"), {from: ppAddress}).then(r => true).catch(e => false);
        assert.equal(true, receipt, "transfer should be possible");


        let spender = ppAddress;
        receipt = await baasToken.approve(receiver, web3.toBigNumber("1e10"), {from: spender}).then(r => true).catch(e => false);
        assert.equal(true, receipt, "approve should not fail");
        assert.equal(web3.toBigNumber("1e10").toNumber(),(await  baasToken.allowance(spender, receiver)).toNumber(), "should be approved correct amount");

        receipt = await baasToken.transferFrom(spender, receiver, web3.toBigNumber("1e10"), {from: receiver}).then(r => true).catch(e => {return false;});
        assert.equal(true, receipt, "transferFrom should not fail");
    });
});
