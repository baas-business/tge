pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";

interface IBaasIncentives {
    event IncentiveProvided(address indexed account, bytes32 indexed id, uint256 amount);
    event Payout(address indexed account, bytes32 indexed id, uint256 amount);
    event Revoked(address indexed account, bytes32 indexed id, uint256 tokenNotDelivered);

    event SetupCompleted(uint256 supply);
}
/*
    Can claim first stage immediately?
*/
contract BaasIncentives is IBaasIncentives, Ownable {
    using SafeMath for uint256;
    using SafeMath for uint;

    string private constant NAME = "INCENTIVES";
    uint256 private constant INITIAL_SUPPLY = 10 * 10 ** 24;

    struct Incentive {
        address account;        // the beneficiary of this incentive
        uint256 amount;
        uint256 provided;
        bytes32 id;
        bool isValue;
    }

    mapping(address => Incentive) _incentives;
    address[] private _incentivesList;

    uint256 private _incentivesLeft;
    uint256 private _incentivesProvided;

    bool private _isInitialized = false;

    IBaasToken private _token;

    constructor(IBaasToken token) public {
        _token = token;
    }

    function setup() external {
        require(!_isInitialized);
        _incentivesLeft = balance();
        _incentivesProvided = 0;
        _isInitialized = true;
        emit SetupCompleted(_incentivesLeft);
    }

    function reserve(address account, uint256 amount, bytes32 id) external onlyOwner returns (bool) {
        require(_isInitialized, "contract must be initialized");
        require(!_incentives[account].isValue, "address should not been incentivized yet");

        // update incentive storage
        _incentives[account] = Incentive(account, amount, 0, id, true);
        _incentivesList.push(account);

        // update contract balance
        _incentivesLeft = _incentivesLeft.sub(amount);
        _incentivesProvided = _incentivesProvided.add(amount);

        emit IncentiveProvided(account, id, amount);

        return true;
    }

    function revoke(address account) external onlyOwner returns (bool) {
        require(_incentives[account].isValue, "address was never incentivized");
        _incentives[account].isValue = false;

        uint256 remaining = _incentives[account].amount - _incentives[account].provided;
        _incentivesLeft = _incentivesLeft.add(remaining);
        _incentivesProvided = _incentivesProvided.sub(remaining);

        emit Revoked(account, _incentives[account].id, remaining);

        return true;
    }

    function payout(address account, uint256 amount) external returns (uint256) {
        require(_incentives[account].isValue, "address was never incentivized");
        uint256 remaining = _incentives[account].amount - _incentives[account].provided;
        require(amount < remaining, "not enough token left");
        require(_token.transfer(address(this), amount), "token transfer failed");

        _incentives[account].provided += amount;

        emit Payout(account, _incentives[account].id, amount);
    }

    // Views

    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    function tokenAddress() public view returns (address) {
        return _token;
    }

    function incentivesLeft() public view returns (uint256) {
        return _incentivesLeft;
    }

    function incentivesProvided() public view returns (uint256){
        return _incentivesProvided;
    }

    function isInitialized() public view returns (bool) {
        return _isInitialized;
    }

    function incentives() public view returns (address[]) {
        return _incentivesList;
    }

    function getIncentive(address account)
    public view returns (
        uint256 amount,
        uint256 provided,
        bytes32 id,
        bool isValue
    ) {
        Incentive memory i = _incentives[account];

        return (i.amount, i.provided, i.id, i.isValue);
    }

    // pure

    function name() public pure returns (string) {
        return NAME;
    }

    function InitialSupply() public pure returns (uint256) {
        return INITIAL_SUPPLY;
    }
}