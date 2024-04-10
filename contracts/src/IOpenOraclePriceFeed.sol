// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IOpenOraclePriceFeed {
    /// @notice Call the task manager to request latest data
    function requestNewReport() external;

    /// @notice Saves the latest data from task manager in contract
    function saveLatestData(uint256 value) external;

    /// @notice Returns the latest data
    function latestRoundData() view external returns (uint256);
}
