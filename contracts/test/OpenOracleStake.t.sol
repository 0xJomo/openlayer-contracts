pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../src/OpenOracleStake.sol";
import "../src/ERC20Mock2.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OpenOracleStakeTest is Test {

    OpenOracleStake public stakeContract;
    ERC20Mock2 public token;
    event DropPoints(address indexed user, uint256 points);

    function setUp() public {
        token = new ERC20Mock2("ERC20Mock2","E20");
        OpenOracleStake stakeContractImpl = new OpenOracleStake();
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(address(stakeContractImpl),address(proxyAdmin), abi.encodeWithSelector(stakeContractImpl.initialize.selector));
        stakeContract = OpenOracleStake(address(proxy));
    }

    function testDropPoints() public {
        address user = address(this);
        vm.expectEmit(true,false,false,true,address(stakeContract));

        // We emit the event we expect to see.
        emit DropPoints(user, 10**8);
        // drop one point
        stakeContract.dropPoints(user, 10**8);
        assertEq(stakeContract.getPoints(user), 10**8);


    }

    function testFailDropPointsZero() public {
        address user = address(this);
        // drop one point
        stakeContract.dropPoints(user, 0);
    }

    function testFailDropPointsUserNull() public {
        address user = address(this);
        // drop one point
        stakeContract.dropPoints(address(0), 10**8);
    }

    function testFailDropPointsNotOwner() public {
        address user = address(this);
        // drop one point
        vm.startPrank(address(0));
        stakeContract.dropPoints(address(0), 10**8);
        vm.stopPrank();
    }

    function testAdjustBoost() public {
        stakeContract.adjustBoost(address(this), 1000000);
        assertEq(stakeContract.boosts(address(this)), 1000000);
    }

    function testFailAdjustBoostNotOwner() public {
        vm.startPrank(address(0));
        stakeContract.adjustBoost(address(this), 1000000);
        vm.stopPrank();
    }

    function testFailAdjustBoostZero() public {
        stakeContract.adjustBoost(address(this), 0);
    }

    function testFailAdjustBoostUserNull() public {
        stakeContract.adjustBoost(address(0), 1000000);
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
