var BaasToken = artifacts.require("./BaasToken.sol");
var BaasPP = artifacts.require("./BaasPP.sol");
var BaasROI = artifacts.require("./BaasROI.sol");
var BaasFounder = artifacts.require("./BaasFounder.sol");
var BaasInc = artifacts.require("./BaasIncentives.sol");
var BaasEscrow = artifacts.require("./BaasEscrow.sol");

module.exports = function (deployer, network, accounts) {
    deployer.deploy(BaasToken)
        .then(function () {
            return deployer.deploy(BaasPP, BaasToken.address);
        })
        .then(function () {
            return deployer.deploy(BaasROI, BaasToken.address);
        })
        .then(function () {
            return deployer.deploy(BaasInc, BaasToken.address);
        })
        .then(function () {
            return deployer.deploy(BaasEscrow, BaasToken.address);
        })
        .then(function () {
            return deployer.deploy(BaasFounder, BaasToken.address);
        })
};
