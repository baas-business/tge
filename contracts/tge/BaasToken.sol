pragma solidity ^0.4.24;

import "./ERC20.sol";
import "../ownership/Ownable.sol";


interface IBaasToken {
    event SetupCompleted();
}

contract BaasToken is IBaasToken, ERC20, Ownable {
    using SafeMath for uint256;

    string private constant NAME = "TOKEN";

    uint256 constant ESCROW_SUPPLY = 60 * 10 ** 18;                 // 60m Escrow Token
    uint256 constant PP_SUPPLY = 20 * 10 ** 18;                     // 20m Private Placement Token
    uint256 constant FOUNDER_SUPPLY = 10 * 10 ** 18;                // 10m Founder Token
    uint256 constant INCENTIVES_SUPPLY = 10 * 10 ** 18;             // 10m Incentives Token

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
    public onlyOwner returns (bool) {
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

        _isInitialized = true;

        emit SetupCompleted();
        return true;
    }

    function transferFrom(address from, address to, uint256 value)
    public returns (bool){
        if (!_isInitialized) {
            return false;
        }

        return super.transferFrom(from, to, value);
    }

    function transfer(address to, uint256 value) public returns (bool) {
        if (!_isInitialized) {
            return false;
        }

        return super.transfer(to, value);
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
}
