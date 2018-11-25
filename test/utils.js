

var BigNumber = require('bignumber.js');

module.exports =  function utils(){
    function compareBigNumber(expected, actual, label) {
        assert.equal(expected.isEqualTo(actual), true, label + "should have " + expected.toString() + " but has " + actual.toString() + "");
    }

};