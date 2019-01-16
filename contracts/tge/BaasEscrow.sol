pragma solidity 0.5;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";

/**
 * @title BaasEscrow
 * @dev BaasEscrow is an 'ownable' contract that manages the tokens held in escrow.
 * Baas can raise capital by freeing tokens to be issued to investors. Only one capital
 * raise at a time is allowed.
 */
contract BaasEscrow is Ownable {
    /**
    * @dev library for safe math operations
    */
    using SafeMath for uint256;

    /**
    * @dev initial supply of escrow
    */
    uint256 private constant INITIAL_SUPPLY = 60 * 10 ** 24;

    /**
    * @dev tracks a capital raise
    */
    struct CapitalRaise {
        uint id;
        uint256 amount;
        uint256 provided;
        uint atBlock;
        bool isFinalized;
        bool isValue;
    }

    /**
    * @dev map and list of all capital raises
    */
    mapping(uint => CapitalRaise) private _capitalRaises;
    uint[] private _capitalRaiseIds;

    /**
    * @dev total amount of token being raised in all capital raises
    */
    uint256 private _raisedCapitalTotal;

    /**
    * @dev total amount of token being issued in all capital raises
    */
    uint256 private _issuedTokenTotal;

    /**
    * @dev total amount of token being burned in all capital raises
    */
    uint256 private _burnedTokenTotal;

    /**
    * @dev the id of the current capital raise
    */
    uint private _currentCapitalRaiseId;

    /**
    * @dev token contract address of BaaSToken
    */
    IBaasToken private _token;

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Events
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
    * @dev emitted when capital is raised
    */
    event CapitalRaised(uint indexed id, uint256 amount);

    /**
    * @dev emitted when tokens were issued to investor
    */
    event TokensIssued(address indexed account, uint256 amount);

    /**
    * @dev emitted when capital raise is finalized
    */
    event Finalized(uint id, uint256 burnedAmount);

    /**
     * @dev constructor
     * @param token IBaasToken The address of the token smart contract
     */
    constructor(IBaasToken token) public {
        _token = token;
    }

    /**
     * @dev frees token to be issued in a capital raise
     * @param amount uint256 The amount of token to be raised
     * @param id uint the id of this capital raise
     */
    function raiseCapital(uint256 amount, uint id)
    external onlyOwner returns (bool) {
        require(id != 0, "capital raise id cannot be 0");
        require(amount <= INITIAL_SUPPLY.sub(_raisedCapitalTotal), "not enough token left for this capital raise");
        require(!_capitalRaises[id].isValue, "don't use the same capital raise id twice");
        require(!isCapitalRaiseOngoing(), "a capital raise is ongoing and not finalized yet");

        // store meta data in contract
        _capitalRaises[id] = CapitalRaise(id, amount, 0, block.number, false, true);
        _capitalRaiseIds.push(id);

        // update contract state
        _currentCapitalRaiseId = id;
        _raisedCapitalTotal = _raisedCapitalTotal.add(amount);

        emit CapitalRaised(id, amount);

        return true;
    }

    /**
     * @dev issues a specific amount of tokens to account by owner of contract
     * @param account address The address to send tokens to
     * @param amount uint256 The amount of token to be sent
     */
    function issue(address account, uint256 amount)
    external onlyOwner returns (bool) {
        require(isCapitalRaiseOngoing(), "no capital raise ongoing");
        CapitalRaise memory capitalRaise = _capitalRaises[_currentCapitalRaiseId];
        require(amount < capitalRaise.amount.sub(capitalRaise.provided), "amount should be less then what is left");
        require(_token.transfer(account, amount), "transfer failed");

        _capitalRaises[_currentCapitalRaiseId].provided += amount;
        _issuedTokenTotal = _issuedTokenTotal.add(amount);

        emit TokensIssued(account, amount);

        return true;
    }

    /**
    * @dev finalization of current capital raise by contract owner. This burns tokens that were not
    * issued to private investors and stops future issuance.
    */
    function finalize()
    external onlyOwner returns (bool) {
        require(isCapitalRaiseOngoing(), "no capital raise ongoing");

        CapitalRaise memory capitalRaise = _capitalRaises[_currentCapitalRaiseId];

        uint256 tokensToBeBurned = capitalRaise.amount.sub(capitalRaise.provided);
        uint currentCapitalRaiseId = _currentCapitalRaiseId;

        require(_token.burnTokensFromPot(address(this), tokensToBeBurned), "burning of tokens failed");

        _currentCapitalRaiseId = 0;

        emit Finalized(currentCapitalRaiseId, tokensToBeBurned);

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

    /**
     * @dev shows if a capital raise is ongoing
     * @return bool indicating if capital raise is ongoing
     */
    function isCapitalRaiseOngoing() public view returns (bool) {
        return _currentCapitalRaiseId != 0;
    }

    /**
     * @dev shows current capital raise id
     * @return uint current capital raise id
     */
    function currentCapitalRaiseId() public view returns (uint) {
        return _currentCapitalRaiseId;
    }

    /**
     * @dev shows the total amount of tokens raised
     * @return uint256 the total amount of tokens raised
     */
    function raisedCapitalTotal() public view returns (uint256) {
        return _raisedCapitalTotal;
    }

    /**
     * @dev shows the total amount of tokens issued
     * @return uint256 the total amount of tokens issued
     */
    function issuedTokenTotal() public view returns (uint256) {
        return _issuedTokenTotal;
    }

    /**
     * @dev shows the total amount of tokens burned
     * @return uint256 the total amount of tokens burned
     */
    function burnedTokenTotal() public view returns (uint256) {
        return _burnedTokenTotal;
    }

    /**
     * @dev shows the capital raise
     * @return address token address of BaaSToken
     */
    function capitalRaise(uint capitalRaiseId) public view returns (
        uint id,
        uint256 amount,
        uint256 provided,
        uint atBlock,
        bool isFinalized,
        bool isValue
    ) {
        CapitalRaise memory c = _capitalRaises[capitalRaiseId];
        return (c.id, c.amount, c.provided, c.atBlock, c.isFinalized, c.isValue);
    }

    /**
     * @dev shows a list of all capital ids
     * @return address token address of BaaSToken
     */
    function capitalRaiseIds() public view returns (uint[] memory) {
        return _capitalRaiseIds;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Pure
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev shows the initial supply of this contract
     * @return uint256 intitial supply
     */
    function initialSupply() public pure returns (uint256) {
        return INITIAL_SUPPLY;
    }
}