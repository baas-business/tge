# BaaS ROI

ROI is a custom contract that has no tokens. The circulating supply and
the token conversion rate determine how many tokens each token holder receives
at the end of the year with interest rate 0f 9 %. The day of payout can be 
chosen by the owner.

**Initial Supply**  | **Is Ownable** 
| :-------------: |:-------------:| 
*0m Tokens* | *true*

## Write Transactions

### payoutAll
**Signature:** 
    
    function payoutAll(uint256 tokenEuroConversionRate) external onlyOwner returns (bool) {...}
    
**Functionality:** pays out 

**Params:**

param | type | meaning
| :-------------: |:-------------:|:-------------|
*tokenEuroConversionRate* | uint256 | the token conversion rate chosen by the owner.



###  withdraw
**Signature:**

    function withdraw(address receiver) external onlyOwner returns (bool) {...}

**Functionality:** Gives the owner the chance to retrieve the tokens that were not paid out.
**Params:** 

param | type | meaning
| :-------------: |:-------------:|:-------------|
*receiver* | address | receiver of tokens

  
    
**Open Questions:**
---
* Should the conversion rate being constantly updated or on day of payout being transmitted as a parameter? 

 