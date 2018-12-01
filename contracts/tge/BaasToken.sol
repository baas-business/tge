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

    /*
        For snap shooting token holders
    */
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
    // 1. not be setup again
    // 2. the owner looses control over the contract
    // 3. token owner can transfer the tokens as they wish
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
        if (!_isInitialized) {
            return false;
        }

        bool result = super.transferFrom(from, to, value);

        if (result) {
            updateTokenHolderList(from);
            updateTokenHolderList(to);
        }

        return result;
    }

    function transfer(address to, uint256 value) public returns (bool) {
        if (!_isInitialized) {
            return false;
        }

        bool result = super.transfer(to, value);

        if (result) {
            updateTokenHolderList(msg.sender);
            updateTokenHolderList(to);
        }

        return result;
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
        // delete if token holder registered
        if (balanceOf(tokenHolderAddress) == 0) {
            if (isTokenHolder(tokenHolderAddress)) {
                deleteTokenHolder(tokenHolderAddress);
            }

        } else {
            // add if token holder not registered yet
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

    /*
        Views
    */
    function isInitialized() public view returns (bool) {
        return _isInitialized;
    }

    function pots() public view returns (address, address, address, address) {
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
}
