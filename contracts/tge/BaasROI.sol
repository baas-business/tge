pragma solidity ^0.4.24;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";

interface IBaasROI {
    function payoutAll(uint256 tokenEuroConversionRate) external returns (bool);

    event PayoutOmitted(address receiver);
    event PaidOut(address receiver, uint256 tokensReceived, uint256 tokensPossessed);
    event PaidOutAll(uint256 tokensProvided, uint256 tokensPossessed, uint tokenHolders);
}

/*
    1 Should the conversion rate being constantly updated or on day of payout being transmitted as a parameter?
*/
contract BaasROI is IBaasROI, Ownable {
    using SafeMath for uint256;

    uint256 constant private INTEREST_RATE = 9;

    string private constant NAME = "RETURN OF INVESTMENT";

    IBaasToken private _token;

    constructor(IBaasToken token) public {
        _token = token;
    }

    function payoutAll(uint256 tokenEuroConversionRate) external onlyOwner returns (bool) {
        require(hasEnoughTokensForPayout(tokenEuroConversionRate));

        uint256 contractBalance = balance();

        address escrowAddress;
        address ppAddress;
        address founderAddress;
        address incentivesAddress;

        (escrowAddress, ppAddress, founderAddress, incentivesAddress) = _token.pots();

        address[] memory tokenHolders = _token.tokenHolderSnapShot();
        uint arrayLength = tokenHolders.length;

        uint256 totalAmount = 0;
        uint totalTokenHolder = 0;

        for (uint i = 0; i < arrayLength; i++) {
            address tokenHolder = tokenHolders[i];
            if (tokenHolder == escrowAddress || tokenHolder == incentivesAddress || tokenHolder == founderAddress || tokenHolder == ppAddress || tokenHolder == address(this)) {
                emit PayoutOmitted(tokenHolder);
            } else {
                uint256 amount = payout(tokenHolders[i], tokenEuroConversionRate);
                totalAmount = totalAmount.add(amount);
                totalTokenHolder = totalTokenHolder + 1;
            }
        }

        emit PaidOutAll(totalAmount, contractBalance, totalTokenHolder);
    }


    function payout(address tokenHolder, uint256 tokenEuroConversionRate)
    internal returns (uint256){
        // get balance and amount to pay out
        uint256 amount;
        uint256 balance;
        (amount, balance) = roiOf(tokenHolder, tokenEuroConversionRate);

        require(_token.transfer(tokenHolder, amount));

        emit PaidOut(tokenHolder, amount, balance);

        return amount;
    }

    // Views
    function balance() public view returns (uint256) {
        return _token.balanceOf(address(this));
    }

    function tokenAddress() public view returns (address) {
        return _token;
    }

    function circulatingSupply() public view returns (uint256) {
        return _token.circulatingSupply();
    }

    function eligibleToken() public view returns (uint256) {
        return _token.circulatingSupply().sub(balance());
    }

    function hasEnoughTokensForPayout(uint256 tokenEuroConversionRate) public view returns (bool) {
        return currentPayout(tokenEuroConversionRate) <= balance();
    }

    function tokensNeededForPayout(uint256 tokenEuroConversionRate) public view returns (uint256) {
        if (hasEnoughTokensForPayout(tokenEuroConversionRate)) {
            return 0;
        }

        return currentPayout(tokenEuroConversionRate) - balance();
    }

    function roi(uint256 token, uint256 tokenEuroConversionRate) public pure returns (uint256) {
        if (tokenEuroConversionRate == 0) {
            return 0;
        }

        return token.mul(INTEREST_RATE).div(tokenEuroConversionRate);
    }

    function roiOf(address wallet, uint256 tokenEuroConversionRate) public view returns (uint256, uint256) {
        uint256 walletBalance = _token.balanceOf(wallet);
        return (roi(walletBalance, tokenEuroConversionRate), walletBalance);
    }

    function currentPayout(uint256 tokenEuroConversionRate) public view returns (uint256) {
        return roi(eligibleToken(), tokenEuroConversionRate);
    }


    function maxTokensToBeRewarded(uint256 tokenEuroConversionRate) public view returns (uint256) {
        uint256 cri = tokenEuroConversionRate.add(INTEREST_RATE);
        return circulatingSupply().mul(tokenEuroConversionRate).div(cri);
    }

    function minPayout(uint256 tokenEuroConversionRate) public view returns (uint256) {
        uint256 cri = INTEREST_RATE.add(tokenEuroConversionRate);
        return circulatingSupply().mul(cri).div(tokenEuroConversionRate);
    }







    // Pure

    function interestRate() public pure returns (uint256) {
        return INTEREST_RATE;
    }

    function name() public pure returns (string) {
        return NAME;
    }
}