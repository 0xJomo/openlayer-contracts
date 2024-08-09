// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./IOpenOracleFeedConsumer.sol";
import "./IOpenOracleVRFFeed.sol";

contract OpenOracleVRFFeedConsumer is IOpenOracleFeedConsumer {

    uint256 public requestId;
    address public vrfFeed;

    mapping(uint256 => Response) public results;

    event Request(uint256 requestId);

    struct Response{
        bytes result;
        bytes data;
    }
    function request() public{
        bytes memory data;
        IOpenOracleVRFFeed(vrfFeed).createNewTaskWithCallback(data,requestId);
        emit Request(requestId);
        requestId++;
    }

    function fulfillResult(
        uint256 _requestId,
        bytes calldata _result,
        bytes calldata _data
    ) external {
        bytes memory result = _result;
        bytes memory data = _data;
        results[requestId] = Response(result,data);
    }

    function otherFunction() public{
        // todo deal service use result
        return;
    }
}
