pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";


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

    uint256 private constant HARD_CAP_DISCOUNTED_TOKEN = 50 * 10 ** 23;
    uint256 private constant HARD_CAP_NOT_DISCOUNTED_TOKEN = 150 * 10 ** 23;

    IBaasToken private _token;

    constructor(IBaasToken token)  public{
        _token = token;
    }

    uint256 private _providedDiscountedToken;
    uint256 private _providedNotDiscountedToken;

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
        require(amount <= HARD_CAP_DISCOUNTED_TOKEN.sub(_providedDiscountedToken));
        require(_token.transfer(account, amount));
        _providedDiscountedToken = _providedDiscountedToken.add(amount);
    }

    function provideNotDiscountedToken(address account, uint256 amount) internal {
        require(amount <= HARD_CAP_NOT_DISCOUNTED_TOKEN.sub(_providedNotDiscountedToken));
        require(_token.transfer(account, amount));
        _providedNotDiscountedToken = _providedNotDiscountedToken.add(amount);
    }

    function burnRest() external onlyOwner returns (bool) {
        require(_token.burnTokensFromPot(address(this), balance));
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
            return _providedDiscountedToken.add(_providedNotDiscountedToken);
        }

        if (discountType == DISCOUNT_TYPE_DISCOUNTED) {
            return _providedDiscountedToken;
        }

        if (discountType == DISCOUNT_TYPE_NOT_DISCOUNTED) {
            return _providedNotDiscountedToken;
        }

        return 0;
    }

    function Cap(uint discountType) public pure returns (uint256) {
        if (discountType == TOKEN_TYPE_TOTAL) {
            return HARD_CAP_NOT_DISCOUNTED_TOKEN.add(HARD_CAP_DISCOUNTED_TOKEN);
        }

        if (discountType == DISCOUNT_TYPE_DISCOUNTED) {
            return HARD_CAP_DISCOUNTED_TOKEN;
        }

        if (discountType == DISCOUNT_TYPE_NOT_DISCOUNTED) {
            return HARD_CAP_NOT_DISCOUNTED_TOKEN;
        }

        return 0;
    }
}