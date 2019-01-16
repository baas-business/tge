var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var BaasROI = artifacts.require("./BaasROI.sol");
var assertions = require('./assertions');
var BigNumber = require('bignumber.js');
var BN = web3.utils.BN;
var utils = require('../utils');

let baasToken;
let baasPP;
let baasROI;

let tokenHolder1;

contract('BaasROI', function (accounts) {

    beforeEach(async function () {
        baasToken = await BaasToken.deployed();
        baasPP = await BaasPP.deployed();
        baasROI = await BaasROI.deployed();
        tokenHolder1 = accounts[9];
    });

    describe('Setup', function () {
        it("should have correct constant values", async () => {
            utils.compareBigNumber(new BigNumber("9"), await baasROI.interestRate(), "interest rate");
        });

        it("should initialize correctly", async () => {

            assert.equal(await baasROI.tokenAddress(), baasToken.address, "tokenAddress is wrong");

            await assertions.assert(baasROI, {
                balance: new BigNumber("0"),
                circulatingSupply: new BigNumber("0"),
                eligibleToken: new BigNumber("0"),
                tokenEuroConversionRate: [100],
                contractHasEnoughTokensForPayout: [true],
                currentPayoutObligation: [new BigNumber("0")],
                optimalPayoutDistribution: [
                    {
                        contractCanPayMax: new BigNumber("0"),
                        contractShouldHaveMin: new BigNumber("0"),
                        error: new BigNumber("0"),
                    }
                ]
            });
        });

        it("should deliver correct roi for token balance", async () => {
            utils.compareBigNumber(new BigNumber("89"), await baasROI.roi(new BigNumber("999"), 100), "roi");
            utils.compareBigNumber(new BigNumber("115"), await baasROI.roi(new BigNumber("999"), 78), "roi");
            utils.compareBigNumber(new BigNumber("9300"), await baasROI.roi(new BigNumber("12400"), 12), "roi");
            utils.compareBigNumber(new BigNumber("5849"), await baasROI.roi(new BigNumber("81244"), 125), "roi");
            utils.compareBigNumber(new BigNumber("63514"), await baasROI.roi(new BigNumber("917432"), 130), "roi");
        });
    });

    describe('After Eligible Token Holder Received Tokens', function () {

        it("should have correct values when token holder existent", async () => {
            // prepare
            await baasToken.setup(accounts[0], baasPP.address, accounts[1], accounts[2]);
            // send some tokens from private placement to an address
            await baasPP.issue(tokenHolder1, new BN("1 000 000"), 1, {from: accounts[0]});

            let expected = {
                balance: new BigNumber("0"),
                circulatingSupply: new BigNumber("1000000"),
                eligibleToken: new BigNumber("1000000"),
                tokenEuroConversionRate: [50, 100, 130],
                contractHasEnoughTokensForPayout: [false, false, false],
                currentPayoutObligation: [new BigNumber("180000"), new BigNumber("90000"), new BigNumber("69230")],
                optimalPayoutDistribution: [
                    {
                        contractCanPayMax: new BigNumber("847457"),
                        contractShouldHaveMin: new BigNumber("152542"),
                        error: new BigNumber("1"),
                    },
                    {
                        contractCanPayMax: new BigNumber("917431"),
                        contractShouldHaveMin: new BigNumber("82568"),
                        error: new BigNumber("1"),
                    },
                    {
                        contractCanPayMax: new BigNumber("935251"),
                        contractShouldHaveMin: new BigNumber("64748"),
                        error: new BigNumber("1"),
                    }
                ]
            };

            await assertions.assert(baasROI, expected);
        });

        it("should have correct values after contract received enough for euro token conversion rate 130", async () => {

            // send some tokens from private placement to contract
            await baasToken.transfer(baasROI.address, new BN("64748"), {from: tokenHolder1});

            let expected = {
                tokenAddress: baasToken.address,
                balance: new BigNumber("64748"),
                circulatingSupply: new BigNumber("1000000"),
                eligibleToken: new BigNumber("935252"),
                tokenEuroConversionRate: [50, 100, 130],
                contractHasEnoughTokensForPayout: [false, false, true],
                currentPayoutObligation: [new BigNumber("168345"), new BigNumber("84172"), new BigNumber("64748")],
                optimalPayoutDistribution: [
                    {
                        contractCanPayMax: new BigNumber("847457"),
                        contractShouldHaveMin: new BigNumber("152542"),
                        error: new BigNumber("1"),
                    },
                    {
                        contractCanPayMax: new BigNumber("917431"),
                        contractShouldHaveMin: new BigNumber("82568"),
                        error: new BigNumber("1"),
                    },
                    {
                        contractCanPayMax: new BigNumber("935251"),
                        contractShouldHaveMin: new BigNumber("64748"),
                        error: new BigNumber("1"),
                    }
                ]
            };

            await assertions.assert(baasROI, expected);
        });

        it("should have correct values after contract received enough for euro token conversion rate 100", async () => {

            // send some tokens from private placement to contract
            await baasToken.transfer(baasROI.address, new BN("17820"), {from: tokenHolder1});

            let expected = {
                tokenAddress: baasToken.address,
                balance: new BigNumber("82568"),
                circulatingSupply: new BigNumber("1000000"),
                eligibleToken: new BigNumber("917432"),
                tokenEuroConversionRate: [50, 100, 130],
                contractHasEnoughTokensForPayout: [false, true, true],
                currentPayoutObligation: [new BigNumber("165137"), new BigNumber("82568"), new BigNumber("63514")],
                optimalPayoutDistribution: [
                    {
                        contractCanPayMax: new BigNumber("847457"),
                        contractShouldHaveMin: new BigNumber("152542"),
                        error: new BigNumber("1"),
                    },
                    {
                        contractCanPayMax: new BigNumber("917431"),
                        contractShouldHaveMin: new BigNumber("82568"),
                        error: new BigNumber("1"),
                    },
                    {
                        contractCanPayMax: new BigNumber("935251"),
                        contractShouldHaveMin: new BigNumber("64748"),
                        error: new BigNumber("1"),
                    }
                ]
            };

            await assertions.assert(baasROI, expected);
        });

    });

    describe('Payout All', function () {
        it("should fail with wrong owner", async () => {
            let shouldFail = await baasROI.payoutAll(130, {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should fail with converison rate 50", async () => {
            let shouldFail = await baasROI.payoutAll(50, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "not enough tokens for payout available")
        });

        it("should token holders correctly at token conversion rate 130", async () => {

            // send some tokens from private placement to contract
            let result = await baasROI.payoutAll(130, {from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(7, logs.length, "should emit 7 logs");

            utils.compareEvent(logs[0], {
                event: "PayoutOmitted",
                arg: [accounts[0]],
            });

            utils.compareEvent(logs[1], {
                event: "PayoutOmitted",
                arg: [baasPP.address],
            });

            utils.compareEvent(logs[2], {
                event: "PayoutOmitted",
                arg: [accounts[1]],
            });

            utils.compareEvent(logs[3], {
                event: "PayoutOmitted",
                arg: [accounts[2]],
            });

            utils.compareEvent(logs[4], {
                event: "PaidOut",
                arg: [
                    tokenHolder1,
                    new BN("63514"),
                    new BN("917432")
                ],
            });

            utils.compareEvent(logs[5], {
                event: "PayoutOmitted",
                arg: [baasROI.address],
            });


            utils.compareEvent(logs[6], {
                event: "PaidOutAll",
                arg: [
                    new BN("63514"),
                    new BN("82568"),
                    new BN("1")
                ]
            });

            let expected = {
                balance: new BigNumber("19054"),
                circulatingSupply: new BigNumber("1000000"),
                eligibleToken: new BigNumber("980946"),
                tokenEuroConversionRate: [50, 100, 130],
                contractHasEnoughTokensForPayout: [false, false, false],
                currentPayoutObligation: [new BigNumber("176570"), new BigNumber("88285"), new BigNumber("67911")],
                optimalPayoutDistribution: [
                    {
                        contractCanPayMax: new BigNumber("847457"),
                        contractShouldHaveMin: new BigNumber("152542"),
                        error: new BigNumber("1"),
                    },
                    {
                        contractCanPayMax: new BigNumber("917431"),
                        contractShouldHaveMin: new BigNumber("82568"),
                        error: new BigNumber("1"),
                    },
                    {
                        contractCanPayMax: new BigNumber("935251"),
                        contractShouldHaveMin: new BigNumber("64748"),
                        error: new BigNumber("1"),
                    }
                ]
            };

            await assertions.assert(baasROI, expected);
        });
    });

    describe('Withdraw From Contract', function () {
        it("should fail with wrong owner", async () => {
            let shouldFail = await baasROI.withdrawFromContract(accounts[6], {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should receive balance from contract", async () => {
            let result = await baasROI.withdrawFromContract(accounts[6], {from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Withdrawn",
                arg: [
                    accounts[6],
                    new BigNumber("19054")
                ]
            });

            utils.compareBigNumber(new BigNumber("19054"), await baasToken.balanceOf(accounts[6]), "token balance");
        });
    });

});