pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";

interface IBaasIncentives {


    event IncentiveProvided1(address indexed account, uint256 amount);
    event IncentiveProvided2(address indexed account, uint vestingStages, uint vestingBlocks);

    event Claimed(address indexed account, uint256 amount, uint stage);
    event Forfeited(address indexed account, uint256 remainingToken);

    event SetupCompleted(uint256 supply);
}
/*
    Can claim first stage immediately?
*/
contract BaasIncentives is IBaasIncentives, Ownable {
    using SafeMath for uint256;
    using SafeMath for uint;

    string private constant NAME = "INCENTIVES";

    uint8 private constant STATE_CAN_CLAIM = 0;
    uint8 private constant STATE_CAN_NOT_CLAIM_YET = 1;
    uint8 private constant STATE_HAS_CLAIMED = 2;
    uint8 private constant STATE_STAGE_NOT_AVAILABLE = 3;
    uint8 private constant STATE_USER_NON_EXISTENT_OR_FORFEITED = 4;

    struct Incentive {
        address account;        // the beneficiary of this incentive
        uint256 amountPerStage;  // the incentive amount
        uint256 amountClaimed;      // the amount of tokens already withdrawn
        uint atBlock;           // the time in blocks this incentive was initiated
        uint totalVestingStages;             // which stage was withdrawn already
        uint blockTimePerStage;
        bool[] stagesClaimed;
        uint[] stagesBlockTime;
        uint listPointer;
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

    function reward(address account, uint256 amountPerStage, uint totalVestingStages, uint blockTimePerStage) external onlyOwner returns (bool) {
        require(_isInitialized);
        require(blockTimePerStage > 0);
        require(totalVestingStages > 0);
        require(!_incentives[account].isValue);
        require(amountPerStage <= _incentivesLeft);


        // update incentive storage
        _incentives[account] = Incentive(account, amountPerStage, 0, block.number, totalVestingStages, blockTimePerStage, new bool[](0), new uint[](0), _incentivesList.length, true);
        _incentivesList.push(account);

        for (uint i = 0; i < totalVestingStages; i++) {
            _incentives[account].stagesClaimed.push(false);

            uint blockTime = blockTimePerStage.mul(i) + block.number;
            _incentives[account].stagesBlockTime.push(blockTime);
        }


        // update contract balance
        _incentivesLeft = _incentivesLeft.sub(amountPerStage);
        _incentivesProvided = _incentivesProvided.add(amountPerStage);

        emit IncentiveProvided1(account, amountPerStage);
        emit IncentiveProvided2(account, totalVestingStages, blockTimePerStage);

        return true;
    }

    function forfeited(address account) external onlyOwner returns (bool) {
        require(_incentives[account].isValue);
        _incentives[account].isValue = false;

        uint256 remaining = _claim(account);
        _incentivesLeft = _incentivesLeft.add(remaining);
        _incentivesProvided = _incentivesProvided.sub(remaining);
        emit Forfeited(account, remaining);
        return true;
    }

    function claim() external returns (uint256) {
        return _claim(msg.sender);
    }

    function _claim(address account) internal returns (uint256 rest) {
        require(_isInitialized);

        Incentive memory incentive = _incentives[msg.sender];
        require(incentive.isValue);

        uint256 total = 0;
        for (uint i = 0; i < incentive.totalVestingStages; i++) {
            if (!incentive.stagesClaimed[i]) {
                if (incentive.stagesBlockTime[i] <= block.number) {
                    _incentives[msg.sender].amountClaimed = _incentives[msg.sender].amountClaimed.add(incentive.amountPerStage);
                    _incentives[msg.sender].stagesClaimed[i] = true;
                    total = total.add(incentive.amountPerStage);

                    emit Claimed(account, incentive.amountPerStage, i);
                } else {
                    break;
                }
            }
        }


        require(_token.transfer(account, total));

        return incentive.amountPerStage.mul(incentive.totalVestingStages).sub(incentive.amountClaimed);
    }

    function state(address account, uint stage, uint blockNumber) public view returns (uint8) {
        Incentive memory i = _incentives[account];

        if (!i.isValue) {
            return STATE_USER_NON_EXISTENT_OR_FORFEITED;
        }

        if (stage > i.stagesClaimed.length - 1) {
            return STATE_STAGE_NOT_AVAILABLE;
        }

        if (i.stagesClaimed[stage]) {
            return STATE_HAS_CLAIMED;
        }

        if (i.stagesBlockTime[stage] > blockNumber) {
            return STATE_CAN_NOT_CLAIM_YET;
        }

        return STATE_CAN_CLAIM;
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
        uint256 initialAmount,
        uint256 withdrawn,
        uint atBlock,
        uint totalStages,
        uint totalBlocks,
        bool[] stagesClaimed,
        uint[] stagesBlockTime,
        bool isValue
    ) {
        Incentive memory i = _incentives[account];

        return (i.amountPerStage, i.amountClaimed, i.atBlock, i.totalVestingStages, i.blockTimePerStage, i.stagesClaimed, i.stagesBlockTime, i.isValue);
    }

    // pure

    function name() public pure returns (string) {
        return NAME;
    }
}