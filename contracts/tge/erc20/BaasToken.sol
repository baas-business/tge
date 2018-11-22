pragma solidity ^0.4.24;

import "./ERC20.sol";
import "../../ownership/Ownable.sol";

// @title BaasToken
contract BaasToken is ERC20, Ownable {
    using SafeMath for uint256;

    uint256 constant initialEscrowSupply = 60 * 10 ** 18;               // 60m Escrow Token
    uint256 constant initialPrivatePlacementSupply = 20 * 10 ** 18;     // 20m Private Placement Token
    uint256 constant initialFounderSupply = 10 * 10 ** 18;              // 10m Founder Token
    uint256 constant initialIncentiveSupply = 10 * 10 ** 18;            // 10m Incentive Token

    // prohibits usage of transfer and transferFrom
    bool isInitialized = false;

    address escrowContractAddress = address(0);
    address privatePlacementContractAddress = address(0);
    address founderContractAddress = address(0);
    address incentiveContractAddress = address(0);


    // setup(...) sets the address of deployed contracts that govern
    // the incentive, private placement, founder and escrow pots
    function setup(
        address _escrowContractAddress,
        address _privatePlacementContractAddress,
        address _founderContractAddress,
        address _incentiveContractAddress)
    public onlyOwner returns (bool) {
        require(!isInitialized);

        // delete existing balances to reset the supply
        _burn(escrowContractAddress, initialEscrowSupply);
        _burn(privatePlacementContractAddress, initialPrivatePlacementSupply);
        _burn(founderContractAddress, initialFounderSupply);
        _burn(incentiveContractAddress, initialIncentiveSupply);

        // mints the tokens for the pots according to their defined amounts
        _mint(escrowContractAddress, initialEscrowSupply);
        _mint(privatePlacementContractAddress, initialPrivatePlacementSupply);
        _mint(founderContractAddress, initialFounderSupply);
        _mint(incentiveContractAddress, initialIncentiveSupply);

        // stores the current address of the pot contracts to be able to reset the balances if contract is setup again
        escrowContractAddress = _escrowContractAddress;
        privatePlacementContractAddress = _privatePlacementContractAddress;
        founderContractAddress = _founderContractAddress;
        incentiveContractAddress = _incentiveContractAddress;

        return true;
    }

    // initialize() sets isInitialized flag to true
    //
    // once trading has started the contract can
    // 1. not be setup again
    // 2. the owner looses control over the contract
    // 3. token owner can transfer the tokens as they wish
    // 4. the escrow, incentive, private placement and founder token can distribute according to their logic
    function initialize() public onlyOwner returns (bool){
        isInitialized = true;
        return true;
    }

    function transferFrom(address from, address to, uint256 value)
    public returns (bool){
        require(isInitialized);
        return super.transferFrom(from, to, value);
    }

    function transfer(address to, uint256 value) public returns (bool) {
        require(isInitialized);
        return super.transfer(to, value);
    }
}
