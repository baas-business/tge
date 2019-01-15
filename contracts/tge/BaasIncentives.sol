pragma solidity 0.5;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";

interface IBaasIncentives {
    event IncentiveIssued(address indexed account, bytes32 indexed id, uint256 amount);
    event Payout(address indexed account, bytes32 indexed id, uint256 amount);
    event Revoked(address indexed account, bytes32 indexed id, uint256 tokenNotDelivered);

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

    /**
    * @dev token contract address of BaaSToken
    */
    IBaasToken private _token;

    /**
     * @dev constructor
     * @param token IBaasToken The address of the token smart contract
     */
    constructor(IBaasToken token) public {
        _token = token;
        _incentivesLeft = INITIAL_SUPPLY;
        _incentivesProvided = 0;
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
    //  Views
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev shows balance of this contract
     * @return uint256 the amount of token held by this contract.
     */
    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    /**
     * @dev shows the token address this contracts references
     * @return address token address of BaaSToken
     */
    function tokenAddress() public view returns (address) {
        return address(_token);
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

    function incentives() public view returns (address[] memory) {
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
    //  Pure
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev shows the name of this contract
     * @return string name of this contract
     */
    function name() public pure returns (string memory) {
        return NAME;
    }

    function InitialSupply() public pure returns (uint256) {
        return INITIAL_SUPPLY;
    }
}