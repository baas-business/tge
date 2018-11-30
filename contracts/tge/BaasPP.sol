pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IERC20.sol";


/*
    BaasPP holds tokens for private placement. A total of 20m tokens is deposited
    in this contract. Baas manually delivers tokens to holders. 2 conversion rates
    exist. 2.5m Tokens are delivered at conversion rate 0.5 and 17.5m tokens at
    conversion rate 1.0.
*/
interface IBaasPP {
    function provideToken(address account, uint256 amount, uint8 conversionRate) external returns (bool);

    function burnRest() external returns (bool);

    event TokenDelivered(address indexed to, uint256 amount, uint8 discountType);
}

/*
    Contract BaasPP

    BaasPP delivers token to private placement investors.
    2.5m tokens are provided with a discount of 50%.
    17.5m tokens are provided with no discount.
*/
contract BaasPP is IBaasPP, Ownable {
    using SafeMath for uint256;

    string private constant NAME = "PRIVATE PLACEMENT";

    uint8 private constant TOKEN_TYPE_TOTAL = 0;
    uint8 private constant DISCOUNT_TYPE_DISCOUNTED = 1;
    uint8 private constant DISCOUNT_TYPE_NOT_DISCOUNTED = 2;

    uint256 private constant HARD_CAP_DISCOUNTED_TOKEN = 25 * 10 ** 24;
    uint256 private constant HARD_CAP_NOT_DISCOUNTED_TOKEN = 175 * 10 ** 24;

    IERC20 private _token;

    constructor(IERC20 token)  public{
        _token = token;
    }

    uint256 private providedDiscountedToken;
    uint256 private providedNotDiscountedToken;

    /*
        deliverToken sends amount number of tokens to account if
        enough tokens are in the balance of this contract.


    */
    function provideToken(address account, uint256 amount, uint8 discountType)
    external onlyOwner returns (bool) {
        require(discountType == DISCOUNT_TYPE_DISCOUNTED || discountType == DISCOUNT_TYPE_NOT_DISCOUNTED);

        if (discountType == DISCOUNT_TYPE_DISCOUNTED) {
            provideDiscountedToken(account, amount);
        } else {
            provideNotDiscountedToken(account, amount);
        }

        emit TokenDelivered(account, amount, discountType);

        return true;
    }

    function provideDiscountedToken(address account, uint256 amount) internal {
        require(amount <= HARD_CAP_DISCOUNTED_TOKEN.sub(providedDiscountedToken));
        require(_token.transfer(account, amount));
        providedDiscountedToken = providedDiscountedToken.add(amount);
    }

    function provideNotDiscountedToken(address account, uint256 amount) internal {
        require(amount <= HARD_CAP_NOT_DISCOUNTED_TOKEN.sub(providedNotDiscountedToken));
        require(_token.transfer(account, amount));
        providedNotDiscountedToken = providedNotDiscountedToken.add(amount);
    }


    /*
    */
    function burnRest()
    external onlyOwner returns (bool) {
        // TODO
        return true;
    }

    // Views

    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    function tokenAddress() public view returns (address) {
        return _token;
    }

    function name() public pure returns (string) {
        return NAME;
    }

    function totalTokenProvided(uint8 discountType) public view returns (uint256) {
        if (discountType == TOKEN_TYPE_TOTAL) {
            return providedDiscountedToken.add(providedNotDiscountedToken);
        }

        if (discountType == DISCOUNT_TYPE_DISCOUNTED) {
            return providedDiscountedToken;
        }

        if (discountType == DISCOUNT_TYPE_NOT_DISCOUNTED) {
            return providedNotDiscountedToken;
        }

        return 0;
    }
}