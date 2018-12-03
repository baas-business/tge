var BaasToken = artifacts.require("./BaasToken.sol");
var BaasROI = artifacts.require("./BaasROI.sol");
const BigNumber = web3.BigNumber;
const should = require('chai')
    .use(require('chai-bignumber')(BigNumber))
    .should();


contract('BaasROI', function (accounts) {
    it("should initialize correctly", async () => {
            const escrowAddress = accounts[0];
            const ppAddress = accounts[1];
            const founderAddress = accounts[2];
            const incentivesAddress = accounts[3];


            const baasToken = await BaasToken.deployed();
            const baasROI = await BaasROI.deployed();
            let setup = await baasToken.setup(escrowAddress, ppAddress, founderAddress, incentivesAddress);

            let expected = new BigNumber("60e24");
            let result = await baasToken.balanceOf(escrowAddress);
            console.log(expected.toNumber() + " " + result.toNumber());
            console.log(expected + " " + result);
            console.log(expected.e);
            console.log(result.e);
            console.log(expected.c);
            console.log(result.c);
            console.log(expected.s);
            console.log(result.s);
            result.should.be.bignumber.equal(expected);

        }
    )
});