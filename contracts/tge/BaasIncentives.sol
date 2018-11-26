pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IERC20.sol";

interface IBaasIncentives {
    function provideIncentive(address account, uint256 amount) external returns (bool);

    function withdraw() external returns (bool);

    event IncentiveProvided(address indexed account, uint256 amount);

    event Withdrawn(address indexed account, uint256 amount, uint stage);
}

contract BaasIncentives is IBaasIncentives, Ownable {
    using SafeMath for uint256;

    string private constant NAME = "INCENTIVES  ";
    uint constant STAGES = 4;

    struct Incentive {
        address account;        // the beneficiary of this incentive
        uint256 initialAmount;  // the incentive amount
        uint256 withdrawn;      // the amount of tokens already withdrawn
        uint atBlock;           // the time in blocks this incentive was initiated
        uint stage;            // which stage was withdrawn already
        bool isValue;
    }

    mapping(address => Incentive) _incentives;
    uint256 private _incentivesLeft;
    uint256 private _incentivesProvided;
    uint private _vestingPeriodInBlocks;
    bool private _isInitialized = false;

    IERC20 private _token;

    constructor(IERC20 token) public {
        _token = token;
    }

    function setup(uint vestingPeriodInBlocks, uint256 supply) external {
        require(!_isInitialized);
        _incentivesLeft = supply;
        _incentivesProvided = 0;
        _vestingPeriodInBlocks = vestingPeriodInBlocks;
    }

    function provideIncentive(address account, uint256 amount) external onlyOwner returns (bool) {
        require(amount <= _incentivesLeft);
        require(!_incentives[account].isValue);

        _incentivesLeft = _incentivesLeft.sub(amount);
        _incentivesProvided = _incentivesProvided.add(amount);

        _incentives[account] = Incentive(account, amount, 0, block.number, 0, true);

        emit IncentiveProvided(account, amount);

        return true;
    }

    function withdraw() external returns (bool) {
        require(_incentives[msg.sender].isValue);

        uint256 chunks = _incentives[msg.sender].initialAmount / STAGES;

        uint stagesCurrentlyAllowed = (block.number - _incentives[msg.sender].atBlock) / _vestingPeriodInBlocks;
        uint stagesToBePaidOut = stagesCurrentlyAllowed - _incentives[msg.sender].stage + 1;
        uint256 amount = chunks.mul(stagesToBePaidOut);

        require(stagesToBePaidOut > 0);
        require(_token.transfer(msg.sender, amount));

        _incentives[msg.sender].stage = stagesToBePaidOut;
        _incentives[msg.sender].withdrawn = _incentives[msg.sender].withdrawn.add(amount);

        emit Withdrawn(msg.sender, amount, stagesToBePaidOut);
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

    function vestingPeriodInBlocks() public view returns (uint) {
        return _vestingPeriodInBlocks;
    }

    function isInitialized() public view returns (bool) {
        return _isInitialized;
    }

    function getIncentive(address account)
    public view returns (uint256 initialAmount, uint256 withdrawn, uint atBlock, uint stage, bool isValue) {
        Incentive memory i = _incentives[account];

        return (i.initialAmount, i.withdrawn, i.atBlock, i.stage, i.isValue);
    }

    /*
        canWithdraw returns the amount, block number at which account can withdraw this stage
        and a flag if it already happened.
    */
    function canWithdraw(address account, uint stage) public view returns (uint256 amount, uint atBlock, bool alreadyWithdrawn) {
        if (stage == 0) {
            return (0, 0, false);
        }

        Incentive memory i = _incentives[account];

        amount = i.initialAmount / STAGES;
        atBlock = i.atBlock + (stage - 1) * _vestingPeriodInBlocks;
        alreadyWithdrawn = i.stage >= stage;

        return (amount, atBlock, alreadyWithdrawn);
    }

    function name() public pure returns (string) {
        return NAME;
    }
}