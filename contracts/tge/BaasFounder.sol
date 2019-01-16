pragma solidity 0.5;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


contract BaasFounder is Ownable {
    /**
    * @dev library for safe math operations
    */
    using SafeMath for uint256;

    /**
    * @dev initial founder supply. 8m token for founder 1 and 2m token for founder 2.
    */
    uint256 constant FOUNDER1_SUPPLY = 8 * 10 ** 24;                // 8m Founder1 Token
    uint256 constant FOUNDER2_SUPPLY = 2 * 10 ** 24;                // 2m Founder2 Token

    /**
    * @dev tracks if founder 1/2 have received their tokens
    */
    bool[] private _founderHasReceivedTokens;


    /**
    * @dev token contract address of BaaSToken
    */
    IBaasToken private _token;

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Events
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
    * @dev emitted when a tokens were issued to private investor
    */
    event TokensIssued(address indexed receiver, uint indexed founderId, uint256 amount);

    /**
     * @dev constructor
     * @param token IBaasToken The address of the token smart contract
     */
    constructor(IBaasToken token) public {
        _token = token;
        _founderHasReceivedTokens.push(false);
        _founderHasReceivedTokens.push(false);
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  External/Public Functions
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev issues tokens to founder.
     * @param receiver address wallet address of founder
     * @param founderId int the id of the founder
     */
    function issue(address receiver, uint founderId)
    external onlyOwner returns (bool) {
        require(founderId <= 1, "founderId not existent");
        require(!_founderHasReceivedTokens[founderId], "founder has already received tokens");

        uint256 amount = 0;

        if (founderId == 0) {
            amount = FOUNDER1_SUPPLY;
        } else {
            amount = FOUNDER2_SUPPLY;
        }

        require(amount <= balance(), "not enough tokens left");
        require(_token.transfer(receiver, amount), "transfer of tokens failed");

        _founderHasReceivedTokens[founderId] = true;

        emit TokensIssued(receiver, founderId, amount);

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
     * @dev shows if founder has received tokens.
     * @param  founderId int the id of the founder. Either 0 or 1.
     * @return bool flag indicating if founder received tokens
     */
    function hasFounderReceivedTokens(uint founderId) public view returns (bool){
        return _founderHasReceivedTokens[founderId];
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Pure
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev shows the founder supply of founder 1
     * @return uint256 founder supply
     */
    function founder1Supply() public pure returns (uint256) {
        return FOUNDER1_SUPPLY;
    }

    /**
     * @dev shows the founder supply of founder 2
     * @return uint256 founder supply
     */
    function founder2Supply() public pure returns (uint256) {
        return FOUNDER2_SUPPLY;
    }
}