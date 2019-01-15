var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var utils = require('../utils')



module.exports = {
    assert: async function (contractPP, expected) {
        assert.equal(await contractPP.tokenAddress(), expected.tokenAddress, "tokenAddress is wrong");

        utils.compareBigNumber(expected.balance, await contractPP.balance(), "balance");
        utils.compareBigNumber(expected.incentivesLeft, await contractPP.incentivesLeft(), "incentives left");
        utils.compareBigNumber(expected.incentivesIssued, await contractPP.incentivesIssued(), "incentives issued");
    }
};