pragma solidity 0.5;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


contract BaasEscrow is Ownable {
    using SafeMath for uint256;

    IBaasToken private _token;


    uint256 private constant INITIAL_SUPPLY = 60 * 10 ** 24;

    uint256 private _vestingStart;
    uint256 private _vestingPeriod;

    uint256 private raisedCapital;

    struct CapitalRaise {
        uint id;
        uint256 amount;
        uint256 provided;
        uint atBlock;
        bool isValue;
    }

    mapping(uint => CapitalRaise) private capitalRaises;
    mapping(uint => mapping(address => uint256)) private whiteListed;
    mapping(uint => address[]) private whiteListedAddresses;


    bool private _isInitialized = false;

    event CapitalRaised(uint indexed id, uint256 amount);

    event SetupCompleted(uint vestingStartBlock, uint vestingPeriod, uint vestingEndBlock);

    event TokenDelivered(address indexed to, uint256 amount, uint256 tokenPrice);

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
        require(amount <= INITIAL_SUPPLY.sub(raisedCapital));
        require(!capitalRaises[id].isValue);

        capitalRaises[id] = CapitalRaise(id, amount,0, block.number, true);

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
                whiteListedAddresses[id].push(tokenHolder);
            }
        }

        emit CapitalRaised(id, amount);
        return true;
    }


    function provideToken(address account, uint256 amount, uint256 conversionRate, uint capitalRaiseId)
    external onlyOwner returns (bool) {
        CapitalRaise memory capitalRaise = capitalRaises[capitalRaiseId];
        require(capitalRaise.isValue, "wrong capital raise id");
        require(amount < capitalRaise.amount.sub(capitalRaise.provided), "amount should be less then what is left");
        require(_token.transfer(account, amount), "transfer failed");

        capitalRaises[capitalRaiseId].provided += amount;

        emit TokenDelivered(account, amount, conversionRate);

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

    function isInitialized() public view returns (bool) {
        return _isInitialized;
    }

    function capitalRaise(uint capitalRaiseId) public view returns
    (
        uint id,
        uint256 amount,
        uint256 provided,
        uint atBlock,
        bool isValue
    ) {
        CapitalRaise memory c = capitalRaises[capitalRaiseId];
        return (c.id, c.amount, c.provided, c.atBlock, c.isValue);
    }


    function whiteList(uint capitalRaiseId) public view returns (address[] memory) {
        return whiteListedAddresses[capitalRaiseId];
    }

    function whiteListedAccount(uint capitalRaiseId, address account) public view returns (uint256) {
        return whiteListed[capitalRaiseId][account];
    }
}