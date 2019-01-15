pragma solidity 0.5;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


/**
 * @title BaasIncentives
 * @dev BaasIncentives is an 'ownable' contract that manages the tokens for incentivized participants.
 * A total of 10m tokens is deposited in this contract. Baas manually delivers tokens.
 */
contract BaasIncentives is Ownable {
    /**
    * @dev library for safe math operations
    */
    using SafeMath for uint256;

    /**
    * @dev initial supply for incentives
    */
    uint256 private constant INITIAL_SUPPLY = 10 * 10 ** 24;


    /**
    * @dev tracks the issued and left over token amounts
    */
    uint256 private _incentivesLeft;
    uint256 private _incentivesIssued;


    /**
    * @dev token contract address of BaaSToken
    */
    IBaasToken private _token;


    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Events
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
    * @dev emitted when tokens were issued to private investor
    */
    event TokensIssued(address indexed account, uint256 amount);

    /**
     * @dev constructor
     * @param token IBaasToken The address of the token smart contract
     */
    constructor(IBaasToken token) public {
        _token = token;
        _incentivesLeft = INITIAL_SUPPLY;
        _incentivesIssued = 0;
    }

    /**
     * @dev issues a specific amount of tokens to account by owner of contract
     * @param account address The address to send tokens to
     * @param amount uint256 The amount of token to be sent
     */
    function issue(address account, uint256 amount) external onlyOwner returns (bool) {
        require(amount < _incentivesLeft, "not enough token left");
        require(_token.transfer(account, amount), "token transfer failed");

        // update balance
        _incentivesLeft = _incentivesLeft.sub(amount);
        _incentivesIssued = _incentivesIssued.add(amount);

        emit TokensIssued(account, amount);

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
     * @dev shows incentives left in this contract
     * @return uint256 the amount of token left for incentives
     */
    function incentivesLeft() public view returns (uint256) {
        return _incentivesLeft;
    }

    /**
     * @dev shows incentives issued in this contract
     * @return uint256 the amount of token already issued
     */
    function incentivesIssued() public view returns (uint256){
        return _incentivesIssued;
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