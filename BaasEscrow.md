# BaaS Escrow


**Initial Supply**  | **Is Ownable** 
| :-------------: |:-------------:| 
*60m Tokens* | *true*

## Write Transactions

### setup
**Signature:** 
    
    function setup(uint256 vestingPeriod) external onlyOwner returns (bool){...}
    
**Functionality:** Activates the contract and specifies a period in blocks until the first raise can happen.

**Params:**

param | type | meaning
| :-------------: |:-------------:|:-------------|
*vestingPeriod* | uint256 | the time in blocks until the first raise can is unlocked.


### raiseCapital
**Signature:**
    
    raiseCapital(uint256 amount, uint id) external onlyOwner returns (bool) {...}

**Functionality:** Raises capital by making amount tokens transferable. It also snapshots the current
token holders and stores them in contract.

**Params:**


param | type | meaning
| :-------------: |:-------------:|:-------------| 
*amount*| uint256 | the total amount raised in this round
*id* | uint | the id to refer to at delivery 
      
         
### provideToken
**Signature:**

    function provideToken(address account, uint256 amount, uint256 conversionRate) external onlyOwner returns (bool)
    
**Functionality:** Provides token to account and stores the provision for the record. 

**Params:**

param | type | meaning
| :-------------: |:-------------:|:-------------|
*account* | address | the receiver(investor) of tokens.
*amount* | uint256 | the amount the ince.
*conversionRate* | uint256 | The token to Euro conversion rate the investor purhchased the token (Only for the records).

    function provideToken(address account, uint256 amount, uint256 conversionRate) external onlyOwner returns (bool)
    
    
**Open Questions:**
---
* Is there a halting period for each raise?
* Do we lock the time until the investor can claim the tokens and block provision to others (its just masquerade)?
* Is there a necessity to burn token if they weren't sold during a capital raise. 

**Difficulties:**
---
* The snapshot can cost a lot of gas if the holder count goes exceeds 10.000! 