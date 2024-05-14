pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "./interfaces/IUserStakingHandler.sol";

contract OpenOracleUserStakingHandler is IUserStakingHandler, OwnableUpgradeable {

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

    // user => boost (1 > BASE_FRACTION)
    mapping(address => uint256) public boosts;

    event Stake(address indexed user, address indexed token, uint256 amount, uint256 totalPoints);

    event UnStake(address indexed user, address indexed token, uint256 amount, uint256 totalPoints);

    event DropPoints(address indexed user, uint256 points);

    event AddToken(address indexed token, uint256 fraction);

    event AdjustBoost(address indexed user, uint256 boost);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() initializer public {
        __Ownable_init();
    }

    function dropPoints(address user, uint256 _points) external onlyOwner{
        require(_points > 0, "Points must be greater than 0");
        require(user != address(0), "User cannot be 0 address");
        points[user] += (_points * BASE_FRACTION);
        emit DropPoints(user, _points);
    }

    function adjustBoost(address user, uint256 boost) external onlyOwner{
        // The precision of boost is BASE_FRACTION
        require(boost > 0, "Boost must be greater than 0");
        require(user != address(0), "User cannot be 0 address");
        boosts[user] = boost;
        emit AdjustBoost(user, boost);
    }

    function addToken(address token, uint256 fraction) external onlyOwner{
        require(fraction > 0, "Fraction must be greater than 0");
        require(tokenConfig[token] <= 0, "Token is already in the list");
        tokenConfig[token] = fraction;
        emit AddToken(token, fraction);
    }

    function afterStake(address user, address token, uint256 amount) override external {
        require(tokenConfig[token] > 0, "Token is not in the list");
        require(amount > 0, "Amount must be greater than 0");
        IERC20Metadata tokenContract = IERC20Metadata(token);
        updatePoint(user, tokenContract, amount, true);
        emit Stake(user, token, amount, points[user]);
    }

    function afterUnstake(address user, address token, uint256 amount) override external {
        require(tokenConfig[token] > 0, "Token is not in the list");
        require(amount > 0, "Amount must be greater than 0");
        require(tokenStake[user][token] >= amount, "Not enough tokens staked");
        IERC20Metadata tokenContract = IERC20Metadata(token);
        updatePoint(user, tokenContract, amount, false);
        emit UnStake(user, token, amount, points[user]);
    }

    function updatePoint(address user, IERC20Metadata tokenContract, uint256 amount, bool isStake) private{
        address token = address(tokenContract);
        points[user] += _calcPoints(user);
        lastUpdateTime[user] = block.timestamp;
        if(totalStake[user] > 0){
            totalStake[user] -= (tokenStake[user][token] * tokenConfig[token] / 10 ** tokenContract.decimals());
        }
        isStake ? (tokenStake[user][token] += amount) : (tokenStake[user][token] -= amount);
        totalStake[user] += (tokenStake[user][token] * tokenConfig[token] / 10 ** tokenContract.decimals());
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
        bool isBoost = boosts[user] > 0;
        uint256 _points = totalStake[user] * PER_REWARD * (block.timestamp - lastUpdateTime[user]) / SECONDS_PER_HOUR;
        return isBoost ? (_points * boosts[user] / BASE_FRACTION) : _points;
    }
}
