pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IERC20.sol";

interface IBaasROI {
}

contract BaasROI is IBaasROI, Ownable {
    using SafeMath for uint256;

    string private constant NAME = "RETURN OF INVESTMENT";

    IERC20 private _token;

    constructor(IERC20 token) public {
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