pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";

interface IBaasIncentives {
    event IncentiveIssued(address indexed account, bytes32 indexed id, uint256 amount);
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
        address account;
        uint256 amount;
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

    function issue(address account, uint256 amount, bytes32 id) external onlyOwner returns (bool) {
        require(_isInitialized, "contract must be initialized");
        require(amount < _incentivesLeft, "not enough token left");
        require(!_incentives[account].isValue, "address should not been incentivized yet");

        // send token
        require(_token.transfer(address(this), amount), "token transfer failed");

        // update incentive storage
        _incentives[account] = Incentive(account, amount, id, true);
        _incentivesList.push(account);

        // update balance
        _incentivesLeft = _incentivesLeft.sub(amount);
        _incentivesProvided = _incentivesProvided.add(amount);

        emit IncentiveIssued(account, id, amount);

        return true;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  VIEWS
    ///////////////////////////////////////////////////////////////////////////////////////////

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
        bytes32 id,
        bool isValue
    ) {
        Incentive memory i = _incentives[account];

        return (i.amount, i.id, i.isValue);
    }


    ///////////////////////////////////////////////////////////////////////////////////////////
    //  PURE
    ///////////////////////////////////////////////////////////////////////////////////////////

    function name() public pure returns (string) {
        return NAME;
    }

    function InitialSupply() public pure returns (uint256) {
        return INITIAL_SUPPLY;
    }
}