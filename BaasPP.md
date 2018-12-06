# BaaS Private Placement 

Private Placement is a custom contract that holds 20m Token reserved
for early investors in a private round in custody. It allows to provide
token to given investors and keeps book on the provisions.

**Initial Supply**  | **Is Ownable** 
| :-------------: |:-------------:| 
*20m Tokens* | *true*

## Write Transactions

### provideToken
**Signature:** 
    
    function provideToken(address account, uint256 amount, uint8 discountType) external onlyOwner returns (bool) {
    
**Functionality:** Sends amoutn token to account.

**Params:**

param | type | meaning
| :-------------: |:-------------:|:-------------|
*account* | address | receiver of tokens
*amount* | uint256 | amount sent
*discountType* | uint8 | either 1 = DISCOUNTED or 2 = NOT_DISCOUNTED


###  burnRest
**Signature:**

    function burnRest() external onlyOwner returns (bool) {...}

**Functionality:** Burns token if not sold in private placement
**Params:** no

  
    
**Open Questions:**
---
* Do we need to burn the tokens or would be enough to send them to Escrow? 

**Difficulties:**
---
* no 