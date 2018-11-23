pragma solidity ^0.4.24;


import "../../math/SafeMath.sol";
import "../../ownership/Ownable.sol";


contract BaasIncentive is Ownable {
    using SafeMath for uint256;


    function raiseCapital(uint256 amount) public onlyOwner returns (bool) {
        return true;
    }
}