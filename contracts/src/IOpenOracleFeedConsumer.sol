// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;


interface IOpenOracleFeedConsumer {

    function fulfillResult(
        uint256 _requestId,
        bytes calldata _result,
        bytes calldata _data
    ) external;
}
