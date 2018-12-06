# BaaS Token

Token is a ERC20 compatible contract. At setup the contract receives 
4 other contract address representing escrow, incentives, founder, private placement.
It adds functionality to the ERC20 Standard:

1. It provides the circulating supply by subtracting the pots from the total supply. 
2. It is pausable in terms of token transfers
3. It manages a list of token holders giving the possibility to snapshot the list of holders
at any time.

**Total Supply**  | **Is Ownable** 
| :-------------: |:-------------:| 
*100m Tokens* | *true*

## Write Transactions

### setup
**Signature:** 
    
     function setup(
            address escrowAddress,
            address ppAddress,
            address founderAddress,
            address incentivesAddress)
        external onlyOwner returns (bool) {...}
    
**Functionality:** stores the pot addresses and mints tokens according to the prospect.

**Params:**

param | type | meaning
| :-------------: |:-------------:|:-------------|
*escrowAddress* | address | escrow address supplied with 60m tokens
*ppAddress* | address | private placement address supplied with 20m tokens 
*founderAddress* | address | founder address supplied with 10m tokens
*incentivesAddress* | address | incentives placement address supplied with 10m tokens



###  setPause
**Signature:**

    function setPause(bool pause) external onlyOwner {...}

**Functionality:** Gives the owner the chance to pause transfers for maintainance
**Params:** 

param | type | meaning
| :-------------: |:-------------:|:-------------|
*pause* | bool | pauses or unpauses transfers



###  burnTokensFromPot
**Signature:**

    function burnTokensFromPot(address potAddress, uint256 amount) external returns (bool) {...}

**Functionality:** Gives the pot a way to burn tokens.
**Params:** 

param | type | meaning
| :-------------: |:-------------:|:-------------|
*potAddress* | address | the address tokens shall be burned from
*amount* | amount | the amount of tokens shall be burned
    
    
**Open Questions:**
---
* Do we need a function recoverFromFailure to shift remaining tokens to a substitute pot contract (I would say yes)? 
* Does the owner really need to have the possibility to burn tokens?

**Difficulties:**
---
* none 