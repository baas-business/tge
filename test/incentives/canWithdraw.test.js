var BaasToken = artifacts.require("./BaasToken.sol");
var BaasInc = artifacts.require("./BaasIncentives.sol");
var setup = require("./setup").setup;

const BigNumber = web3.BigNumber;
const should = require('chai')
    .use(require('chai-bignumber')(BigNumber))
    .should();

const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";

contract('BaasIncentives', function (accounts) {

    it("should provide incentive correctly", async () => {
        await setup(accounts);
        const baasInc = await BaasInc.deployed();
        arg = {
            receiver: accounts[5],
            amount: web3.toBigNumber("123e21"),
            vestingStages: 4,
            stagesBlockTime: 49
        };

        const result = await baasInc.reward(arg.receiver, arg.amount, arg.vestingStages, arg.stagesBlockTime);
        validateLogs(result, arg);
        await validateContractState(baasInc, arg);
    });
});


const validateLogs = (result, arg) => {
    assert.equal(2, result.logs.length, "logs size");

    let incentiveProvided1 = result.logs[0];
    assert.equal('IncentiveProvided1', incentiveProvided1.event);
    assert.equal(arg.receiver, incentiveProvided1.args.account);
    incentiveProvided1.args.amount.should.be.bignumber.equal(new BigNumber("123e21"));

    let incentiveProvided2 = result.logs[1];
    assert.equal('IncentiveProvided2', incentiveProvided2.event);
    assert.equal(arg.receiver, incentiveProvided2.args.account);
    incentiveProvided2.args.vestingStages.toNumber().should.be.equal(arg.vestingStages);
    incentiveProvided2.args.vestingBlocks.toNumber().should.be.equal(arg.stagesBlockTime);
};

const validateContractState = async (baasInc, arg) => {
    (await baasInc.incentivesProvided()).should.be.bignumber.equal(arg.amount);
    (await baasInc.incentivesLeft()).should.be.bignumber.equal(new BigNumber("9.877e+24"));

    const accounts = await baasInc.incentives();
    assert.equal(1, accounts.length);
    assert.equal(arg.receiver, accounts[0]);

    const i = await baasInc.getIncentive(arg.receiver);

    // uint256 initialAmount
    i[0].should.be.bignumber.equal(arg.amount);
    //     uint256 withdrawn
    i[1].should.be.bignumber.equal(0);
    //     uint atBlock
    const blocknumber = i[2];
    //     uint currentStage
    i[3].should.be.bignumber.equal(0);
    //     uint totalStages
    i[4].should.be.bignumber.equal(arg.vestingStages);
    //     uint totalBlocks
    i[5].should.be.bignumber.equal(arg.stagesBlockTime);
    //     bool[] stagesClaimed
    assert.equal(arg.vestingStages, i[6].length);
    i[6].forEach((e) => {
        assert.equal(false, e);
    });
    //     uint[] stagesBlockTime
    assert.equal(arg.vestingStages, i[7].length);
    i[7].forEach((e,i) => {
        e.should.be.bignumber.equal(blocknumber.add(arg.stagesBlockTime * i));
    });
    //     bool isValue
    assert.equal(true, i[8]);
};