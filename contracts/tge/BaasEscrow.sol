pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IERC20.sol";


contract IBaasEscrow {
    function tokenAddress() public view returns (address);

    function balance() public view returns (uint256);
}

contract BaasEscrow is Ownable {
    using SafeMath for uint256;

    IERC20 private _token;

    constructor(IERC20 token) public {
        _token = token;
    }

    function raiseCapital(uint256 amount) public onlyOwner returns (bool) {
        return true;
    }

    // Views

    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    function tokenAddress() public view returns (address) {
        return _token;
    }
}