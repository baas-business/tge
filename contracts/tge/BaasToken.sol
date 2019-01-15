pragma solidity ^0.4.24;

import "./ERC20.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


contract BaasToken is IBaasToken, ERC20, Ownable {
    using SafeMath for uint256;

    string private constant NAME = "TOKEN";

    uint256 constant ESCROW_SUPPLY = 60 * 10 ** 24;                 // 60m Escrow Token
    uint256 constant PP_SUPPLY = 20 * 10 ** 24;                     // 20m Private Placement Token
    uint256 constant FOUNDER_SUPPLY = 10 * 10 ** 24;                // 10m Founder Token
    uint256 constant INCENTIVES_SUPPLY = 10 * 10 ** 24;             // 10m Incentives Token

    bool private _paused;

    struct TokenHolder {
        address wallet;
        uint listPointer;
    }

    mapping(address => TokenHolder) private _tokenHolders;
    address[] private _tokenHolderList;

    // prohibits usage of transfer and transferFrom
    bool private _isInitialized = false;

    address _escrowAddress = address(0);
    address _ppAddress = address(0);
    address _founderAddress = address(0);
    address _incentivesAddress = address(0);

    // setup(...) sets the address of deployed contracts that govern
    // the incentive, private placement, founder and escrow pots
    //
    // once trading has started the contract can
    //
    // 1. not be setup again
    // 2. the owner looses control over the contract
    // 3. token owner can transfer tokens
    // 4. the escrow, incentive, private placement and founder token can distribute according to their logic
    function setup(
        address escrowAddress,
        address ppAddress,
        address founderAddress,
        address incentivesAddress)
    external onlyOwner returns (bool) {
        require(!_isInitialized);

        // stores the current address of the pot contracts to be able to reset the balances if contract is setup again
        _escrowAddress = escrowAddress;
        _ppAddress = ppAddress;
        _founderAddress = founderAddress;
        _incentivesAddress = incentivesAddress;

        // mints the tokens for the pots according to their defined amounts
        _mint(_escrowAddress, ESCROW_SUPPLY);
        _mint(_ppAddress, PP_SUPPLY);
        _mint(_founderAddress, FOUNDER_SUPPLY);
        _mint(_incentivesAddress, INCENTIVES_SUPPLY);

        updateTokenHolderList(_escrowAddress);
        updateTokenHolderList(_ppAddress);
        updateTokenHolderList(_founderAddress);
        updateTokenHolderList(_incentivesAddress);

        _isInitialized = true;

        emit SetupCompleted();
        return true;
    }

    function transferFrom(address from, address to, uint256 value)
    public returns (bool){
        require(_isInitialized && !_paused);
        require(!isPotAddress(to), "cannot send tokens to pots");

        require(super.transferFrom(from, to, value), "failed to transfer tokens");

        updateTokenHolderList(from);
        updateTokenHolderList(to);

        return true;
    }

    function transfer(address to, uint256 value) public returns (bool) {
        require(_isInitialized && !_paused);
        require(!isPotAddress(to), "cannot send tokens to pots");

        require(super.transfer(to, value), "failed to transfer tokens");

        updateTokenHolderList(msg.sender);
        updateTokenHolderList(to);

        return true;
    }

    function burnTokensFromPot(address potAddress, uint256 amount)
    external returns (bool) {
        require(_isInitialized);
        // only pots themselves and owner can burn tokens
        require(msg.sender == potAddress || msg.sender == owner());
        // only tokens from pot can be burned
        require(potAddress == _escrowAddress || potAddress == _incentivesAddress || potAddress == _founderAddress || potAddress == _ppAddress);
        // amount must be smaller then tokens held by the pot
        require(amount <= balanceOf(potAddress));

        _burn(potAddress, amount);

        return true;
    }

    function updateTokenHolderList(address tokenHolderAddress) internal returns (bool success) {
        // delete if token holder is registered
        if (balanceOf(tokenHolderAddress) == 0) {
            if (isTokenHolder(tokenHolderAddress)) {
                deleteTokenHolder(tokenHolderAddress);
            }
        } else {
            // add if token holder is not registered yet
            if (!isTokenHolder(tokenHolderAddress)) {
                _tokenHolders[tokenHolderAddress].wallet = tokenHolderAddress;
                _tokenHolders[tokenHolderAddress].listPointer = _tokenHolderList.push(tokenHolderAddress) - 1;
            }
        }

        return true;
    }

    function deleteTokenHolder(address tokenHolderAddress) internal returns (bool success) {
        if (!isTokenHolder(tokenHolderAddress)) revert();

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

    function setPause(bool pause) external onlyOwner {
        require(_paused != pause);
        _paused = pause;
        emit Paused(pause);
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  VIEWS
    ///////////////////////////////////////////////////////////////////////////////////////////

    function isInitialized() public view returns (bool) {
        return _isInitialized;
    }

    function isPotAddress(address account) public view returns (bool){
        return account == _escrowAddress || account == _incentivesAddress || account == _founderAddress || account == _ppAddress;
    }

    function pots() external view returns (address, address, address, address) {
        return (_escrowAddress, _ppAddress, _founderAddress, _incentivesAddress);
    }

    function name() public pure returns (string) {
        return NAME;
    }

    function isTokenHolder(address tokenHolderAddress) public view returns (bool) {
        if (_tokenHolderList.length == 0) return false;
        return (_tokenHolderList[_tokenHolders[tokenHolderAddress].listPointer] == tokenHolderAddress);
    }

    function tokenHolderCount() public view returns (uint) {
        return _tokenHolderList.length;
    }

    function tokenHolderSnapShot() public view returns (address[]) {
        return _tokenHolderList;
    }

    function paused() public view returns (bool) {
        return _paused;
    }

    function circulatingSupply() public view returns (uint256) {
        return totalSupply().sub(potSupply());
    }

    function potSupply() public view returns (uint256) {
        if (!_isInitialized) {
            return 0;
        }

        // get balances of pots
        uint256 escrowBalance = balanceOf(_escrowAddress);
        uint256 ppBalance = balanceOf(_ppAddress);
        //uint256 founderBalance = balanceOf(_founderAddress);
        uint256 incentivesBalance = balanceOf(_incentivesAddress);

        return escrowBalance.add(incentivesBalance).add(ppBalance);
        //.add(founderBalance);
    }
}
