pragma solidity 0.5;


import "../math/SafeMath.sol";
import "../ownership/Ownable.sol";
import "./IBaasToken.sol";



/**
 * @title BaasROI
 * @dev BaasROI is an 'ownable' contract that manages the pays out all eligible token holder a return of investment.
 * Baas manually triggers a payout to all token holders and tokens are transferred to token holders at given interest
 * rate.
 */
contract BaasROI is Ownable {
    using SafeMath for uint256;

    /**
    * @dev the interest rate paid for each token in percent
    */
    uint256 constant private INTEREST_RATE = 9;

    /**
    * @dev token contract address of BaaSToken
    */
    IBaasToken private _token;

    /**
     * @dev constructor
     * @param token IBaasToken The address of the token smart contract
     */
    constructor(IBaasToken token) public {
        _token = token;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Events
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
    * @dev emitted when a token holder was omitted during payout
    */
    event PayoutOmitted(address omittedReceiver);

    /**
    * @dev emitted when a token holder was paid out
    */
    event PaidOut(address indexed receiver, uint256 tokensReceived, uint256 tokensPossessed);

    /**
    * @dev emitted when all token holders were paid out
    */
    event PaidOutAll(uint256 tokensProvided, uint256 tokensPossessed, uint tokenHolders);

    /**
    * @dev emitted when tokens are withdrawn from this contract
    */
    event Withdrawn(address indexed receiver, uint256 amount);

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  External/Public Functions
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**

     * @dev pays out all token holders registered at BaaS Token a dividend based on the token euro conversion rate.
     * It omits the pots (Escrow, Founder, Incentives, Private Placement) and tokens held by this contract.
     * @param tokenEuroConversionRate address The token euro conversion rate
     * @return bool If Successful
     */
    function payoutAll(uint256 tokenEuroConversionRate)
    external onlyOwner returns (bool) {
        require(contractHasEnoughTokensForPayout(tokenEuroConversionRate), "contracts balance is too low");

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
                uint256 amount = _payout(tokenHolders[i], tokenEuroConversionRate);
                totalAmount = totalAmount.add(amount);
                totalTokenHolder = totalTokenHolder + 1;
            }
        }

        emit PaidOutAll(totalAmount, contractBalance, totalTokenHolder);
    }

    /**
     * @dev withdraws tokens from this contract if available
     * @param receiver address The receiver of the tokens
     * @return bool If Successful
     */
    function withdrawFromContract(address receiver) external onlyOwner returns (bool) {
        uint256 amount = balance();
        require(amount > 0, "no tokens left in contract");
        require(_token.transfer(receiver, amount), "token transfer failed");

        emit Withdrawn(receiver, amount);

        return true;
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Internal Functions
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev pays out one token holder on payout day
     * @param receiver address The receiver of the tokens
     * @param tokenEuroConversionRate uint256 The token euro conversion rate used for this payout
     * @return uint256 amount that was delivered
     */
    function _payout(address receiver, uint256 tokenEuroConversionRate)
    internal returns (uint256){
        // get balance and amount to pay out
        uint256 balance = _token.balanceOf(receiver);
        uint256 amount = roi(balance, tokenEuroConversionRate);

        require(_token.transfer(receiver, amount), "transfer of tokens failed");

        emit PaidOut(receiver, amount, balance);

        return amount;
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
     * @dev shows the circulating supply of Baas tokens
     * @return uint256 the amount of token circulating
     */
    function circulatingSupply() public view returns (uint256) {
        return _token.circulatingSupply();
    }

    /**
     * @dev shows the eligible tokens for payout
     * @return uint256 The amount of eligible tokens
     */
    function eligibleToken() public view returns (uint256) {
        return _token.circulatingSupply().sub(balance());
    }

    /**
     * @dev indicates if contract itself has enough tokens to pay out all token holders
     * @return bool True if enough tokens are deposited in this contract
     */
    function contractHasEnoughTokensForPayout(uint256 tokenEuroConversionRate) public view returns (bool) {
        uint256 contractCanPayMax;
        uint256 contractShouldHaveMin;
        uint256 error;

        (contractCanPayMax, contractShouldHaveMin, error) = optimalPayoutDistribution(tokenEuroConversionRate);

        return contractShouldHaveMin <= balance();
    }

    /**
     * @dev with the current circulating supply - tokens from this contract and given token conversion rate
     * what would be the total roi to be paid.
     * @return uint256 roi of eligible token supply
     */
    function currentPayoutObligation(uint256 tokenEuroConversionRate) public view returns (uint256) {
        return roi(eligibleToken(), tokenEuroConversionRate);
    }

    /**
     * @dev the optimal payout distribution indicates the number of tokens the contract needs to hold to logically
     * be able to pay out all token holders.
     * with
     *
     * cs = circulating supply
     * ir = interest rate
     * tcr = token euro conversion rate
     *
     * contractShouldHaveMin := cs * ir / (ir + tcr)
     *
     * contractCanPayMax := cs * ctr / (ir + tcr)
     *
     * It is easily verifiable that
     *
     * cs = contractShouldHaveMin + contractCanPayMax + error
     *
     * @return (uint256, uint256, uint256) the maximum amount of tokens that logical can be paid, the minimum balance the contract should have to to fullfile obligation, the rounding error
    */
    function optimalPayoutDistribution(uint256 tokenEuroConversionRate) public view returns (uint256 contractCanPayMax, uint256 contractShouldHaveMin, uint256 error) {
        uint256 cri = INTEREST_RATE.add(tokenEuroConversionRate);

        contractShouldHaveMin = circulatingSupply().mul(INTEREST_RATE).div(cri);
        contractCanPayMax = circulatingSupply().mul(tokenEuroConversionRate).div(cri);
        error = circulatingSupply().sub(contractShouldHaveMin).sub(contractCanPayMax);

        return (contractCanPayMax, contractShouldHaveMin, error);
    }

    ///////////////////////////////////////////////////////////////////////////////////////////
    //  Pure
    ///////////////////////////////////////////////////////////////////////////////////////////

    /**
     * @dev shows the interest rate of BaaSToken
     * @return uint256 The interest rate
     */
    function interestRate() public pure returns (uint256) {
        return INTEREST_RATE;
    }

    /**
     * @dev shows the amount of token dividend at given interest and euro token conversion rate
     * @return uint256 The interest rate
     */
    function roi(uint256 token, uint256 tokenEuroConversionRate) public pure returns (uint256) {
        if (tokenEuroConversionRate == 0) {
            return 0;
        }

        return token.mul(INTEREST_RATE).div(tokenEuroConversionRate);
    }
}