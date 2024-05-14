pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/IUserStakingHandler.sol";

contract OpenOracleUserStaking is Ownable {

    // token => true
    mapping(address => bool) public tokenConfig;
    // user => lastUpdateTime
    mapping(address => uint256) public lastUpdateTime;
    // user => stake
    mapping(address => uint256) public totalStake;
    // user => token => stake
    mapping(address => mapping(address => uint256)) public tokenStake;

    IUserStakingHandler private userStakingHandler;

    event Stake(address indexed user, address indexed token, uint256 amount);

    event UnStake(address indexed user, address indexed token, uint256 amount);

    event AddToken(address indexed token, bool active);


    constructor(IUserStakingHandler _userStakingHandler)Ownable() {
        userStakingHandler = _userStakingHandler;
    }

    function updateHandler(IUserStakingHandler _userStakingHandler) public onlyOwner{
        require(address(_userStakingHandler) != address(0), "UserStakingHandler cannot be 0 address");
        userStakingHandler = _userStakingHandler;
    }

    function addToken(address token, bool active) external onlyOwner{
        tokenConfig[token] = active;
        emit AddToken(token, active);
    }

    function stake(address token, uint256 amount) external {
        require(tokenConfig[token], "Token is not active");
        require(amount > 0, "Amount must be greater than 0");
        IERC20Metadata tokenContract = IERC20Metadata(token);
        tokenContract.transferFrom(msg.sender, address(this), amount);
        tokenStake[msg.sender][token] += amount;
        emit Stake(msg.sender, token, amount);
        try userStakingHandler.afterStake(msg.sender, token, amount) {
        } catch Error(string memory reason) {
            // catch failing revert() and require()
        } catch (bytes memory reason) {
            // catch failing assert()
        }
    }

    function unStake(address token, uint256 amount) external {
        require(tokenConfig[token], "Token is not active");
        require(amount > 0, "Amount must be greater than 0");
        require(tokenStake[msg.sender][token] >= amount, "Not enough tokens staked");
        tokenStake[msg.sender][token] -= amount;
        IERC20Metadata tokenContract = IERC20Metadata(token);
        tokenContract.transfer(msg.sender, amount);
        emit UnStake(msg.sender, token, amount);
        try userStakingHandler.afterUnstake(msg.sender, token, amount) {
        } catch Error(string memory reason) {
            // catch failing revert() and require()
        } catch (bytes memory reason) {
            // catch failing assert()
        }
    }

}
