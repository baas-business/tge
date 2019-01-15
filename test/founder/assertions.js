var utils = require('../utils')



module.exports = {
    assert: async function (contract, expected) {
        assert.equal(await contract.tokenAddress(), expected.tokenAddress, "tokenAddress is wrong");

        utils.compareBigNumber(expected.balance, await contract.balance(), "balance");
        assert.equal(expected.hasFounder1ReceivedTokens, await contract.hasFounderReceivedTokens(0), "founder 1 token reception");
        assert.equal(expected.hasFounder2ReceivedTokens, await contract.hasFounderReceivedTokens(1), "founder 2 token reception");
    }
};