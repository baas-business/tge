var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var utils = require('./utils.js');
var BigNumber = require('bignumber.js');
var BN = require('bn.js');
const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";

contract('BaasPP', function (accounts) {
    it("should initialize correctly", async () => {

        const baasToken = await BaasToken.deployed();
        const baasPP = await BaasPP.deployed();

        let tokenAddress = await baasPP.tokenAddress();
        assert.equal(tokenAddress, baasToken.address, "baasPP should have baasToken address registered!");

        let balance = await baasPP.balance();
        assert.equal(balance, 0, "baasPP should have 0 token!");
    });

    it("should have enough tokens after BaasToken was initialized", async () => {

        const baasToken = await BaasToken.deployed();
        const baasPP = await BaasPP.deployed();

        let token = await baasToken.setup(accounts[0], baasPP.address, accounts[1], accounts[2]).catch(e => console.log(e));
        let balance = await baasPP.balance();
        utils.compareBigNumber(new BigNumber('20e24'), balance, "PP Balance");

    });

    it("should provide right amount of discounted tokens", async () => {

        const baasToken = await BaasToken.deployed();
        const baasPP = await BaasPP.deployed();


        // first provision
        let firstAccountProvision = web3.toBigNumber("1000000000000000000000000");


        const investor = accounts[9];
        // correct modifier
        let shouldFail = await baasPP.provideToken(investor, firstAccountProvision, 1, {from: accounts[1]}).catch(e => true);
        assert.equal(true, shouldFail, "account that is not owner should no be able to execute function")

        //    console.log(firstAccountProvision.toNumber())
        let provide1 = await baasPP.provideToken(investor, firstAccountProvision, 1, {from: accounts[0]});


        // check logs
        let logs = provide1.receipt.logs;
        assert.equal(logs.length, 2, "provision should emit 2 logs");

        // log 1 Transfer
        assert.equal(logs[0].topics.length, 3, "first log should have 3 topics");
        assert.equal("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef", logs[0].topics[0], "topic1 of first log is wrong");
        assert.equal(baasPP.address.substr(2), logs[0].topics[1].substr(26), "topic2 of first log is wrong");
        assert.equal("0x000000000000000000000000"+investor.substr(2), logs[0].topics[2], "topic3 of first log is wrong");
        assert.equal("0x00000000000000000000000000000000000000000000d3c21bcecceda1000000", logs[0].data, "data of first log is wrong");

        // log 2 TokenDelivered
        assert.equal(logs[1].topics.length, 2, "second log should have 3 topics");
        assert.equal("0xa254c04eefa2b3459dff6e70a8a28dc5876401ccfb247ad27c56f0c32970cf06", logs[1].topics[0], "topic1 of second log is wrong");
        assert.equal("0x000000000000000000000000"+investor.substr(2), logs[1].topics[1], "topic2 of second log is wrong");
        let discountType = "0000000000000000000000000000000000000000000000000000000000000001"
        assert.equal("0x00000000000000000000000000000000000000000000d3c21bcecceda1000000" + discountType, logs[1].data, "data of second log is wrong");


        // balance is correct
        let balance = await baasToken.balanceOf(investor);
        utils.compareBigNumber(new BigNumber('1e24'), balance, "provided Balance ");


        let discounted = await baasPP.totalTokenProvided(1);
        utils.compareBigNumber(new BigNumber('1e24'), discounted, "total discounted tokens provided ");

        let total = await baasPP.totalTokenProvided(0);
        utils.compareBigNumber(new BigNumber('1e24'), total, "total tokens provided ");

        let not_discounted = await baasPP.totalTokenProvided(2);
        utils.compareBigNumber(new BigNumber('0'), not_discounted, "provided total not discounted ");

    });
});