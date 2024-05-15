pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../src/OpenOracleUserStaking.sol";
import "../src/interfaces/IUserStakingHandler.sol";
import "../src/OpenOracleUserStakingHandler.sol";
import "../src/ERC20Mock2.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OpenOracleUserStakingTest is Test {

    OpenOracleUserStaking public stakeContract;
    ERC20Mock2 public token;
    OpenOracleUserStakingHandler public stakeContractHandler;

    function setUp() public {
        token = new ERC20Mock2("ERC20Mock2","E20");
        OpenOracleUserStakingHandler stakeContractImpl = new OpenOracleUserStakingHandler();
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(address(stakeContractImpl),address(proxyAdmin), abi.encodeWithSelector(stakeContractImpl.initialize.selector));
        stakeContractHandler = OpenOracleUserStakingHandler(address(proxy));
        stakeContract = new OpenOracleUserStaking(IUserStakingHandler(address(stakeContractHandler)));

    }


    function testAddToken() public {
        stakeContract.addToken(address(token), true);
        assertEq(stakeContract.tokenConfig(address(token)), true);
    }

    function testFailAddTokenNotOwner() public {
        vm.startPrank(address(0));
        stakeContract.addToken(address(token), true);
        vm.stopPrank();
    }

    function testStake() public {
        address user = address(this);
        stakeContract.addToken(address(token), true);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(token), 10**18);
        assertEq(stakeContract.tokenStake(user, address(token)),10**18);
    }

    function testFailStakeZero() public {
        address user = address(this);
        stakeContract.addToken(address(token), true);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(token), 0);
    }

    function testFailStakeNotExist() public {
        address user = address(this);
        stakeContract.addToken(address(token), true);
        token.mint(user, 10**18);
        token.approve(address(stakeContract), 10**18);
        stakeContract.stake(address(0), 0);
    }

    function testUnStake() public {
        testStake();
        address user = address(this);
        stakeContract.unStake(address(token), 10**18);
        assertEq(stakeContract.tokenStake(user, address(token)),0);
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


    function testUpdateHandler() public {
        stakeContract.updateHandler(IUserStakingHandler(address(1)));
    }

    function testFailUpdateHandler() public {
        stakeContract.updateHandler(IUserStakingHandler(address(0)));
    }

}
