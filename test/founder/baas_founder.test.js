var BaasToken = artifacts.require("./BaasToken.sol");
var BaasFounder = artifacts.require("./BaasFounder.sol");
var assertions = require('./assertions');
var BigNumber = require('bignumber.js');
var BN = web3.utils.BN;
var utils = require('../utils');

let baasToken;
let baasFounder;

contract('BaasFounder', function (accounts) {

    beforeEach(async function () {
        baasToken = await BaasToken.deployed();
        baasFounder = await BaasFounder.deployed();
    });

    describe('Setup', function () {
        it("should have correct constant values", async () => {
            utils.compareBigNumber(new BigNumber("8e24"), await baasFounder.founder1Supply(), "founder 1 supply");
            utils.compareBigNumber(new BigNumber("2e24"), await baasFounder.founder2Supply(), "founder 2 supply");
        });

        it("should initialize correctly", async () => {
            await assertions.assert(baasFounder, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("0"),
                hasFounder1ReceivedTokens: false,
                hasFounder2ReceivedTokens: false,
            });
        });

        it("should have enough tokens after BaasToken was initialized", async () => {
            let token = await baasToken.setup(accounts[0], accounts[1], baasFounder.address, accounts[2]);

            await assertions.assert(baasFounder, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("10e24"),
                hasFounder1ReceivedTokens: false,
                hasFounder2ReceivedTokens: false,
            });
        });
    });


    describe('Issue Token', function () {
        it("should fail with wrong owner", async () => {
            let shouldFail = await baasFounder.issue(accounts[9], 0, {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should issue founder2 tokens correctly", async () => {
            let result = await baasFounder.issue(accounts[9],  1, {from: accounts[0]})
            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "TokensIssued",
                arg:[
                    accounts[9],
                    1,
                    new BN("2 000 000 000 000 000 000 000 000")
                ],
                argLength: 3
            });

            await assertions.assert(baasFounder, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("8e24"),
                hasFounder1ReceivedTokens: false,
                hasFounder2ReceivedTokens: true,
            });
        });


        it("founder 2 cannot receive again", async () => {
            let shouldFail = await baasFounder.issue(accounts[9], 1, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "founder 2 cannot receive again");
        });

        it("should issue founder1 tokens correctly", async () => {
            let result = await baasFounder.issue(accounts[9],  0, {from: accounts[0]})
            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "TokensIssued",
                arg:[
                    accounts[9],
                    0,
                    new BN("8 000 000 000 000 000 000 000 000")
                ],
                argLength: 3
            });

            await assertions.assert(baasFounder, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("0"),
                hasFounder1ReceivedTokens: true,
                hasFounder2ReceivedTokens: true,
            });
        });


        it("founder 1 cannot receive again", async () => {
            let shouldFail = await baasFounder.issue(accounts[9], 0, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "founder 1 cannot receive again");
        });
    });
});