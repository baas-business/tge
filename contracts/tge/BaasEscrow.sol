pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IERC20.sol";


contract IBaasEscrow {
}

contract BaasEscrow is Ownable {
    using SafeMath for uint256;

    string private constant NAME = "ESCROW";
    IERC20 private _token;


    constructor(IERC20 token) public {
        _token = token;
    }

    function raiseCapital(uint256 amount) public onlyOwner returns (bool) {
        require(amount <= balance());
        return true;
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