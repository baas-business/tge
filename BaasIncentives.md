# BaaS Incentives 


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

Creates an entry in storage with the param information. The total amount of rewarded token is given implicitly by **totalVestingStages x amountPerStage** and
the total time by **totalVestingStages x blockTimePerStage**.
 
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
  
    
**Open Questions:**
* When is the first stage unlocked, do we need an offset? 
* Do we rather want to work with timestamps then blocktime?

**Difficulties:**
* no 