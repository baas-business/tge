pragma solidity 0.5;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


/**
 * @title BaasPP
 * @dev BaasPP is an 'ownable' contract that manages the tokens for private placement.
 * A total of 20m tokens is deposited in this contract. Baas manually delivers tokens to holders.
 * 5m tokens are provided with a discount of 50%. 15m tokens are provided with no discount.
 */
contract BaasPP is Ownable {
    /**
    * @dev library for safe math operations
    */
    using SafeMath for uint256;

    /**
    * @dev constants to distinguish between discounted and not discounted investments
    */
    uint8 private constant DISCOUNT_TYPE_DISCOUNTED_AND_NOT_DISCOUNTED = 0;
    uint8 private constant DISCOUNT_TYPE_DISCOUNTED = 1;
    uint8 private constant DISCOUNT_TYPE_NOT_DISCOUNTED = 2;

    /**
    * @dev caps for discounted and not discounted tokens
    */
    uint256 private constant CAP_DISCOUNTED_TOKEN = 5 * 10 ** 24;
    uint256 private constant CAP_NOT_DISCOUNTED_TOKEN = 15 * 10 ** 24;

    /**
    * @dev token contract address of BaaSToken
    */
    IBaasToken private _token;

    /**
    * @dev tracker of issued token
    */
    uint256 private _issuedDiscountedToken;
    uint256 private _issuedNotDiscountedToken;

    /**
    * @dev flag indicating private placement is finished
    */
    bool private _isFinalized;

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Events
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
    * @dev emitted when a tokens were issued to private investor
    */
    event TokensIssued(address indexed account, uint8 indexed discountType, uint256 amount);

    /**
    * @dev emitted when contract is finalized
    */
    event Finalized(uint256 burnedAmount);

    /**
     * @dev constructor
     * @param token IBaasToken The address of the token smart contract
     */
    constructor(IBaasToken token)  public{
        _token = token;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  External/Public Functions
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev issues a specific amount of tokens to account by owner of contract
     * @param account address The address to send tokens to
     * @param amount uint256 The amount of token to be sent
     * @param discountType uint8 Either 1 (for discounted) or 2 (for not discounted)
     */
    function issue(address account, uint256 amount, uint8 discountType)
    external onlyOwner returns (bool) {
        require(!_isFinalized, "contract was already finalized");
        require(discountType == DISCOUNT_TYPE_DISCOUNTED || discountType == DISCOUNT_TYPE_NOT_DISCOUNTED, "bad discount type passed");

        if (discountType == DISCOUNT_TYPE_DISCOUNTED) {
            _issueDiscounted(account, amount);
        } else {
            _issueNotDiscountedToken(account, amount);
        }

        emit TokensIssued(account, discountType, amount);

        return true;
    }


    /**
     * @dev finalization of private placement by contract owner. This burns tokens that were not
     * issued to private investors and stops future issuance.
     */
    function finalize() external onlyOwner returns (bool) {
        require(!_isFinalized, "contract was already finalized");
        uint256 balance = balance();
        require(_token.burnTokensFromPot(address(this), balance), "burning of tokens failed");

        _isFinalized = true;

        emit Finalized(balance);

        return true;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Internal Functions
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev internal function to issue discounted tokens
     * @param account address The address to send tokens to
     * @param amount uint256 The amount of token to be sent
     */
    function _issueDiscounted(address account, uint256 amount) internal {
        require(amount <= CAP_DISCOUNTED_TOKEN.sub(_issuedDiscountedToken), "amount exceeds hard cap of discounted token");
        require(_token.transfer(account, amount), "token transfer failed");
        _issuedDiscountedToken = _issuedDiscountedToken.add(amount);
    }

    /**
     * @dev internal function to issue not discounted tokens
     * @param account address The address to send tokens to
     * @param amount uint256 The amount of token to be sent
     */
    function _issueNotDiscountedToken(address account, uint256 amount) internal {
        require(amount <= CAP_NOT_DISCOUNTED_TOKEN.sub(_issuedNotDiscountedToken), "amount exceeds hard cap of not discounted token");
        require(_token.transfer(account, amount), "token transfer failed");
        _issuedNotDiscountedToken = _issuedNotDiscountedToken.add(amount);
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Views
    ///////////////////////////////////////////////////////////////////////////////////////////
    /**
     * @dev shows state of this contract
     * @return bool flag if contract was already finalized.
     */
    function isFinalized() public view returns (bool) {
        return _isFinalized;
    }

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
     * @dev shows amount of tokens issued for discount type
     * @param discountType uint8 either 0, 1 or 2 for total, discounted and not discounted
     * @return uint256 amount of tokens
     */
    function tokensIssued(uint8 discountType) public view returns (uint256) {
        if (discountType == DISCOUNT_TYPE_DISCOUNTED_AND_NOT_DISCOUNTED) {
            return _issuedDiscountedToken.add(_issuedNotDiscountedToken);
        }

        if (discountType == DISCOUNT_TYPE_DISCOUNTED) {
            return _issuedDiscountedToken;
        }

        if (discountType == DISCOUNT_TYPE_NOT_DISCOUNTED) {
            return _issuedNotDiscountedToken;
        }

        return 0;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Pure
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev shows cap for discount type
     * @param discountType uint8 either 0, 1 or 2 for total, discounted and not discounted
     * @return uint256 cap
     */
    function cap(uint discountType) public pure returns (uint256) {
        if (discountType == DISCOUNT_TYPE_DISCOUNTED_AND_NOT_DISCOUNTED) {
            return CAP_NOT_DISCOUNTED_TOKEN.add(CAP_DISCOUNTED_TOKEN);
        }

        if (discountType == DISCOUNT_TYPE_DISCOUNTED) {
            return CAP_DISCOUNTED_TOKEN;
        }

        if (discountType == DISCOUNT_TYPE_NOT_DISCOUNTED) {
            return CAP_NOT_DISCOUNTED_TOKEN;
        }

        return 0;
    }
}