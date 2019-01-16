var BigNumber = require('bignumber.js');

const toType = function(obj) {
    return ({}).toString.call(obj).match(/\s([a-zA-Z]+)/)[1].toLowerCase()
};

module.exports = {
    compareBigNumber: function (expected, actual, label) {
        assert.equal(expected.isEqualTo(actual), true, label + " should have " + expected.toString() + " but has " + actual.toString() + "");
    },

    compareEvent: function (actual, expected) {

        assert.equal(expected.event, actual.event);
        let length = 0;

        if(typeof actual.args['0'] !== "undefined") {
            assert.equal(expected.arg[0].toString(), actual.args['0'].toString(), "arg 1 is wrong for " + actual.event);
            length++;
        }

        if(typeof actual.args['1'] !== "undefined") {
            assert.equal(expected.arg[1].toString(), actual.args['1'].toString(), "arg 2 is wrong for " + actual.event);
            length++;
        }

        if(typeof actual.args['2'] !== "undefined") {
            assert.equal(expected.arg[2].toString(), actual.args['2'].toString(), "arg 3 is wrong for " + actual.event);
            length++;
        }

        assert.equal(expected.arg.length, length, "wrong argument length for " + actual.event)
    }
};