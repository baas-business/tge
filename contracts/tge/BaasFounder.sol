pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


interface IBaasFounder {
    function setup(uint256 vestingPeriod) external returns (bool);

    event FounderWithdraw(address indexed receiver, uint256 amount);

    event SetupCompleted(uint vestingStartBlock, uint vestingPeriod, uint vestingEndBlock);
}

contract BaasFounder is Ownable, IBaasFounder {
    using SafeMath for uint256;
    using SafeMath for uint;

    string private constant NAME = "FOUNDER";

    uint256 constant FOUNDER1_SUPPLY = 8 * 10 ** 18;                // 8m Founder1 Token
    uint256 constant FOUNDER2_SUPPLY = 2 * 10 ** 18;                // 2m Founder2 Token

    uint256 private _vestingStart;
    uint256 private _vestingPeriod;

    bool private _isInitialized = false;

    IBaasToken private _token;

    constructor(IBaasToken token) public {
        _token = token;
    }

    function setup(uint256 vestingPeriod)
    external onlyOwner returns (bool) {
        _vestingStart = block.number;
        _vestingPeriod = vestingPeriod;
        _isInitialized = true;
        emit SetupCompleted(_vestingStart, _vestingPeriod, _vestingStart.add(_vestingPeriod));
    }

    function withdraw(address receiver, uint256 amount)
    external onlyOwner returns (bool) {
        require(_isInitialized);
        require(canWithdraw(block.number));
        require(amount <= balance());

        require(_token.transfer(receiver, amount));

        emit FounderWithdraw(receiver, amount);

        return true;
    }

    /*
        VIEWS
    */

    function isInitialized() public view returns (bool) {
        return _isInitialized;
    }

    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    function tokenAddress() public view returns (address) {
        return _token;
    }

    function vestingStartBlock() public view returns (uint) {
        return _vestingStart;
    }

    function vestingPeriod() public view returns (uint) {
        return _vestingPeriod;
    }

    function vestingEndBlock() public view returns (uint) {
        return _vestingStart.add(_vestingPeriod);
    }

    function canWithdraw(uint blocknumber) public view returns (bool) {
        return vestingEndBlock() <= blocknumber;
    }

    /*
        pure
    */
    function name() public pure returns (string) {
        return NAME;
    }
}