/*
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../src/OpenOracleUserStaking.sol";
import "../src/ERC20Mock2.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OpenOracleUserStakingTest is Test {

    OpenOracleUserStaking public stakeContract;
    ERC20Mock2 public token;
    event DropPoints(address indexed user, uint256 points);

    function setUp() public {
        token = new ERC20Mock2("ERC20Mock2","E20");
        stakeContract = new OpenOracleUserStaking(address(0));
    }


    function testAddToken() public {
        stakeContract.addToken(address(token), true);
        assertEq(stakeContract.tokenConfig(address(token)), true);
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
        assertEq(stakeContract.tokenStake(user, address(token)),10**18);
        assertEq(stakeContract.totalStake(user),1*10**6);
    }

    function testFailStakeZero() public {
        address user = address(this);
        stakeContract.addToken(address(token), 1000000);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(token), 0);
    }

    function testFailStakeNotExist() public {
        address user = address(this);
        stakeContract.addToken(address(token), 1000000);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(0), 0);
    }

    function testManyStake() public {
        address user = address(this);
        stakeContract.addToken(address(token), 1000000);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(token), 10**18);
        assertEq(stakeContract.tokenStake(user, address(token)),10**18);
        assertEq(stakeContract.totalStake(user),1*10**6);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(token), 10**18);

        assertEq(stakeContract.tokenStake(user, address(token)),2*10**18);
        assertEq(stakeContract.totalStake(user),1*2*10**6);
    }

    function testLessStake() public {
        address user = address(this);
        stakeContract.addToken(address(token), 1000000);
        token.mint(user, 10**17);
        token.approve(address(stakeContract), 10**17);
        stakeContract.stake(address(token), 10**17);
        assertEq(stakeContract.tokenStake(user, address(token)),10**17);
        assertEq(stakeContract.totalStake(user),10**5);
    }

    function testUnStake() public {
        testStake();
        address user = address(this);
        stakeContract.unStake(address(token), 10**18);
        assertEq(stakeContract.tokenStake(user, address(token)),0);
        assertEq(stakeContract.totalStake(user),0);
    }

    function testManyUnStake() public {
        testManyStake();
        address user = address(this);
        stakeContract.unStake(address(token), 11*10**17);
        assertEq(stakeContract.tokenStake(user, address(token)),9*10**17);
        assertEq(stakeContract.totalStake(user),9*10**5);
        stakeContract.unStake(address(token), 9*10**17);
        assertEq(stakeContract.tokenStake(user, address(token)),0);
        assertEq(stakeContract.totalStake(user),0);
    }

    function testUnStakeWithLess() public {
        testLessStake();
        address user = address(this);
        stakeContract.unStake(address(token), 10**17);
        assertEq(stakeContract.tokenStake(user, address(token)),0);
        assertEq(stakeContract.totalStake(user),0);
    }

    function testFailUnStakeZero() public {
        testStake();
        stakeContract.unStake(address(token), 0);
    }

    function testFailUnStakeNotExist() public {
        testStake();
        stakeContract.unStake(address(0), 0);
    }

    function testFailUnStakeNotEnough() public {
        testStake();
        stakeContract.unStake(address(token), 2*10**18);
    }

    function testLessUnStake() public {
        testStake();
        address user = address(this);
        stakeContract.unStake(address(token), 10**17);
        assertEq(stakeContract.tokenStake(user, address(token)),9*10**17);
        assertEq(stakeContract.totalStake(user),9*10**5);
    }

    function testCheckPoint() public {
        address user = address(this);
        testStake();
        assertEq(stakeContract.getRealTimePoints(user),0);
        uint256 currentTime = block.timestamp;
        vm.warp(currentTime + 3600);
        assertEq(stakeContract.getRealTimePoints(user),10**8);
        assertEq(stakeContract.getPoints(user),0);
    }

    function testLessOneHourCheckPoint() public {
        address user = address(this);
        testStake();
        assertEq(stakeContract.getRealTimePoints(user),0);
        uint256 currentTime = block.timestamp;
        vm.warp(currentTime + 3500);
        assertEq(stakeContract.getRealTimePoints(user),97222222);
    }

}
*/
