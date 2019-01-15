var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var state = require('./assert');
const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";


var BigNumber = require('bignumber.js');
var BN = web3.utils.BN;
var utils = require('../utils');

let baasToken;
let baasPP;

contract('BaasPP', function (accounts) {

    beforeEach(async function () {
        baasToken = await BaasToken.deployed();
        baasPP = await BaasPP.deployed();
    });

    describe('Setup', function () {
        it("should correct constant values", async () => {
            utils.compareBigNumber(new BigNumber("20e24"), await baasPP.cap(0), "total cap");
            utils.compareBigNumber(new BigNumber("5e24"), await baasPP.cap(1), "discounted cap");
            utils.compareBigNumber(new BigNumber("15e24"), await baasPP.cap(2), "not discounted cap");
        });

        it("should initialize correctly", async () => {
            await state.assert(baasPP, {
                isFinalized: false,
                tokenAddress: baasToken.address,
                balance: new BigNumber("0"),
                tokensIssued0: new BigNumber("0"),
                tokensIssued1: new BigNumber("0"),
                tokensIssued2: new BigNumber("0"),
            });
        });

        it("should have enough tokens after BaasToken was initialized", async () => {
            let token = await baasToken.setup(accounts[0], baasPP.address, accounts[1], accounts[2]);

            await state.assert(baasPP, {
                isFinalized: false,
                tokenAddress: baasToken.address,
                balance: new BigNumber('20e24'),
                tokensIssued0: new BigNumber("0"),
                tokensIssued1: new BigNumber("0"),
                tokensIssued2: new BigNumber("0"),
            });
        });
    });


    describe('Issue Token', function () {
        it("should fail with wrong owner", async () => {
            let shouldFail = await baasPP.issue(accounts[9], new BN("1000000000000000000000000"), 1, {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should issue discounted tokens correctly", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let result = await baasPP.issue(accounts[9], amount, 1, {from: accounts[0]})
            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 2 logs");

            utils.compareEvent(logs[0], {
                event: "TokensIssued",
                arg:[
                    "0xF59f27a9c0a5f926983C6a69AEbA809fb6b25f3e",
                    1,
                    amount
                ],
                argLength: 3
            });

            await state.assert(baasPP, {
                isFinalized: false,
                tokenAddress: baasToken.address,
                balance: new BigNumber('19e24'),
                tokensIssued0: new BigNumber("1e24"),
                tokensIssued1: new BigNumber("1e24"),
                tokensIssued2: new BigNumber("0"),
            });

        });

        it("should fail when issued discounted tokens above cap", async () => {
            const amount = new BN("4 100 000 000 000 000 000 000 000");
            let result = await baasPP.issue(accounts[9], amount, 1, {from: accounts[0]}).catch(e => true);
            assert.equal(true, result, "should fail when above cap")
        });

        it("should issue NOT discounted tokens correctly", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let result = await baasPP.issue(accounts[9], amount, 2, {from: accounts[0]});

            // Logs
            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 2 logs");

            utils.compareEvent(logs[0], {
                event: "TokensIssued",
                arg:[
                    "0xF59f27a9c0a5f926983C6a69AEbA809fb6b25f3e",
                    2,
                    amount
                ],
                argLength: 3
            });

            // State
            await state.assert(baasPP, {
                isFinalized: false,
                tokenAddress: baasToken.address,
                balance: new BigNumber('18e24'),
                tokensIssued0: new BigNumber("2e24"),
                tokensIssued1: new BigNumber("1e24"),
                tokensIssued2: new BigNumber("1e24"),
            });
        });

        it("should fail when issued NOT discounted tokens above cap", async () => {
            const amount = new BN("14 100 000 000 000 000 000 000 000");
            let result = await baasPP.issue(accounts[9], amount, 2, {from: accounts[0]}).catch(e => true);
            assert.equal(true, result, "should fail when above cap")
        });
    });


    describe('Finalize', function () {
        it("should revert if not finalized by owner", async () => {
            let shouldFail = await baasPP.finalize({from: accounts[2]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should be finalized by owner", async () => {
            let result = await baasPP.finalize({from: accounts[0]})

            // Logs
            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Finalized",
                arg:[
                    new BN('18 000 000 000 000 000 000 000 000')
                ],
                argLength: 1
            });

            // State
            await state.assert(baasPP, {
                isFinalized: true,
                tokenAddress: baasToken.address,
                balance: new BigNumber('0'),
                tokensIssued0: new BigNumber("2e24"),
                tokensIssued1: new BigNumber("1e24"),
                tokensIssued2: new BigNumber("1e24"),
            });
        });
        it("should fail to issue tokens after finalization", async () => {
            const amount = new BN("1");
            let shouldFail = await baasPP.issue(accounts[9], amount, 2, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });
    });
});