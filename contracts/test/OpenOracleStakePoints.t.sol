pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../src/OpenOracleStakePoints.sol";
import "../src/ERC20Mock2.sol";

contract OpenOracleStakePointsTest is Test {

    OpenOracleStakePoints public stakeContract;
    ERC20Mock2 public token;

    function setUp() public {
        token = new ERC20Mock2("ERC20Mock2","E20");
        stakeContract = new OpenOracleStakePoints();
    }

    function testAddToken() public {
        stakeContract.addToken(address(token), 1000000);
        assertEq(stakeContract.tokenConfig(address(token)), 1000000);
    }

    function testFailAddTokenNotOwner() public {
        vm.startPrank(address(0));
        stakeContract.addToken(address(token), 1000000);
        vm.stopPrank();
    }

    function testFailAddTokenNotUnique() public {
        stakeContract.addToken(address(token), 1000000);
        stakeContract.addToken(address(token), 1000000);
    }

    function testFailAddTokenFractionIsZero() public {
        stakeContract.addToken(address(token), 0);
    }

    function testStake() public {
        address user = address(this);
        stakeContract.addToken(address(token), 1000000);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(token), 10**18);
        uint256 stake = stakeContract.tokenStake(user, address(token));
        assertEq(stakeContract.tokenStake(user, address(token)),10**18);
        assertEq(stakeContract.totalStake(user),1);
    }

    function testManyStake() public {
        address user = address(this);
        stakeContract.addToken(address(token), 1000000);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(token), 10**18);
        uint256 stake = stakeContract.tokenStake(user, address(token));
        assertEq(stakeContract.tokenStake(user, address(token)),10**18);
        assertEq(stakeContract.totalStake(user),1);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(token), 10**18);

        assertEq(stakeContract.tokenStake(user, address(token)),2*10**18);
        assertEq(stakeContract.totalStake(user),1*2);
    }

    function testLessStake() public {
        address user = address(this);
        stakeContract.addToken(address(token), 1000000);
        token.mint(user, 10**17);
        token.approve(address(stakeContract), 10**17);
        stakeContract.stake(address(token), 10**17);
        uint256 stake = stakeContract.tokenStake(user, address(token));
        assertEq(stakeContract.tokenStake(user, address(token)),10**17);
        assertEq(stakeContract.totalStake(user),0);
    }

    function testUnStake() public {
        testStake();
        address user = address(this);
        stakeContract.unStake(address(token), 10**18);
        uint256 stake = stakeContract.tokenStake(user, address(token));
        assertEq(stakeContract.tokenStake(user, address(token)),0);
        assertEq(stakeContract.totalStake(user),0);
    }

    function testLessUnStake() public {
        testStake();
        address user = address(this);
        stakeContract.unStake(address(token), 10**17);
        uint256 stake = stakeContract.tokenStake(user, address(token));
        assertEq(stakeContract.tokenStake(user, address(token)),9*10**17);
        assertEq(stakeContract.totalStake(user),0);
    }

    function testCheckPoint() public {
        address user = address(this);
        testStake();
        assertEq(stakeContract.getRealTimePoints(user),0);
        uint256 currentTime = block.timestamp;
        vm.warp(currentTime + 3600);
        assertEq(stakeContract.getRealTimePoints(user),1);
    }

    function testFailCheckPoint() public {
        address user = address(this);
        testStake();
        assertEq(stakeContract.getRealTimePoints(user),0);
        uint256 currentTime = block.timestamp;
        vm.warp(currentTime + 3500);
        assertEq(stakeContract.getRealTimePoints(user),1);
    }

}
