// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Counter {
    uint256 public count;

    function increment() external {
        count += 1;
    }
    
    function decrement() external {
        count -= 1;
    }
    
    function get() external view returns (uint256) {
        return count;
    }

    function set(uint256 _count) external {
        count = _count;
    }

    function reset() external {
        count = 0;
    }

    function add(uint256 _count) external {
        count += _count;
    }   

    function subtract(uint256 _count) external {
        count -= _count;
    }

    function multiply(uint256 _count) external {
        count *= _count;
    }

    function divide(uint256 _count) external {
        require(_count != 0, "Cannot divide by zero");
        count /= _count;
    }
}