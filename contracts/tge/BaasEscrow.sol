pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


interface IBaasEscrow {
    function raiseCapital(uint256 amount) external returns (bool);

    event CapitalRaised(uint256 amount);
}


/*
    Question: any limitation on raising time?
*/
contract BaasEscrow is IBaasEscrow, Ownable {
    using SafeMath for uint256;

    string private constant NAME = "ESCROW";
    IBaasToken private _token;

    uint256 private raisedCapital;

    constructor(IBaasToken token) public {
        _token = token;
    }

    function raiseCapital(uint256 amount) public onlyOwner returns (bool) {
        require(amount <= balance());
        require(amount <= balance());

        emit CapitalRaised(amount);
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