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
        let tokenConversionRate = 100;

        const baasToken = await BaasToken.deployed();
        const baasROI = await BaasROI.deployed();


        let tokenAddress = await baasROI.tokenAddress();
        assert.equal(tokenAddress, baasToken.address, "baasROI should have baasToken address registered!");

        // Setup token
        let setup = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);

        await baasToken.transfer(accounts[4], web3.toBigNumber('100000'), {from: ppAddress});

        let feedContract = "100000";
        // get tokens from circulating supply - simulating market purchase
        await baasToken.transfer(baasROI.address, web3.toBigNumber(feedContract), {from: accounts[0]});
        (await baasROI.balance()).should.be.bignumber.equal(new BigNumber(feedContract));

        let expectedReturn = [];
        expectedReturn.push(await baasROI.roiOf(accounts[4], tokenConversionRate));

        optimum = await baasROI.optimalPayoutDistribution(tokenConversionRate);
        optimum[0].should.be.bignumber.equal(new BigNumber('183486'));
        optimum[1].should.be.bignumber.equal(new BigNumber('16513'));
        optimum[2].should.be.bignumber.equal(new BigNumber('1'));
        (await baasROI.currentPayoutObligation(tokenConversionRate)).should.be.bignumber.equal(new BigNumber('9000'));


        console.log("Circulating Supply: " + (await baasROI.circulatingSupply()).toNumber());
        console.log("Current Payout Obligation: " + (await baasROI.currentPayoutObligation(tokenConversionRate)).toNumber());
        console.log("Balance: " + (await baasROI.balance()).toNumber());

        optimum = await baasROI.optimalPayoutDistribution(tokenConversionRate);
        console.log("maxTokensToBeRewarded: "+ optimum[0].toNumber());
        console.log("minPayout: "+ optimum[1].toNumber());
        console.log("error: "+ optimum[2].toNumber());
        // payout
        let result = await baasROI.payoutAll(tokenConversionRate);
        let logs = result.logs;
        //
        // let paidOut = logs.filter((event) => event.event === 'PaidOut');
        // let payoutOmitted = logs.filter((event) => event.event === 'PayoutOmitted');
        // let paidAll = logs.filter((event) => event.event === 'PaidOutAll')[0];
        //
        // //Paid Out
        // paidOut.forEach((l, i) => {
        //     console.log("............." + i);
        //     testCorrectlyReceived(expectedReturn[i], accounts[5 + i], l)
        // });
        //
        //
        // // Payout Omitted
        // assert.equal(5, payoutOmitted.length, "payment omitted should have right elements");
        // assertPayoutOmittedEventContainsAddress(baasROI.address, payoutOmitted);
        // assertPayoutOmittedEventContainsAddress(escrowAddress, payoutOmitted);
        // assertPayoutOmittedEventContainsAddress(founderAddress, payoutOmitted);
        // assertPayoutOmittedEventContainsAddress(incentivesAddress, payoutOmitted);
        // assertPayoutOmittedEventContainsAddress(ppAddress, payoutOmitted);
        //
        // // // PaidOut All
        // paidAll.args.tokensProvided.should.be.bignumber.equal(9173);
        // paidAll.args.tokensPossessed.should.be.bignumber.equal(9173);
        // paidAll.args.tokenHolders.should.be.bignumber.equal(4);
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
