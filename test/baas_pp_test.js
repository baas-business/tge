var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var BigNumber = require('bignumber.js');
var utils = require('utils.js');
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
        console.log(token)

        let balance = await baasPP.balance();
        console.log(balance);

        const expectedBT = new BigNumber('1e9');

        assert.equal(expectedBT.isEqualTo(balance), true, "baasPP should have " + expectedBT.toString() + " token but has " + balance.toString() + "!");
    });
});