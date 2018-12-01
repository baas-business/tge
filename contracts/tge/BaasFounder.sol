pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


interface IBaasFounder {
    function setup(address founder1, address founder2, uint256 vestingStart, uint256 vestingPeriod) external returns (bool);

    event FounderChanged(address oldAddress, address newAddress, int index);
}

contract BaasFounder is Ownable, IBaasFounder {
    using SafeMath for uint256;

    string private constant NAME = "FOUNDER";

    uint256 constant FOUNDER1_SUPPLY = 8 * 10 ** 18;                // 8m Founder1 Token
    uint256 constant FOUNDER2_SUPPLY = 2 * 10 ** 18;                // 2m Founder2 Token

    address private _founder1;
    address private _founder2;

    uint256 private _vestingStart;
    uint256 private _vestingPeriod;

    IBaasToken private _token;

    constructor(IBaasToken token) public {
        _token = token;
    }

    function setup(
        address founder1,
        address founder2,
        uint256 vestingStart,
        uint256 vestingPeriod)
    external onlyOwner returns (bool) {
        _founder1 = founder1;
        _founder2 = founder2;
        _vestingStart = vestingStart;
        _vestingPeriod = vestingPeriod;
    }


    function changeFounder(address newAddress, int index)
    external onlyOwner returns (bool) {
        address oldAddress;
        if (index == 0) {
            oldAddress = _founder1;
            _founder1 = newAddress;
        } else {
            oldAddress = _founder2;
            _founder2 = newAddress;
        }

        emit FounderChanged(oldAddress, newAddress, index);

        return true;
    }

    /*
        VIEWS
    */
    function Founder1() public view returns (address) {
        return _founder1;
    }

    function Founder2() public view returns (address) {
        return _founder2;
    }

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