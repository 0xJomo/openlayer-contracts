// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./IOpenOracleFeedConsumer.sol";
import "./IOpenOracleCommonDataFeed.sol";

contract OpenOracleDataFeedConsumer is IOpenOracleFeedConsumer {

    uint256 public requestId;
    address public dateFeed;

    mapping(uint256 => Response) public results;

    event Request(uint256 requestId);

    struct Response{
        bytes result;
        bytes data;
    }

    constructor(address _dateFeed){
        dateFeed = _dateFeed;
    }

    function request() public{
        IOpenOracleCommonDataFeed(dateFeed).requestNewReportCallback(uint8(0),requestId);
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
        results[_requestId] = Response(result,data);
    }

    function otherFunction() public{
        // todo deal service use result
        return;
    }
}
