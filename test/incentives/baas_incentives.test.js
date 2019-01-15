var BaasToken = artifacts.require("./BaasToken.sol");
var BaasInc = artifacts.require("./BaasIncentives.sol");
var assertions = require('./assertions');
var BigNumber = require('bignumber.js');
var BN = web3.utils.BN;
var utils = require('../utils');

let baasToken;
let baasInc;

contract('BaasPP', function (accounts) {

    beforeEach(async function () {
        baasToken = await BaasToken.deployed();
        baasInc = await BaasInc.deployed();
    });

    describe('Setup', function () {
        it("should correct constant values", async () => {
            utils.compareBigNumber(new BigNumber("10e24"), await baasInc.initialSupply(), "total cap");
        });

        it("should initialize correctly", async () => {
            await assertions.assert(baasInc, {
                isFinalized: false,
                tokenAddress: baasToken.address,
                balance: new BigNumber("0"),
                incentivesLeft: new BigNumber("10e24"),
                incentivesIssued: new BigNumber("0"),
            });
        });

        it("should have enough tokens after BaasToken was initialized", async () => {
            await baasToken.setup(accounts[0], accounts[1], accounts[2], baasInc.address);

            await assertions.assert(baasInc, {
                isFinalized: false,
                tokenAddress: baasToken.address,
                balance: new BigNumber('10e24'),
                incentivesLeft: new BigNumber("10e24"),
                incentivesIssued: new BigNumber("0"),
            });
        });
    });


    describe('Issue Token', function () {
        it("should fail with wrong owner", async () => {
            let shouldFail = await baasInc.issue(accounts[9], new BN("1000000000000000000000000"), {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should issue tokens correctly", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let result = await baasInc.issue(accounts[9], amount,  {from: accounts[0]});
            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 2 logs");

            utils.compareEvent(logs[0], {
                event: "TokensIssued",
                arg:[
                    accounts[9],
                    amount
                ],
                argLength: 2
            });
            await assertions.assert(baasInc, {
                isFinalized: false,
                tokenAddress: baasToken.address,
                balance: new BigNumber('9e24'),
                incentivesLeft: new BigNumber("9e24"),
                incentivesIssued: new BigNumber("1e24"),
            });
        });

        it("should fail when issued discounted tokens above cap", async () => {
            const amount = new BN("9 100 000 000 000 000 000 000 000");
            let result = await baasInc.issue(accounts[9], amount, 1, {from: accounts[0]}).catch(e => true);
            assert.equal(true, result, "should fail when above cap")
        });
    });
});