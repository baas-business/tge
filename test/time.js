const {ethGetBlock} = require('./web3');
const {promisify} = require('util');

async function advanceBlockSteps(steps) {
    for (let i = 0; i < steps; i++) {
        await advanceBlock();
    }
}

function advanceBlock() {
    return promisify(web3.currentProvider.sendAsync)({
        jsonrpc: '2.0',
        method: 'evm_mine',
    });
}

// Returns the time of the last mined block in seconds
async function latest() {
    const block = await ethGetBlock('latest');
    return block.timestamp;
}

// Returns the time of the last mined block in seconds
async function latestBlockNumber() {
    const block = await ethGetBlock('latest');
    return block.number;
}

// Increases ganache time by the passed duration in seconds
async function increase(duration) {
    if (duration < 0) throw Error(`Cannot increase time by a negative amount (${duration})`);

    await promisify(web3.currentProvider.sendAsync)({
        jsonrpc: '2.0',
        method: 'evm_increaseTime',
        params: [duration],
    });

    await advanceBlock();
}

/**
 * Beware that due to the need of calling two separate ganache methods and rpc calls overhead
 * it's hard to increase time precisely to a target point so design your test to tolerate
 * small fluctuations from time to time.
 *
 * @param target time in seconds
 */
async function increaseTo(target) {
    const now = (await latest());

    if (target < now) throw Error(`Cannot increase current time (${now}) to a moment in the past (${target})`);
    const diff = target - now;
    return increase(diff);
}

const duration = {
    seconds: function (val) {
        return val;
    },
    minutes: function (val) {
        return val * this.seconds(60);
    },
    hours: function (val) {
        return val * this.minutes(60);
    },
    days: function (val) {
        return val * this.hours(24);
    },
    weeks: function (val) {
        return val * this.days(7);
    },
    years: function (val) {
        return val * this.days(365);
    },
};

module.exports = {
    advanceBlock,
    latest,
    increase,
    increaseTo,
    duration,
    latestBlockNumber,
    advanceBlockSteps
};
