var ItemRegistry = artifacts.require("./UserRegistry.sol");

module.exports = function(deployer) {
    deployer.deploy(ItemRegistry);
};
