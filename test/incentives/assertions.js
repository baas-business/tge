var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var utils = require('../utils')



module.exports = {
    assert: async function (contract, expected) {
        assert.equal(await contract.tokenAddress(), expected.tokenAddress, "tokenAddress is wrong");

        utils.compareBigNumber(expected.balance, await contract.balance(), "balance");
        utils.compareBigNumber(expected.incentivesLeft, await contract.incentivesLeft(), "incentives left");
        utils.compareBigNumber(expected.incentivesIssued, await contract.incentivesIssued(), "incentives issued");
    }
};