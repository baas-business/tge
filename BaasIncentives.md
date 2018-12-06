# BaaS Incentives 

Incentives is a custom smart contract that incentivizes contributor to the baas project.
The incentives are unlocked partially in a given time frame. The contract offers enough
flexibility to determine the vesting period and stages.

**Initial Supply**  | **Is Ownable** 
| :-------------: |:-------------:| 
*10m Tokens* | *true*

## Write Transactions

## setup()
**Signature:** 
    
    function setup() external onlyOwner returns (bool){...}
    
**Functionality:** Activates the contract and retrieves the balance information of the token contract to prepare its books.

**Params:**
no


## reward(account, amountPerStage, totalVestingStages, blockTimePerStage) 
**Signature:**

    function reward(address account, uint256 amountPerStage, uint totalVestingStages, uint blockTimePerStage) external onlyOwner returns (bool) {...}

**Functionality:** 

Creates an entry in storage with the param information. The total amount of rewarded token is given implicitly by 
    
    totalAmount = totalVestingStages x amountPerStage 

and the total time in blocks by 
    
    totalTime = totalVestingStages x blockTimePerStage
 
**Params:**

param | type | meaning
| :-------------: |:-------------:|:-------------|
*account* | address | the receiver of the incentive
*amountPerStage* | uint256 | the token amount unlocked per stage
*totalVestingStages* | uint | the total number of stages
*blockTimePerStage* | uint | blocks between 2 stages
  

## claim() 
**Signature:**

    function claim() external returns (uint256) {...}

**Functionality:** 

Claims all unlocked tokens if not claimed yet or not forfeited. This can claim more than one stage at the same time. 
*msg.sender* is the claimer.
 
**Params:** no  
  
    
    
## forfeited
**Signature:**

    function forfeited(address account) external onlyOwner returns (bool) {...}

**Functionality:** 

Stops vesting of an account and sends all unlocked tokens to given address. 
*msg.sender* is the claimer.
 
**Params:** 
 

param | type | meaning
| :-------------: |:-------------:|:-------------|
*account* | address | the account that will forfeit its vesting  
  
    
**Open Questions:**
* When is the first stage unlocked, do we need an offset? 
* Do we rather want to work with timestamps then blocktime?

**Difficulties:**
* no 