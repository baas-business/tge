pragma solidity 0.5;

import "./ERC20.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


/**
 * @title BaasEscrow
 * @dev BaasEscrow is an 'ownable' contract that manages the tokens held in escrow.
 * Baas can raise capital by freeing tokens to be issued to investors. Only one capital
 * raise at a time is allowed.
 */
contract BaasToken is IBaasToken, ERC20, Ownable {
    /**
     * @dev library for safe math operations
     */
    using SafeMath for uint256;

    /**
     * @dev the initial supply of Escrow, Private Placement, Founder and Incentives contracts
     */
    uint256 constant ESCROW_SUPPLY = 60 * 10 ** 24;                 // 60m Escrow Token
    uint256 constant PP_SUPPLY = 20 * 10 ** 24;                     // 20m Private Placement Token
    uint256 constant FOUNDER_SUPPLY = 10 * 10 ** 24;                // 10m Founder Token
    uint256 constant INCENTIVES_SUPPLY = 10 * 10 ** 24;             // 10m Incentives Token

    /**
     * @dev ERC20 constants
     */
    string public constant name = "BaaS Token";
    string public constant symbol = "BaaS";
    uint public constant decimals = 18;

    /**
     * @dev stores the addresses of the other contracts
     */
    address private _escrowAddress = address(0);
    address private _ppAddress = address(0);
    address private _founderAddress = address(0);
    address private _incentivesAddress = address(0);

    /**
     * @dev if paused transferring tokens is disabled
     */
    bool private _paused;

    /**
     * @dev flags if setup(...) was called once
     */
    bool private _isInitialized;

    /**
     * @dev tracks a token holder
     */
    struct TokenHolder {
        address account;
        uint listPointer;
    }

    /**
     * @dev map and list of current token holder
     * goal is to be able to return a whole list of token holders.
     * This combined data type is needed as neither a list nor a map alone can efficiently track struct items.
     * map => to keep track of the position of the item in the list
     * list => to be able to return a complete list
     */
    mapping(address => TokenHolder) private _tokenHolders;
    address[] private _tokenHolderList;

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Events
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev emitted when owner paused trading
     */
    event Paused(bool paused);

    /**
     * @dev emitted when setup is completed
     */
    event SetupCompleted();

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  External/Public Functions
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev sets the address of deployed contracts that govern the incentive, private placement, founder and escrow pots
     */
    function setup(
        address escrowAddress,
        address ppAddress,
        address founderAddress,
        address incentivesAddress)
    external onlyOwner returns (bool) {
        require(!_isInitialized);

        _escrowAddress = escrowAddress;
        _ppAddress = ppAddress;
        _founderAddress = founderAddress;
        _incentivesAddress = incentivesAddress;

        _mint(_escrowAddress, ESCROW_SUPPLY);
        _mint(_ppAddress, PP_SUPPLY);
        _mint(_founderAddress, FOUNDER_SUPPLY);
        _mint(_incentivesAddress, INCENTIVES_SUPPLY);

        _updateTokenHolderList(_escrowAddress);
        _updateTokenHolderList(_ppAddress);
        _updateTokenHolderList(_founderAddress);
        _updateTokenHolderList(_incentivesAddress);

        _isInitialized = true;

        emit SetupCompleted();

        return true;
    }

    /**
     * @dev as in ERC20 with additional logic
     *
     * 1. transferFrom can only be called if this contract is initialized and unpaused
     * 2. cannot be called by pot contract
     * 3. the token holder list is updated
     *
     * @param from address The address spending the token
     * @param to address The address receiving the token
     * @param value uint256 The amount being send
     * @return bool True if successful
     */
    function transferFrom(address from, address to, uint256 value)
    public returns (bool){
        require(_isInitialized && !_paused, "either not initialized or paused");
        require(!isPotAddress(to), "cannot send tokens to pots");

        require(super.transferFrom(from, to, value), "failed to transfer tokens");

        _updateTokenHolderList(from);
        _updateTokenHolderList(to);

        return true;
    }


    /**
     * @dev as in ERC20 with additional logic
     *
     * 1. transferFrom can only be called if this contract is initialized and unpaused
     * 2. cannot be called by pot contract
     * 3. the token holder list is updated
     *
     * @param to address The address receiving the token
     * @param value uint256 The amount being send
     * @return bool True if successful
     */
    function transfer(address to, uint256 value) public returns (bool) {
        require(_isInitialized && !_paused, "either not initialized or paused");
        require(!isPotAddress(to), "cannot send tokens to pots");

        require(super.transfer(to, value), "failed to transfer tokens");

        _updateTokenHolderList(msg.sender);
        _updateTokenHolderList(to);

        return true;
    }

    /**
     * @dev pauses or unpauses token transfers
     *
     * @param pause bool True if paused
     */
    function setPause(bool pause) external onlyOwner {
        require(_paused != pause, "cannot pause/unpause if already paused/unpaused");
        _paused = pause;
        emit Paused(pause);
    }

    /**
     * @dev pots can burn their tokens none other
     *
     * @param potAddress address The address which balance shall be burned
     * @param amount uint256 The amount that shall be burned
     */
    function burnTokensFromPot(address potAddress, uint256 amount)
    external returns (bool) {
        require(_isInitialized, "needs to be initialized");
        require(msg.sender == potAddress, "only pots themselves can burn tokens");
        require(isPotAddress(potAddress), "only tokens from pot can be burned");
        require(amount <= balanceOf(potAddress), "amount must be smaller then tokens held by the pot");

        _burn(potAddress, amount);

        _updateTokenHolderList(potAddress);

        return true;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Internal Functions
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev adds token holder to list of token holder if not registered yet or
     * deletes token holder if balance turned 0.
     *
     * @param tokenHolderAddress address The address of token holder that should be added or deleted
     */
    function _updateTokenHolderList(address tokenHolderAddress) internal returns (bool success) {
        // delete if token holder is registered and owns 0 tokens
        if (balanceOf(tokenHolderAddress) == 0) {
            if (isTokenHolder(tokenHolderAddress)) {
                _deleteTokenHolder(tokenHolderAddress);
            }
        } else {
            // add if token holder is not registered yet
            if (!isTokenHolder(tokenHolderAddress)) {
                _tokenHolders[tokenHolderAddress].account = tokenHolderAddress;
                _tokenHolders[tokenHolderAddress].listPointer = _tokenHolderList.push(tokenHolderAddress) - 1;
            }
        }

        return true;
    }

    /**
     * @dev deletes token holder from map and list
     *
     * @param tokenHolderAddress address The address of token holder that should deleted
     */
    function _deleteTokenHolder(address tokenHolderAddress) internal returns (bool success) {
        // get element to remove
        uint rowToDelete = _tokenHolders[tokenHolderAddress].listPointer;

        // pick last element to move it to the
        address keyToMove = _tokenHolderList[_tokenHolderList.length - 1];

        // put last element into position of removed element
        _tokenHolderList[rowToDelete] = keyToMove;

        // update listPointer of moved element
        _tokenHolders[keyToMove].listPointer = rowToDelete;

        // update length of list
        _tokenHolderList.length--;

        return true;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  VIEWS
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev indicates if setup was executed
     *
     * @return bool True if setup was executed
     */
    function isInitialized() public view returns (bool) {
        return _isInitialized;
    }

    /**
     * @dev indicates if setup was executed
     *
     * @return bool True if setup was executed
     */
    function isPotAddress(address account) public view returns (bool){
        return account == _escrowAddress || account == _incentivesAddress || account == _founderAddress || account == _ppAddress;
    }

    /**
     * @dev shows the addresses of escrow, private placement, founder and incentives pot
     *
     * @return (address,address,address,address) The pot addresses
     */
    function pots() public view returns (address, address, address, address) {
        return (_escrowAddress, _ppAddress, _founderAddress, _incentivesAddress);
    }

    /**
     * @dev indicates if account owns token
     *
     * @param account address token holder address
     * @return bool True if account ownds tokens
     */
    function isTokenHolder(address account) public view returns (bool) {
        if (_tokenHolderList.length == 0) return false;
        return (_tokenHolderList[_tokenHolders[account].listPointer] == account);
    }

    /**
     * @dev Shows the number of token holder that own tokens
     *
     * @return uint number of token holders
     */
    function tokenHolderCount() public view returns (uint) {
        return _tokenHolderList.length;
    }

    /**
     * @dev Shows a list of all token holders that own tokens
     *
     * @return address[] Addresses of all token holders
     */
    function tokenHolderSnapShot() public view returns (address[] memory) {
        return _tokenHolderList;
    }

    /**
     * @dev indicates trading is paused
     *
     * @return bool True if trading is paused
     */
    function paused() public view returns (bool) {
        return _paused;
    }

    /**
     * @dev Shows the circulating supply which is  = total supply - sum of pot balances
     *
     * @return uint256 Total number of tokens that can be traded
     */
    function circulatingSupply() public view returns (uint256) {
        return totalSupply().sub(potSupply());
    }

    /**
     * @dev Shows the pot supply which is the sum of pot balances
     *
     * @return uint256 Sum of pot's balances
     */
    function potSupply() public view returns (uint256) {
        if (!_isInitialized) {
            return 0;
        }

        uint256 escrowBalance = balanceOf(_escrowAddress);
        uint256 ppBalance = balanceOf(_ppAddress);
        uint256 founderBalance = balanceOf(_founderAddress);
        uint256 incentivesBalance = balanceOf(_incentivesAddress);

        return escrowBalance.add(incentivesBalance).add(ppBalance).add(founderBalance);
    }
}
