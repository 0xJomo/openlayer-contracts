// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IUserStakingHandler {
    function afterStake(address user, address token, uint256 amount) external;
    function afterUnstake(address user, address token, uint256 amount) external;
}
