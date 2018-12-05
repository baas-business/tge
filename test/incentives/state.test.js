var BaasToken = artifacts.require("./BaasToken.sol");
var BaasInc = artifacts.require("./BaasIncentives.sol");
var setup = require("./setup").setup;
var {advanceBlockSteps, latestBlockNumber} = require('../time')

const BigNumber = web3.BigNumber;
const should = require('chai')
    .use(require('chai-bignumber')(BigNumber))
    .should();

const contractDeployer = "0x260502fd8202ad46e1e0cb555e4efa778e568e7c";

const STATE_CAN_CLAIM = 0;
const STATE_CAN_NOT_CLAIM_YET = 1;
const STATE_HAS_CLAIMED = 2;
const STATE_STAGE_NOT_AVAILABLE = 3;
const STATE_USER_NON_EXISTENT_OR_FORFEITED = 4;

contract('BaasIncentives', function (accounts) {

    it("should provide incentive correctly", async () => {
        await setup(accounts);
        const baasInc = await BaasInc.deployed();
        arg = {
            receiver: accounts[5],
            amount: web3.toBigNumber("123e21"),
            vestingStages: 4,
            stagesBlockTime: 49,
            atBlock: -1
        };


        let t = await baasInc.reward(arg.receiver, arg.amount, arg.vestingStages, arg.stagesBlockTime);
        arg.atBlock = t.receipt.blockNumber;


        const i = await baasInc.getIncentive(arg.receiver);
        console.log(i[6]);
        console.log("BlockNumber: " +await latestBlockNumber());
        await confirmStates(arg, baasInc, [STATE_CAN_CLAIM, STATE_CAN_NOT_CLAIM_YET, STATE_CAN_NOT_CLAIM_YET, STATE_CAN_NOT_CLAIM_YET, STATE_STAGE_NOT_AVAILABLE])
        await advanceBlockSteps(arg.stagesBlockTime);
        console.log("BlockNumber: " +await latestBlockNumber());
        arg.atBlock = await latestBlockNumber();
        await confirmStates(arg, baasInc, [STATE_CAN_CLAIM, STATE_CAN_CLAIM, STATE_CAN_NOT_CLAIM_YET, STATE_CAN_NOT_CLAIM_YET, STATE_STAGE_NOT_AVAILABLE])
    });
});

const confirmStates = async (arg, contract, states) => {

    console.log(states);
    console.log(arg.atBlock);
    for (let i = 0; i < states.length; i++) {
        let s = await contract.state(arg.receiver, i, arg.atBlock);
        assert.equal(states[i], s)
        console.log(states[i]+ " "+ s)
    }
};
