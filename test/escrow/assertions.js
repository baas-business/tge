var utils = require('../utils')
var BN = web3.utils.BN;

module.exports = {
    assertCapitalRaise: async function (contract, expected) {
        cr = [];
        ids = await contract.capitalRaiseIds();


        for (var i = 0; i < ids.length; i++) {
            cr.push(await contract.capitalRaise(new BN(ids[i])));
        }

        assert.equal(expected.length, cr.length, "number of capital raises");

        for (var i = 0; i < cr.length; i++) {
            let capitalRaise = cr[i];
            utils.compareBigNumber(expected[i].id, capitalRaise.id, "id not equal");
            utils.compareBigNumber(expected[i].amount, capitalRaise.amount, "amount not equal");
            utils.compareBigNumber(expected[i].provided, capitalRaise.provided, "provided not equal");
            utils.compareBigNumber(expected[i].atBlock, capitalRaise.atBlock, "atBlock not equal");

            assert.equal(expected[i].isFinalized, capitalRaise.isFinalized, "isFinalized");
            assert.equal(expected[i].isValue, capitalRaise.isValue, "isValue");
        }
    },
    assert: async function (contract, expected) {
        assert.equal(await contract.tokenAddress(), expected.tokenAddress, "tokenAddress is wrong");

        utils.compareBigNumber(expected.balance, await contract.balance(), "balance");
        assert.equal(expected.isCapitalRaiseOngoing, await contract.isCapitalRaiseOngoing(), "capital raise ongoing");
        assert.equal(expected.currentCapitalRaiseId, await contract.currentCapitalRaiseId(), "current capital raise id");
        utils.compareBigNumber(expected.raisedCapitalTotal, await contract.raisedCapitalTotal(), "raised capital total");
        utils.compareBigNumber(expected.issuedTokenTotal, await contract.issuedTokenTotal(), "issued token total");
        utils.compareBigNumber(expected.burnedTokenTotal, await contract.burnedTokenTotal(), "burned token total");
    }
};