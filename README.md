# Baas Main Net Token And Token Distribution Contracts

## This repositories consists of 6 Smart Contracts creating the token and managing the token distribution

* ERC20 Token Contract
* ROI Contract for rewarding token holders with interest
* Incentive Contract managing wallets of collaborates with vesting period
* Private Placement Contract managing the initial investment 
* Founder Contract
* Escrow Contract


## BaasFounder

The BaasFounder Contract manages the token funds for founder1 and founder2.
A vesting period (in block time) dictates the possibility to withdraw tokens.
Every year (2.200.000 Blocks) 25% of the funds dedicated to the founders are
freed.

Test
   
    truffle test
   
Build

    truffle compile
    
Current Contract:

//https://ropsten.etherscan.io/address/0xe27396555f2a682cef62f7ec54d4fa9ef0a70264
    