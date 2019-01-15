pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


interface IBaasFounder {
    function setup() external returns (bool);

    function issue(address receiver, uint256 amount, bytes32 id, uint8 founderId) external returns (bool);

    event Issued(address indexed receiver, uint256 amount, uint8 founderId);

    event SetupCompleted();
}

contract BaasFounder is Ownable, IBaasFounder {
    using SafeMath for uint256;

    string private constant NAME = "FOUNDER";

    uint256 constant FOUNDER1_SUPPLY = 8 * 10 ** 24;                // 8m Founder1 Token
    uint256 constant FOUNDER2_SUPPLY = 2 * 10 ** 24;                // 2m Founder2 Token

    bool private _isInitialized = false;
    uint256[] private _left;

    IBaasToken private _token;

    constructor(IBaasToken token) public {
        _token = token;
    }

    function setup()
    external onlyOwner returns (bool) {
        require(balance() == FOUNDER1_SUPPLY.add(FOUNDER2_SUPPLY), "not enough tokens in contract");
        _left.push(FOUNDER1_SUPPLY);
        _left.push(FOUNDER2_SUPPLY);

        _isInitialized = true;
        emit SetupCompleted();
    }

    function issue(address receiver, uint256 amount, bytes32 id, int founderId)
    external onlyOwner returns (bool) {
        require(_isInitialized, "contract not initialized yet");
        require(founderId <= 1, "founderId not existent");
        require(amount <= balance(), "not enough tokens left");
        require(amount <= _left[founderId], "not enough tokens for this founder");
        require(_token.transfer(receiver, amount));

        _left[founderId] = _left[founderId].sub(amount);

        emit Issued(receiver, amount);

        return true;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  VIEWS
    ///////////////////////////////////////////////////////////////////////////////////////////

    function isInitialized() public view returns (bool) {
        return _isInitialized;
    }

    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    function tokenAddress() public view returns (address) {
        return _token;
    }

    function left(int founderId) public view returns (uint256){
        return _left[founderId];
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  PURE
    ///////////////////////////////////////////////////////////////////////////////////////////
    function name() public pure returns (string) {
        return NAME;
    }
}