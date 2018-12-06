# BaasFounder
 

. | .
--- | --- 
**Initial Supply** | *20m Tokens* 
**Is Ownable** | *true*

### Write Transactions

#### 1 setup(vestingPeriod)
Activates the contract and specifies a period in blocks until the a raise can happen.

    function setup(uint256 vestingPeriod) external onlyOwner returns (bool)

#### 2 raiseCapital(amount, id)
Raises capital by making amount tokens transferable. It also snapshots the current
toke holders.
     
    raiseCapital(uint256 amount, uint id) external onlyOwner returns (bool) 
         
#### 3 provideToken(account, amount, conversionRate)
Provides token to account and for the record stores the conversion rate.

    function provideToken(address account, uint256 amount, uint256 conversionRate) external onlyOwner returns (bool) 