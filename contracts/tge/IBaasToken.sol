pragma solidity 0.5;


interface IBaasToken  {
    function setPause(bool pause) external;

    function burnTokensFromPot(address potAddress, uint256 amount) external returns(bool);

    function tokenHolderSnapShot() external returns (address[] memory);

    function circulatingSupply() external view returns (uint256);

    function pots() external view returns (address, address, address, address);

    function totalSupply() external view returns (uint256);

    function balanceOf(address who) external view returns (uint256);

    function allowance(address owner, address spender)
    external view returns (uint256);

    function transfer(address to, uint256 value) external returns (bool);

    function approve(address spender, uint256 value)
    external returns (bool);

    function transferFrom(address from, address to, uint256 value)
    external returns (bool);

    event Transfer(
        address indexed from,
        address indexed to,
        uint256 value
    );

    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    event Paused(bool paused);

    event SetupCompleted();
}