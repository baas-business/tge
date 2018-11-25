var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");

module.exports = function (deployer, network, accounts) {
    deployer.deploy(BaasToken)
        .then(function () {
            return deployer.deploy(BaasPP, BaasToken.address);
        });
};
