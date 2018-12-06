# BaaS Founder 


The Founder contract is a custom smart contract that acts as a custodian for 
the tokens rewarded to the founder. After the vesting time the founders in the
form of the contract owner can withdraw tokens.

**Initial Supply**  | **Is Ownable** 
| :-------------: |:-------------:| 
*10m Tokens* | *true*

## Write Transactions

### setup
**Signature:** 
    
    function setup(uint256 vestingPeriod) external onlyOwner returns (bool){...}
    
**Functionality:** Activates the contract and specifies a period in blocks until the founder can withdraw the tokens.

**Params:**

param | type | meaning
| :-------------: |:-------------:|:-------------|
*vestingPeriod* | uint256 | the time in blocks until the first raise can is unlocked.


### withdraw
**Signature:**

    function withdraw(address receiver, uint256 amount) external onlyOwner returns (bool) {...}

**Functionality:** Sends tokens to the receiver if enough tokens are available.
**Params:**

param | type | meaning
| :-------------: |:-------------:|:-------------|
*receiver* | address | the receiver of the tokens
*amount* | address | the amount receiver receives
  
    
**Open Questions:**
* Do we specify the founder address beforehand (not necessary)?
* Do we want to have more than one receiver address (this would communicate to the public who was served with how much)? 

**Difficulties:**
* no 