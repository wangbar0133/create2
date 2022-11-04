// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Deployer {

    bytes public initCode;
    
    function  getInitCode() public {
        initCode = type(Test).creationCode;
    }

    function deploy(uint _salt) public returns(address){
      
        address addr;
        bytes memory bytecode = type(Test).creationCode;
        assembly {
            addr := create2(0, add(bytecode, 0x20), mload(bytecode), _salt)
        }
        return addr;
  }
}

contract Test {

    address public owner;
    event Ev(); 

    function met() public {
        owner = msg.sender;
    }

    function get() public {
        require(owner == msg.sender);
        emit Ev();
    }
}