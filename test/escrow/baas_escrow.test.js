var BaasToken = artifacts.require("./BaasToken.sol");
var BaasEscrow = artifacts.require("./BaasEscrow.sol");
var assertions = require('./assertions');
var BigNumber = require('bignumber.js');
var BN = web3.utils.BN;
var utils = require('../utils');

let baasToken;
let baasEscrow;
let capitalRaiseBlock;
let capitalRaiseBlock2;

contract('BaasFounder', function (accounts) {

    beforeEach(async function () {
        baasToken = await BaasToken.deployed();
        baasEscrow = await BaasEscrow.deployed();
    });

    describe('Setup', function () {
        it("should have correct constant values", async () => {
            utils.compareBigNumber(new BigNumber("60e24"), await baasEscrow.initialSupply(), "initial supply");
        });

        it("should initialize correctly", async () => {
            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("0"),
                isCapitalRaiseOngoing: false,
                currentCapitalRaiseId: 0,
                raisedCapitalTotal: new BigNumber("0"),
                issuedTokenTotal: new BigNumber("0"),
                burnedTokenTotal: new BigNumber("0")
            });
        });

        it("should have enough tokens after BaasToken was initialized", async () => {
            let token = await baasToken.setup(baasEscrow.address, accounts[0], accounts[1], accounts[2]);

            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("60e24"),
                isCapitalRaiseOngoing: false,
                currentCapitalRaiseId: 0,
                raisedCapitalTotal: new BigNumber("0"),
                issuedTokenTotal: new BigNumber("0"),
                burnedTokenTotal: new BigNumber("0")
            });
        });
    });

    describe('Raise Capital', function () {
        it("should fail with wrong owner", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let shouldFail = await baasEscrow.raiseCapital(amount, 5, {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should fail capital raise id 0", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let shouldFail = await baasEscrow.raiseCapital(amount, 0, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "capital raise id == 0")
        });

        it("should fail with to much capital raised", async () => {
            const amount = new BN("61 000 000 000 000 000 000 000 000");
            let shouldFail = await baasEscrow.raiseCapital(amount, 5, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "too much capital raised")
        });

        it("should correctly raise capital", async () => {
            const amount = new BN("20 000 000 000 000 000 000 000 000");
            let result = await baasEscrow.raiseCapital(amount, 5, {from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "CapitalRaised",
                arg: [
                    5,
                    new BN("20 000 000 000 000 000 000 000 000")
                ],
                argLength: 2
            });

            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("60e24"),
                isCapitalRaiseOngoing: true,
                currentCapitalRaiseId: 5,
                raisedCapitalTotal: new BigNumber("20e24"),
                issuedTokenTotal: new BigNumber("0"),
                burnedTokenTotal: new BigNumber("0")
            });

            capitalRaiseBlock = result.receipt.blockNumber;

            await assertions.assertCapitalRaise(baasEscrow, [
                {
                    id: new BigNumber("5"),
                    amount: new BigNumber("20e24"),
                    provided: new BigNumber("0"),
                    atBlock: new BigNumber(capitalRaiseBlock),
                    isFinalized: false,
                    isValue: true
                }
            ])
        });

        it("should not be able to raise more capital as capital raise is ongoing", async () => {
            const amount = new BN("10 000 000 000 000 000 000 000 000");
            let shouldFail = await baasEscrow.raiseCapital(amount, 6, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "ongoing capital raise")
        });

    });

    describe('Issue Token', function () {
        it("should fail with wrong owner", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let shouldFail = await baasEscrow.issue(accounts[9], amount, {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should issue tokens correctly", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let result = await baasEscrow.issue(accounts[9], amount, {from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "TokensIssued",
                arg: [
                    accounts[9],
                    amount
                ],
                argLength: 2
            });

            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("59e24"),
                isCapitalRaiseOngoing: true,
                currentCapitalRaiseId: 5,
                raisedCapitalTotal: new BigNumber("20e24"),
                issuedTokenTotal: new BigNumber("1e24"),
                burnedTokenTotal: new BigNumber("0")
            });

            await assertions.assertCapitalRaise(baasEscrow, [
                {
                    id: new BigNumber("5"),
                    amount: new BigNumber("20e24"),
                    provided: new BigNumber("1e24"),
                    atBlock: new BigNumber(capitalRaiseBlock),
                    isFinalized: false,
                    isValue: true
                }
            ])
        });

        it("should issue tokens correctly second time", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let result = await baasEscrow.issue(accounts[9], amount, {from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "TokensIssued",
                arg: [
                    accounts[9],
                    amount
                ],
                argLength: 2
            });

            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("58e24"),
                isCapitalRaiseOngoing: true,
                currentCapitalRaiseId: 5,
                raisedCapitalTotal: new BigNumber("20e24"),
                issuedTokenTotal: new BigNumber("2e24"),
                burnedTokenTotal: new BigNumber("0")
            });

            await assertions.assertCapitalRaise(baasEscrow, [
                {
                    id: new BigNumber("5"),
                    amount: new BigNumber("20e24"),
                    provided: new BigNumber("2e24"),
                    atBlock: new BigNumber(capitalRaiseBlock),
                    isFinalized: false,
                    isValue: true
                },
            ])
        });
    });


    describe('Finalize', function () {
        it("should fail with wrong owner", async () => {
            let shouldFail = await baasEscrow.finalize({from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "finalize can only be called by owner")
        });

        it("should finalize correctly", async () => {
            let result = await baasEscrow.finalize({from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "Finalized",
                arg: [
                    5,
                    new BN("18 000 000 000 000 000 000 000 000")
                ],
                argLength: 2
            });

            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("40e24"),
                isCapitalRaiseOngoing: false,
                currentCapitalRaiseId: 0,
                raisedCapitalTotal: new BigNumber("20e24"),
                issuedTokenTotal: new BigNumber("2e24"),
                burnedTokenTotal: new BigNumber("18e24")
            });

            await assertions.assertCapitalRaise(baasEscrow, [
                {
                    id: new BigNumber("5"),
                    amount: new BigNumber("20e24"),
                    provided: new BigNumber("2e24"),
                    atBlock: new BigNumber(capitalRaiseBlock),
                    isFinalized: true,
                    isValue: true
                }

            ]);
        });

        it("should fail when already finalized", async () => {
            let shouldFail = await baasEscrow.finalize({from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "finalize can only be called if capital raise is ongoing")
        });
    });


    describe('Raise Capital (Second Raise)', function () {

        it("should fail if capital raise id was already used", async () => {
            const amount = new BN("10 000 000 000 000 000 000 000 000");
            let shouldFail = await baasEscrow.raiseCapital(amount, 5, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "too much capital raised")
        });

        it("should correctly raise capital a second time", async () => {
            const amount = new BN("10 000 000 000 000 000 000 000 000");
            let result = await baasEscrow.raiseCapital(amount, 6, {from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "CapitalRaised",
                arg: [
                    6,
                    new BN("10 000 000 000 000 000 000 000 000")
                ],
                argLength: 2
            });

            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("40e24"),
                isCapitalRaiseOngoing: true,
                currentCapitalRaiseId: 6,
                raisedCapitalTotal: new BigNumber("30e24"),
                issuedTokenTotal: new BigNumber("2e24"),
                burnedTokenTotal: new BigNumber("18e24")
            });

            capitalRaiseBlock2 = result.receipt.blockNumber;

            await assertions.assertCapitalRaise(baasEscrow, [
                {
                    id: new BigNumber("5"),
                    amount: new BigNumber("20e24"),
                    provided: new BigNumber("2e24"),
                    atBlock: new BigNumber(capitalRaiseBlock),
                    isFinalized: true,
                    isValue: true
                },
                {
                    id: new BigNumber("6"),
                    amount: new BigNumber("10e24"),
                    provided: new BigNumber("0"),
                    atBlock: new BigNumber(capitalRaiseBlock2),
                    isFinalized: false,
                    isValue: true
                }
            ])
        });

        it("should not be able to raise more capital as capital raise is ongoing", async () => {
            const amount = new BN("10 000 000 000 000 000 000 000 000");
            let shouldFail = await baasEscrow.raiseCapital(amount, 7, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "ongoing capital raise")
        });
    });


    describe('Issue Token (Second Time)', function () {
        it("should issue tokens correctly", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let result = await baasEscrow.issue(accounts[9], amount, {from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "provision should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "TokensIssued",
                arg: [
                    accounts[9],
                    amount
                ],
                argLength: 2
            });

            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("39e24"),
                isCapitalRaiseOngoing: true,
                currentCapitalRaiseId: 6,
                raisedCapitalTotal: new BigNumber("30e24"),
                issuedTokenTotal: new BigNumber("3e24"),
                burnedTokenTotal: new BigNumber("18e24")
            });

            await assertions.assertCapitalRaise(baasEscrow, [
                {
                    id: new BigNumber("5"),
                    amount: new BigNumber("20e24"),
                    provided: new BigNumber("2e24"),
                    atBlock: new BigNumber(capitalRaiseBlock),
                    isFinalized: true,
                    isValue: true
                },
                {
                    id: new BigNumber("6"),
                    amount: new BigNumber("10e24"),
                    provided: new BigNumber("1e24"),
                    atBlock: new BigNumber(capitalRaiseBlock2),
                    isFinalized: false,
                    isValue: true
                }
            ])
        });
    });


    describe('Finalize (Second Time)', function () {
        it("should finalize correctly", async () => {
            let result = await baasEscrow.finalize({from: accounts[0]});

            let logs = result.receipt.logs;
            assert.equal(1, logs.length, "should emit 1 logs");

            utils.compareEvent(logs[0], {
                event: "Finalized",
                arg: [
                    6,
                    new BN("9 000 000 000 000 000 000 000 000")
                ],
                argLength: 2
            });

            await assertions.assert(baasEscrow, {
                tokenAddress: baasToken.address,
                balance: new BigNumber("30e24"),
                isCapitalRaiseOngoing: false,
                currentCapitalRaiseId: 0,
                raisedCapitalTotal: new BigNumber("30e24"),
                issuedTokenTotal: new BigNumber("3e24"),
                burnedTokenTotal: new BigNumber("27e24")
            });

            await assertions.assertCapitalRaise(baasEscrow, [
                {
                    id: new BigNumber("5"),
                    amount: new BigNumber("20e24"),
                    provided: new BigNumber("2e24"),
                    atBlock: new BigNumber(capitalRaiseBlock),
                    isFinalized: true,
                    isValue: true
                },
                {
                    id: new BigNumber("6"),
                    amount: new BigNumber("10e24"),
                    provided: new BigNumber("1e24"),
                    atBlock: new BigNumber(capitalRaiseBlock2),
                    isFinalized: true,
                    isValue: true
                }
            ]);
        });

        it("should fail when already finalized", async () => {
            let shouldFail = await baasEscrow.finalize({from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "finalize can only be called if capital raise is ongoing")
        });

        it("issue token should fail when already finalized", async () => {
            const amount = new BN("1 000 000 000 000 000 000 000 000");
            let shouldFail = await baasEscrow.finalize(amount, 9, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "issue token should fail")
        });
    });
});