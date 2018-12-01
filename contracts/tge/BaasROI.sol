pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";

interface IBaasROI {
}

contract BaasROI is IBaasROI, Ownable {
    using SafeMath for uint256;

    string private constant NAME = "RETURN OF INVESTMENT";

    IBaasToken private _token;

    constructor(IBaasToken token) public {
        _token = token;
    }



    // Views

    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    function tokenAddress() public view returns (address) {
        return _token;
    }

    function name() public pure returns (string) {
        return NAME;
    }
}