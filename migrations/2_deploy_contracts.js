var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var BaasROI = artifacts.require("./BaasROI.sol");

module.exports = function (deployer, network, accounts) {
    deployer.deploy(BaasToken)
        .then(function () {
            return deployer.deploy(BaasPP, BaasToken.address);
        })
        .then(function () {
            return deployer.deploy(BaasROI, BaasToken.address);
        });
};
