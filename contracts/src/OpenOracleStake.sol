pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract OpenOracleStake is OwnableUpgradeable {

    // deal amount
    uint256 private constant BASE_FRACTION = 10**6;
    uint256 private constant SECONDS_PER_HOUR = 3600;
    // deal time reward
    uint256 public constant PER_REWARD = 10**8;

    // token => fraction
    mapping(address => uint256) public tokenConfig;
    // user => lastUpdateTime
    mapping(address => uint256) public lastUpdateTime;
    // user => stake
    mapping(address => uint256) public totalStake;
    // user => token => stake
    mapping(address => mapping(address => uint256)) public tokenStake;
    // user => points
    mapping(address => uint256) private points;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() initializer public {
        __Ownable_init();
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
        points[msg.sender] += _calcPoints(msg.sender);
        lastUpdateTime[msg.sender] = block.timestamp;
        if(totalStake[msg.sender] > 0){
            totalStake[msg.sender] -= (tokenStake[msg.sender][token] * tokenConfig[token] / 10 ** tokenContract.decimals());
        }
        isStake ? (tokenStake[msg.sender][token] += amount) : (tokenStake[msg.sender][token] -= amount);
        totalStake[msg.sender] += (tokenStake[msg.sender][token] * tokenConfig[token] / 10 ** tokenContract.decimals());
    }

    /**
     * @dev Returns the amount of points a user has earned.
     */
    function getPoints(address user) external view returns(uint256){
        return points[user] / BASE_FRACTION;
    }

    /**
     * @dev Returns the amount of points a user has earned and the current time.
     */
    function getRealTimePoints(address user) external view returns(uint256){
        return (points[user] + _calcPoints(user)) / BASE_FRACTION;
    }

    function _calcPoints(address user) private view returns(uint256){
        return totalStake[user] * PER_REWARD * (block.timestamp - lastUpdateTime[user]) / SECONDS_PER_HOUR;
    }
}
