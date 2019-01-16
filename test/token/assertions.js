var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var utils = require('../utils')



module.exports = {
    assert: async function (contract, expected) {
        assert.equal(await contract.paused(), expected.paused, "paused");

        utils.compareBigNumber(expected.potSupply, await contract.potSupply(), "potSupply");
        utils.compareBigNumber(expected.circulatingSupply, await contract.circulatingSupply(), "circulatingSupply");


        let tokenHolders = await contract.tokenHolderSnapShot();

        assert.equal(expected.tokenHolders.length, await contract.tokenHolderCount(), "tokenHolderCount");
        assert.equal(expected.tokenHolders.length, tokenHolders.length, "tokenHolders length");

        for(let i = 0; i < tokenHolders.length; i++) {
            assert.equal(expected.tokenHolders[i].address, tokenHolders[i], "tokenHolder " + i);
            utils.compareBigNumber(expected.tokenHolders[i].balance, await contract.balanceOf(expected.tokenHolders[i].address), "token holder balance " + i);
        }
    }
};