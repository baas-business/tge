pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


interface IBaasEscrow {

    event CapitalRaised(uint indexed id, uint256 amount);

    event SetupCompleted(uint vestingStartBlock, uint vestingPeriod, uint vestingEndBlock);

    event TokenDelivered(address indexed to, uint256 amount, uint256 tokenPrice);
}


/*
    Question:
    Any limitation on time between raises?
    Any limitations on time limit to claim
*/
contract BaasEscrow is IBaasEscrow, Ownable {
    using SafeMath for uint256;

    string private constant NAME = "ESCROW";
    IBaasToken private _token;


    uint256 private _vestingStart;
    uint256 private _vestingPeriod;

    uint256 private raisedCapital;

    struct CapitalRaise {
        uint id;
        uint256 amount;
        bool isValue;
    }

    mapping(uint => CapitalRaise) private capitalRaises;
    mapping(uint => mapping(address => uint256)) private whiteListed;


    bool private _isInitialized = false;

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

    function raiseCapital(uint256 amount, uint id) external onlyOwner returns (bool) {
        require(canRaise(block.number));
        require(amount <= balance());
        require(!capitalRaises[id].isValue);

        capitalRaises[id] = CapitalRaise(id, amount, true);


        address escrowAddress;
        address ppAddress;
        address founderAddress;
        address incentivesAddress;

        (escrowAddress, ppAddress, founderAddress, incentivesAddress) = _token.pots();

        address[] memory tokenHolders = _token.tokenHolderSnapShot();
        uint arrayLength = tokenHolders.length;


        for (uint i = 0; i < arrayLength; i++) {
            address tokenHolder = tokenHolders[i];
            if (tokenHolder == escrowAddress || tokenHolder == incentivesAddress || tokenHolder == founderAddress || tokenHolder == ppAddress) {
            } else {
                whiteListed[id][tokenHolder] = _token.balanceOf(tokenHolder);
            }
        }

        emit CapitalRaised(id, amount);
        return true;
    }


    function provideToken(address account, uint256 amount, uint256 conversionRate)
    external onlyOwner returns (bool) {
        require(_token.transfer(account, amount));

        emit TokenDelivered(account, amount, conversionRate);

        return true;
    }

    // Views

    function isInitialized() public view returns (bool) {
        return _isInitialized;
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


    function vestingStartBlock() public view returns (uint) {
        return _vestingStart;
    }

    function vestingPeriod() public view returns (uint) {
        return _vestingPeriod;
    }

    function vestingEndBlock() public view returns (uint) {
        return _vestingStart.add(_vestingPeriod);
    }

    function canRaise(uint blocknumber) public view returns (bool) {
        return vestingEndBlock() <= blocknumber;
    }

}