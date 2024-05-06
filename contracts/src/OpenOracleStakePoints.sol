pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract OpenOracleStakePoints is Ownable {

    uint256 private constant BASE_FRACTION = 1000000;
    uint256 private constant SECONDS_PER_HOUR = 3600;


    uint256 private perReward = 1;

    // token => fraction
    mapping(address => uint256) public tokenConfig;
    // user => lastUpdateTime
    mapping(address => uint256) public lastUpdateTime;
    // user => stake
    mapping(address => uint256) public totalStake;
    // user => token => stake
    mapping(address => mapping(address => uint256)) public tokenStake;
    // user => points
    mapping(address => uint256) public points;

    constructor() {
    }

    function addToken(address token, uint256 fraction) external onlyOwner{
        require(fraction > 0, "Fraction must be greater than 0");
        require(tokenConfig[token] <= 0, "Token is already in the list");
        tokenConfig[token] = fraction;
    }

    function stake(address token, uint256 amount) external {
        require(tokenConfig[token] > 0, "Token is not in the list");
        require(amount > 0, "Amount must be greater than 0");
        IERC20Metadata tokenContract = IERC20Metadata(token);
        tokenContract.transferFrom(msg.sender, address(this), amount);
        updatePoint(tokenContract, amount, true);
    }

    function unStake(address token, uint256 amount) external {
        require(tokenConfig[token] > 0, "Token is not in the list");
        require(amount > 0, "Amount must be greater than 0");
        require(tokenStake[msg.sender][token] >= amount, "Not enough tokens staked");
        IERC20Metadata tokenContract = IERC20Metadata(token);
        updatePoint(tokenContract, amount, false);
        tokenContract.transfer(msg.sender, amount);
    }

    function updatePoint(IERC20Metadata tokenContract, uint256 amount, bool isStake) private{
        address token = address(tokenContract);
        points[msg.sender] += (totalStake[msg.sender] * perReward * (block.timestamp - lastUpdateTime[msg.sender]) / SECONDS_PER_HOUR);
        lastUpdateTime[msg.sender] = block.timestamp;
        if(totalStake[msg.sender] > 0){
            totalStake[msg.sender] -= (tokenStake[msg.sender][token] * tokenConfig[token] / BASE_FRACTION / 10 ** tokenContract.decimals());
        }
        isStake ? (tokenStake[msg.sender][token] += amount) : (tokenStake[msg.sender][token] -= amount);
        totalStake[msg.sender] += (tokenStake[msg.sender][token] * tokenConfig[token] / BASE_FRACTION / 10 ** tokenContract.decimals());
    }

    function getPoints(address user) external view returns(uint256){
        return points[user];
    }

    function getRealTimePoints(address user) external view returns(uint256){
        return points[user] + (totalStake[user] * perReward * (block.timestamp - lastUpdateTime[user]) / SECONDS_PER_HOUR);
    }
}
