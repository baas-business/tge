var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var utils = require('../utils')



module.exports = {
    assert: async function (contractPP, expected) {
        assert.equal(await contractPP.isFinalized(), expected.isFinalized, "isFinalized state is wrong");
        assert.equal(await contractPP.tokenAddress(), expected.tokenAddress, "tokenAddress is wrong");

        utils.compareBigNumber(expected.balance, await contractPP.balance(), "balance");
        utils.compareBigNumber(expected.tokensIssued0, await contractPP.tokensIssued(0), "tokens issued 0");
        utils.compareBigNumber(expected.tokensIssued1, await contractPP.tokensIssued(1), "tokens issued 1");
        utils.compareBigNumber(expected.tokensIssued2, await contractPP.tokensIssued(2), "tokens issued 2");
    }
};