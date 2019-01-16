var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var utils = require('../utils.js');
var BigNumber = require('bignumber.js');
var BN = web3.utils.BN;
var assertions = require('./assertions');

let baasToken;
let baasPP;
let escrowAddress;
let ppAddress;
let founderAddress;
let incentivesAddress;
let tokenReceiver1;
let tokenReceiver2;

contract('BaasToken', function (accounts) {
    beforeEach(async function () {
        baasToken = await BaasToken.deployed();
        baasPP = await BaasPP.deployed();

        escrowAddress = accounts[0];
        ppAddress = accounts[1];
        founderAddress = accounts[2];
        incentivesAddress = accounts[3];
        tokenReceiver1 = accounts[4];
        tokenReceiver2 = accounts[5];
    });

    describe('Setup', function () {
        it("should initialize with 0 total supply and isInitialized false", async () => {
            assert.equal(0, await baasToken.totalSupply(), "totalSupply should be 0");
            assert.equal(false, await baasToken.isInitialized(), "isInitialized should be false");
            assert.equal(accounts[0], await baasToken.owner(), "owner should be contract deployer");
            let transferResult = await baasToken.transfer(accounts[0], 200).catch(e => false);
            assert.equal(false, transferResult, "transfer should fail")

            assert.equal("BaaS Token", await baasToken.name(), "name");
            assert.equal("BaaS", await baasToken.symbol(), "symbol");
            assert.equal(18, await baasToken.decimals(), "decimals");
        });


        it("should should only be callable by owner", async () => {
            let shouldFail = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress, {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")
        });

        it("should setup correctly", async () => {
            let result = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);

            let logs = result.receipt.logs;

            assert.equal(5, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Transfer",
                arg: [
                    "0x0000000000000000000000000000000000000000",
                    escrowAddress,
                    new BN("60 000 000 000 000 000 000 000 000")
                ]
            });

            utils.compareEvent(logs[1], {
                event: "Transfer",
                arg: [
                    "0x0000000000000000000000000000000000000000",
                    ppAddress,
                    new BN("20 000 000 000 000 000 000 000 000")
                ]
            });

            utils.compareEvent(logs[2], {
                event: "Transfer",
                arg: [
                    "0x0000000000000000000000000000000000000000",
                    founderAddress,
                    new BN("10 000 000 000 000 000 000 000 000")
                ]
            });

            utils.compareEvent(logs[3], {
                event: "Transfer",
                arg: [
                    "0x0000000000000000000000000000000000000000",
                    incentivesAddress,
                    new BN("10 000 000 000 000 000 000 000 000")
                ]
            });

            utils.compareEvent(logs[4], {
                event: "SetupCompleted",
                arg: []
            });

            assert.equal(true, await baasToken.isInitialized(), "isInitialized");
            utils.compareBigNumber(new BigNumber('100e24'), await baasToken.totalSupply(), "Total Supply");

            await assertions.assert(baasToken, {
                paused: false,
                potSupply: new BigNumber("100e24"),
                circulatingSupply: new BigNumber("0"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("60e24"),
                    },
                    {
                        address: ppAddress,
                        balance: new BigNumber("20e24"),
                    },
                    {
                        address: founderAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("10e24"),
                    }]
            });
        });

        it("should fail to setup again", async () => {
            let shouldFail = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "transaction reverted")
        });
    });


    describe('Transfer', function () {
        it("should update token holder list and balances", async () => {
            let amount = new BN("1 000 000 000 000 000 000 000 000");
            let result = await baasToken.transfer(tokenReceiver1, amount, {from: accounts[0]});
            let logs = result.receipt.logs;

            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Transfer",
                arg: [
                    escrowAddress,
                    tokenReceiver1,
                    new BN("1 000 000 000 000 000 000 000 000")
                ]
            });

            await assertions.assert(baasToken, {
                paused: false,
                potSupply: new BigNumber("99e24"),
                circulatingSupply: new BigNumber("1e24"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("59e24"),
                    },
                    {
                        address: ppAddress,
                        balance: new BigNumber("20e24"),
                    },
                    {
                        address: founderAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: tokenReceiver1,
                        balance: new BigNumber("1e24"),
                    }
                ]
            });
        });
    });


    describe('Set Pause', function () {
        it("should fail pause if not owner", async () => {
            let shouldFail = await baasToken.setPause(true, {from: accounts[1]}).catch(e => true);
            assert.equal(true, shouldFail, "transaction reverted")
        });

        it("should set pause to true", async () => {
            let amount = new BN("1 000 000 000 000 000 000 000 000");
            let result = await baasToken.setPause(true, {from: accounts[0]});
            let logs = result.receipt.logs;

            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Paused",
                arg: [true]
            });

            await assertions.assert(baasToken, {
                paused: true,
                potSupply: new BigNumber("99e24"),
                circulatingSupply: new BigNumber("1e24"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("59e24"),
                    },
                    {
                        address: ppAddress,
                        balance: new BigNumber("20e24"),
                    },
                    {
                        address: founderAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: tokenReceiver1,
                        balance: new BigNumber("1e24"),
                    }
                ]
            });
        });

        it("should fail pause when already paused", async () => {
            let shouldFail = await baasToken.setPause(true, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "transaction reverted")
        });

        it("should fail transfer when paused", async () => {
            let amount = new BN("2 000 000 000 000 000 000 000 000");
            let shouldFail = await baasToken.transfer(tokenReceiver2, amount, {from: escrowAddress}).catch(e => true);
            assert.equal(true, shouldFail, "transaction reverted")
        });

        it("should set pause to false", async () => {
            let result = await baasToken.setPause(false, {from: accounts[0]});
            let logs = result.receipt.logs;

            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Paused",
                arg: [false]
            });

            await assertions.assert(baasToken, {
                paused: false,
                potSupply: new BigNumber("99e24"),
                circulatingSupply: new BigNumber("1e24"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("59e24"),
                    },
                    {
                        address: ppAddress,
                        balance: new BigNumber("20e24"),
                    },
                    {
                        address: founderAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: tokenReceiver1,
                        balance: new BigNumber("1e24"),
                    }
                ]
            });
        });


        it("should fail unpause when already unpaused", async () => {
            let shouldFail = await baasToken.setPause(false, {from: accounts[0]}).catch(e => true);
            assert.equal(true, shouldFail, "transaction reverted")
        });

        it("should succeed transfer when unpaused", async () => {
            let amount = new BN("2 000 000 000 000 000 000 000 000");
            let result = await baasToken.transfer(tokenReceiver2, amount, {from: escrowAddress});

            let logs = result.receipt.logs;

            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Transfer",
                arg: [
                    escrowAddress,
                    tokenReceiver2,
                    new BN("2 000 000 000 000 000 000 000 000")
                ]
            });

            await assertions.assert(baasToken, {
                paused: false,
                potSupply: new BigNumber("97e24"),
                circulatingSupply: new BigNumber("3e24"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("57e24"),
                    },
                    {
                        address: ppAddress,
                        balance: new BigNumber("20e24"),
                    },
                    {
                        address: founderAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: tokenReceiver1,
                        balance: new BigNumber("1e24"),
                    },
                    {
                        address: tokenReceiver2,
                        balance: new BigNumber("2e24"),
                    }
                ]
            });
        });
    });

    describe('burn tokens from pot', function () {
        it("should fail if sender is not pot address", async () => {
            let amount = new BN("10 000 000 000 000 000 000 000 000");
            let shouldFail = await baasToken.burnTokensFromPot(tokenReceiver1, amount, {from: tokenReceiver1}).catch(e => true);
            assert.equal(true, shouldFail, "transaction reverted");
        });

        it("should fail if msg.sender is not pot address", async () => {
            let amount = new BN("10 000 000 000 000 000 000 000 000");
            let shouldFail = await baasToken.burnTokensFromPot(tokenReceiver1, amount, {from: escrowAddress}).catch(e => true);
            assert.equal(true, shouldFail, "transaction reverted");
        });

        it("should fail if amount is too high", async () => {
            let amount = new BN("65 000 000 000 000 000 000 000 000");
            let shouldFail = await baasToken.burnTokensFromPot(escrowAddress, amount, {from: escrowAddress}).catch(e => true);
            assert.equal(true, shouldFail, "transaction reverted");
        });

        it("should allow escrow to burn tokens", async () => {
            let amount = new BN("20 000 000 000 000 000 000 000 000");
            let result = await baasToken.burnTokensFromPot(escrowAddress, amount, {from: escrowAddress});

            let logs = result.receipt.logs;

            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Transfer",
                arg: [
                    escrowAddress,
                    "0x0000000000000000000000000000000000000000",
                    new BN("20 000 000 000 000 000 000 000 000")
                ]
            });

            await assertions.assert(baasToken, {
                paused: false,
                potSupply: new BigNumber("77e24"),
                circulatingSupply: new BigNumber("3e24"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("37e24"),
                    },
                    {
                        address: ppAddress,
                        balance: new BigNumber("20e24"),
                    },
                    {
                        address: founderAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: tokenReceiver1,
                        balance: new BigNumber("1e24"),
                    },
                    {
                        address: tokenReceiver2,
                        balance: new BigNumber("2e24"),
                    }
                ]
            });
        });

        it("should allow pp to burn tokens", async () => {
            let amount = new BN("20 000 000 000 000 000 000 000 000");
            let result = await baasToken.burnTokensFromPot(ppAddress, amount, {from: ppAddress});

            let logs = result.receipt.logs;

            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Transfer",
                arg: [
                    ppAddress,
                    "0x0000000000000000000000000000000000000000",
                    new BN("20 000 000 000 000 000 000 000 000")
                ]
            });

            await assertions.assert(baasToken, {
                paused: false,
                potSupply: new BigNumber("57e24"),
                circulatingSupply: new BigNumber("3e24"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("37e24"),
                    },
                    {
                        address: tokenReceiver2,
                        balance: new BigNumber("2e24"),
                    },
                    {
                        address: founderAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("10e24"),
                    },
                    {
                        address: tokenReceiver1,
                        balance: new BigNumber("1e24"),
                    }
                ]
            });
        });

        it("should allow founder to burn tokens", async () => {
            let amount = new BN("10 000 000 000 000 000 000 000 000");
            let result = await baasToken.burnTokensFromPot(founderAddress, amount, {from: founderAddress});

            let logs = result.receipt.logs;

            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Transfer",
                arg: [
                    founderAddress,
                    "0x0000000000000000000000000000000000000000",
                    amount
                ]
            });

            await assertions.assert(baasToken, {
                paused: false,
                potSupply: new BigNumber("47e24"),
                circulatingSupply: new BigNumber("3e24"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("37e24"),
                    },
                    {
                        address: tokenReceiver2,
                        balance: new BigNumber("2e24"),
                    },
                    {
                        address: tokenReceiver1,
                        balance: new BigNumber("1e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("10e24"),
                    }
                ]
            });
        });

        it("should allow incentives to burn tokens", async () => {
            let amount = new BN("6 000 000 000 000 000 000 000 000");
            let result = await baasToken.burnTokensFromPot(incentivesAddress, amount, {from: incentivesAddress});

            let logs = result.receipt.logs;

            assert.equal(1, logs.length, "should emit 1 log");

            utils.compareEvent(logs[0], {
                event: "Transfer",
                arg: [
                    incentivesAddress,
                    "0x0000000000000000000000000000000000000000",
                    amount
                ]
            });

            await assertions.assert(baasToken, {
                paused: false,
                potSupply: new BigNumber("41e24"),
                circulatingSupply: new BigNumber("3e24"),
                tokenHolders: [
                    {
                        address: escrowAddress,
                        balance: new BigNumber("37e24"),
                    },
                    {
                        address: tokenReceiver2,
                        balance: new BigNumber("2e24"),
                    },
                    {
                        address: tokenReceiver1,
                        balance: new BigNumber("1e24"),
                    },
                    {
                        address: incentivesAddress,
                        balance: new BigNumber("4e24"),
                    }
                ]
            });
        });
    });

});
