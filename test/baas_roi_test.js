var BaasToken = artifacts.require("./BaasToken.sol");
var BaasROI = artifacts.require("./BaasROI.sol");
var utils = require('./utils.js');
var BigNumber = require('bignumber.js');
var BN = require('bn.js');

const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";

contract('BaasPP', function (accounts) {
    it("should initialize correctly", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];
        const incentivesAddress = accounts[3];

        const baasToken = await BaasToken.deployed();
        const baasROI = await BaasROI.deployed();


        let tokenAddress = await baasROI.tokenAddress();
        assert.equal(tokenAddress, baasToken.address, "baasROI should have baasToken address registered!");

        // Setup token
        let setup = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);
        setup = await baasROI.setup();

        let pots = await baasROI.pots();
        pots.forEach((a, i) => {
            assert.equal(accounts[i], a, "account should be equal")
        });

        assert.equal(9, await baasROI.interestRate(), "wrong interest");
        utils.compareBigNumber(new BigNumber('9'), await baasROI.roi(web3.toBigNumber("100"), 1), "wrong roi");
        utils.compareBigNumber(new BigNumber('45'), await baasROI.roi(web3.toBigNumber("500"), 1), "wrong roi");
        utils.compareBigNumber(new BigNumber('4'), await baasROI.roi(web3.toBigNumber("55"), 1), "wrong roi");

        // check rois
        let escrowRoi  = await baasROI.roiOf(escrowAddress, 1)
        utils.compareBigNumber(new BigNumber('54e23'), escrowRoi[0], "wrong roiOf");
        utils.compareBigNumber(new BigNumber('6e+25'), escrowRoi[1], "wrong roiOf");

        let founderRoi  = await baasROI.roiOf(founderAddress, 1)
        utils.compareBigNumber(new BigNumber('9e23'), founderRoi[0], "wrong roiOf");
        utils.compareBigNumber(new BigNumber('1e+25'), founderRoi[1], "wrong roiOf");


        let incentivesRoi  = await baasROI.roiOf(incentivesAddress, 1)
        utils.compareBigNumber(new BigNumber('9e23'), incentivesRoi[0], "wrong roiOf");
        utils.compareBigNumber(new BigNumber('1e+25'), incentivesRoi[1], "wrong roiOf");

        let ppRoi  = await baasROI.roiOf(ppAddress, 1)
        utils.compareBigNumber(new BigNumber('18e23'), ppRoi[0], "wrong roiOf");
        utils.compareBigNumber(new BigNumber('2e+25'), ppRoi[1], "wrong roiOf");



        let eligibileToken = await baasROI.eligibleToken();
        assert.equal(0, eligibileToken, "wrong token");


    });

    it("should show eligibile tokens correctly", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];
        const incentivesAddress = accounts[3];

        const baasToken = await BaasToken.deployed();
        const baasROI = await BaasROI.deployed();


        // start
        utils.compareBigNumber(new BigNumber('0'), await baasROI.eligibleToken(), "wrong eligible token");

        await baasToken.transfer(accounts[5], web3.toBigNumber("124124312311233"), {from: escrowAddress});
        utils.compareBigNumber(new BigNumber('124124312311233'), await baasROI.eligibleToken(), "wrong eligible token");


        await baasToken.transfer(accounts[6], web3.toBigNumber("10000000000000000"), {from: ppAddress});
        utils.compareBigNumber(new BigNumber('10124124312311233'), await baasROI.eligibleToken(), "wrong eligible token");


        await baasToken.transfer(accounts[7], web3.toBigNumber("100000000000000000000000"), {from: founderAddress});
        utils.compareBigNumber(new BigNumber('1.00000010124124312311233e+23'), await baasROI.eligibleToken(), "wrong eligible token after founder transfer");


        await baasToken.transfer(accounts[8], web3.toBigNumber("1000000000000000000000000"), {from: incentivesAddress});
        utils.compareBigNumber(new BigNumber('1.100000010124124312311233e+24'), await baasROI.eligibleToken(), "wrong eligible token after incentives transfer");

        utils.compareBigNumber(new BigNumber('9.900000091117118810801e+22'), await baasROI.currentPayout(1), "wrong payout");
        utils.compareBigNumber(new BigNumber('3.3000000303723729369336e+22'), await baasROI.currentPayout(3), "wrong payout");

    });

    it("should correctly payout", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];
        const incentivesAddress = accounts[3];

        const baasToken = await BaasToken.deployed();
        const baasROI = await BaasROI.deployed();

        utils.compareBigNumber(new BigNumber('1.100000010124124312311233e+24'), await baasROI.eligibleToken(), "wrong eligible token");

        assert.equal(false, await baasROI.hasEnoughTokensForPayout(1), "not enough tokens");
        utils.compareBigNumber(new BigNumber('9.900000091117118810801e+22'),  await baasROI.tokensNeededForPayout(1), "tokens needed wrong");

        await baasToken.transfer(baasROI.address, web3.toBigNumber("10e22"), {from: accounts[8]});
       // utils.compareBigNumber(new BigNumber('1000000e18'), await baasROI.eligibleToken(), "wrong eligible token");
        // start
//        utils.compareBigNumber(new BigNumber('1.100000010124124312311233e+24'), await baasROI.eligibleToken(), "wrong eligible token");


       // let cp = await baasROI.currentPayout(1);
       // console.log(cp.toNumber());
       // let receipt = await baasROI.payoutAll(1);
       // console.log(receipt);

    });
});