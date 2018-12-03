var BaasToken = artifacts.require("./BaasToken.sol");
var BaasROI = artifacts.require("./BaasROI.sol");
var utils = require('./utils.js');
//var BigNumber = require('bignumber.js');

const BigNumber = web3.BigNumber;
const should = require('chai')
    .use(require('chai-bignumber')(BigNumber))
    .should();

const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";

contract('BaasROI', function (accounts) {
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

        // 1 Euro in cent
        let tokenConversionRate = 100;



        (await baasROI.interestRate()).should.be.bignumber.equal(new BigNumber('9'));
        (await baasROI.roi(web3.toBigNumber("100"), tokenConversionRate)).should.be.bignumber.equal(new BigNumber('9'));
        (await baasROI.roi(web3.toBigNumber("500"), tokenConversionRate)).should.be.bignumber.equal(new BigNumber('45'));
        (await baasROI.roi(web3.toBigNumber("55"), tokenConversionRate)).should.be.bignumber.equal(new BigNumber('4'));


        // check rois
        let escrowRoi = await baasROI.roiOf(escrowAddress, tokenConversionRate);
        escrowRoi[0].should.be.bignumber.equal(new BigNumber('54e23'));
        escrowRoi[1].should.be.bignumber.equal(new BigNumber('6e+25'));

        let founderRoi = await baasROI.roiOf(founderAddress, tokenConversionRate);
        founderRoi[0].should.be.bignumber.equal(new BigNumber('9e23'));
        founderRoi[1].should.be.bignumber.equal(new BigNumber('1e+25'));

        let incentivesRoi = await baasROI.roiOf(incentivesAddress, tokenConversionRate);
        incentivesRoi[0].should.be.bignumber.equal(new BigNumber('9e23'));
        incentivesRoi[1].should.be.bignumber.equal(new BigNumber('1e25'));

        let ppRoi = await baasROI.roiOf(ppAddress, tokenConversionRate);
        ppRoi[0].should.be.bignumber.equal(new BigNumber('18e23'));
        ppRoi[1].should.be.bignumber.equal(new BigNumber('2e25'));

        let eligibileToken = await baasROI.eligibleToken();
        assert.equal(0, eligibileToken, "wrong token");
    });

    it("should show payout infos correctly", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];
        const incentivesAddress = accounts[3];

        const baasToken = await BaasToken.deployed();
        const baasROI = await BaasROI.deployed();


        // 1 Euro in cent
        let tokenConversionRate = 100;
        // start
        (await baasROI.eligibleToken()).should.be.bignumber.equal(new BigNumber('0'));


        await baasToken.transfer(accounts[5], web3.toBigNumber("100"), {from: escrowAddress});
        (await baasROI.eligibleToken()).should.be.bignumber.equal(new BigNumber('100'));


        (await baasROI.currentPayoutObligation(tokenConversionRate)).should.be.bignumber.equal(new BigNumber('9'));
        (await baasROI.currentPayoutObligation(3 * tokenConversionRate)).should.be.bignumber.equal(new BigNumber('3'));
        (await baasROI.circulatingSupply()).should.be.bignumber.equal(new BigNumber('100'));


        let optimum = await baasROI.optimalPayoutDistribution(tokenConversionRate);
        optimum[0].should.be.bignumber.equal(new BigNumber('91'));
        optimum[1].should.be.bignumber.equal(new BigNumber('8'));
        optimum[2].should.be.bignumber.equal(new BigNumber('1'));

        await baasToken.transfer(accounts[6], web3.toBigNumber("1000"), {from: ppAddress});
        (await baasROI.eligibleToken()).should.be.bignumber.equal(new BigNumber('1100'));
        (await baasROI.currentPayoutObligation(tokenConversionRate)).should.be.bignumber.equal(new BigNumber('99'));
        (await baasROI.currentPayoutObligation(3 * tokenConversionRate)).should.be.bignumber.equal(new BigNumber('33'));
        (await baasROI.circulatingSupply()).should.be.bignumber.equal(new BigNumber('1100'));

        optimum = await baasROI.optimalPayoutDistribution(tokenConversionRate);
        optimum[0].should.be.bignumber.equal(new BigNumber('1009'));
        optimum[1].should.be.bignumber.equal(new BigNumber('90'));
        optimum[2].should.be.bignumber.equal(new BigNumber('1'));

        await baasToken.transfer(accounts[7], web3.toBigNumber("10000"), {from: founderAddress});
        (await baasROI.eligibleToken()).should.be.bignumber.equal(new BigNumber('11100'));

        await baasToken.transfer(accounts[8], web3.toBigNumber("100000"), {from: incentivesAddress});
        (await baasROI.eligibleToken()).should.be.bignumber.equal(new BigNumber('111100'));


        (await baasROI.currentPayoutObligation(tokenConversionRate)).should.be.bignumber.equal(new BigNumber('9999'));
        (await baasROI.currentPayoutObligation(3 * tokenConversionRate)).should.be.bignumber.equal(new BigNumber('3333'));
        (await baasROI.circulatingSupply()).should.be.bignumber.equal(new BigNumber('111100'));


        optimum = await baasROI.optimalPayoutDistribution(tokenConversionRate);
        optimum[0].should.be.bignumber.equal(new BigNumber('101926'));
        optimum[1].should.be.bignumber.equal(new BigNumber('9173'));
        optimum[2].should.be.bignumber.equal(new BigNumber('1'));

        assert.equal(false, await baasROI.hasEnoughTokensForPayout(1), "not enough tokens");
    });

    it("should correctly payout", async () => {
        const escrowAddress = accounts[0];
        const ppAddress = accounts[1];
        const founderAddress = accounts[2];
        const incentivesAddress = accounts[3];
        // 1 Euro in cent
        let tokenConversionRate = 100;

        const baasToken = await BaasToken.deployed();
        const baasROI = await BaasROI.deployed();

        (await baasROI.eligibleToken()).should.be.bignumber.equal(new BigNumber('111100'));


        let feedContract = "9173";
        // get tokens from circulating supply - simulating market purchase
        await baasToken.transfer(baasROI.address, web3.toBigNumber(feedContract), {from: accounts[8]});
        (await baasROI.balance()).should.be.bignumber.equal(new BigNumber(feedContract));

        let expectedReturn = [];
        expectedReturn.push(await baasROI.roiOf(accounts[5], tokenConversionRate));
        expectedReturn.push(await baasROI.roiOf(accounts[6], tokenConversionRate));
        expectedReturn.push(await baasROI.roiOf(accounts[7], tokenConversionRate));
        expectedReturn.push(await baasROI.roiOf(accounts[8], tokenConversionRate));

        optimum = await baasROI.optimalPayoutDistribution(tokenConversionRate);
        optimum[0].should.be.bignumber.equal(new BigNumber('101926'));
        optimum[1].should.be.bignumber.equal(new BigNumber('9173'));
        optimum[2].should.be.bignumber.equal(new BigNumber('1'));
        (await baasROI.currentPayoutObligation(tokenConversionRate)).should.be.bignumber.equal(new BigNumber('9173'));

        // payout
        let result = await baasROI.payoutAll(tokenConversionRate);
        let logs = result.logs;

        let paidOut = logs.filter((event) => event.event === 'PaidOut');
        let payoutOmitted = logs.filter((event) => event.event === 'PayoutOmitted');
        let paidAll = logs.filter((event) => event.event === 'PaidOutAll')[0];

        //Paid Out
        paidOut.forEach((l, i) => {
            console.log("............." + i);
            testCorrectlyReceived(expectedReturn[i], accounts[5 + i], l)
        });


        // Payout Omitted
        assert.equal(5, payoutOmitted.length, "payment omitted should have right elements");
        assertPayoutOmittedEventContainsAddress(baasROI.address, payoutOmitted);
        assertPayoutOmittedEventContainsAddress(escrowAddress, payoutOmitted);
        assertPayoutOmittedEventContainsAddress(founderAddress, payoutOmitted);
        assertPayoutOmittedEventContainsAddress(incentivesAddress, payoutOmitted);
        assertPayoutOmittedEventContainsAddress(ppAddress, payoutOmitted);

        // // PaidOut All
        paidAll.args.tokensProvided.should.be.bignumber.equal(9173);
        paidAll.args.tokensPossessed.should.be.bignumber.equal(9173);
        paidAll.args.tokenHolders.should.be.bignumber.equal(4);
    });
});

const testCorrectlyReceived = (expected, account, log) => {
    expected[0].should.be.bignumber.equal(log.args.tokensReceived);
    expected[1].should.be.bignumber.equal(log.args.tokensPossessed);
    assert.equal(account, log.args.receiver, "account");
};


const assertPayoutOmittedEventContainsAddress = (address, payoutOmitted) => {
    assert.equal(1, payoutOmitted.filter((event) => event.args.omittedReceiver === address).length, "payout omitted doesnt contain: " + address);
};
