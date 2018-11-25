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
    function deliverToken(address account, uint256 amount, uint8 conversionRate) external returns (bool);

    function tokenAddress() public view returns (address);
    function balance() public view returns (uint256);

    event TokenDelivered(address indexed to, uint256 amount, uint8 conversionRate);
}

/*
*/
contract BaasPP is IBaasPP, Ownable {
    using SafeMath for uint256;

    IERC20 private _token;

    constructor(IERC20 token){
        _token = token;
    }

    /*
        deliverToken sends amount number of tokens to account if
        enough tokens are in the balance of this contract.
    */
    function deliverToken(address account, uint256 amount, uint8 conversionRate)
    external onlyOwner returns (bool) {
        require(amount <= _token.balanceOf(address(this)));
        require(_token.transfer(account, amount));

        emit TokenDelivered(account, amount, conversionRate);
        return true;
    }

    // Views

    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    function tokenAddress() public view returns (address) {
        return _token;
    }
}