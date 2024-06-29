// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract OpenOracleOperatorRole is AccessControl, Initializable {

    using EnumerableSet for EnumerableSet.Bytes32Set;

    event OperatorRoleAdded(address indexed operator, bytes32 indexed role);

    event OperatorRoleRemoved(address indexed operator, bytes32 indexed role);

    event OperatorTypeAdded(bytes32 indexed role);

    EnumerableSet.Bytes32Set private operatorType;

    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");

    constructor(
    ) {
    }

    function initialize(
        address defaultAdmin
    ) public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
    }

    function addOperatorRole(address operator, bytes32 role) public onlyRole(OPERATOR_ROLE){
        require(operatorType.contains(role), "operator role not exist");
        _grantRole(role, operator);
        emit OperatorRoleAdded(operator, role);
    }

    function removeOperatorRole(address operator, bytes32 role) public onlyRole(OPERATOR_ROLE){
        require(operatorType.contains(role), "operator role not exist");
        require(hasRole(role, operator), "operator role not exist");
        _revokeRole(role, operator);
        emit OperatorRoleRemoved(operator, role);
    }

    function addOperatorType(bytes32 role) public onlyRole(DEFAULT_ADMIN_ROLE){
        //bytes32 public constant OPERATOR_NODE_ROLE = keccak256("OPERATOR_NODE_ROLE");
        //bytes32 public constant OPERATOR_BROWSER_EXTENSION_ROLE = keccak256("OPERATOR_BROWSER_EXTENSION_ROLE");
        operatorType.add(role);
        emit OperatorTypeAdded(role);
    }
}