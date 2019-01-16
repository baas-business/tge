var utils = require('../utils')
var BN = web3.utils.BN;

module.exports = {

    assert: async function (contract, expected) {
        utils.compareBigNumber(expected.balance, await contract.balance(), "balance");
        utils.compareBigNumber(expected.circulatingSupply, await contract.circulatingSupply(), "circulatingSupply");
        utils.compareBigNumber(expected.eligibleToken, await contract.eligibleToken(), "eligibleToken");


        for (let i = 0; i < expected.tokenEuroConversionRate.length; i++) {
            assert.equal(expected.contractHasEnoughTokensForPayout[i], await contract.contractHasEnoughTokensForPayout(expected.tokenEuroConversionRate[i]), "contractHasEnoughTokensForPayout " + expected.tokenEuroConversionRate[i]);
            utils.compareBigNumber(expected.currentPayoutObligation[i], await contract.currentPayoutObligation(expected.tokenEuroConversionRate[i]), "currentPayoutObligation " + expected.tokenEuroConversionRate[i]);

            let opd = await contract.optimalPayoutDistribution(expected.tokenEuroConversionRate[i]);

            utils.compareBigNumber(expected.optimalPayoutDistribution[i].contractCanPayMax, opd.contractCanPayMax, "contractCanPayMax " + expected.tokenEuroConversionRate[i]);
            utils.compareBigNumber(expected.optimalPayoutDistribution[i].contractShouldHaveMin, opd.contractShouldHaveMin, "contractShouldHaveMin " + expected.tokenEuroConversionRate[i]);
            utils.compareBigNumber(expected.optimalPayoutDistribution[i].error, opd.error, "error " + expected.tokenEuroConversionRate[i]);
        }
    }
};