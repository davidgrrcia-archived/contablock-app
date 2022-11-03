// https://eips.ethereum.org/EIPS/eip-20
// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "./IERC20.sol";

contract HoldContract is IERC20 {

    uint256 private supply;
    address private contractOwner;

    mapping(address => uint256) private balances;
    mapping(address => mapping(address => uint256)) private allowed;

    mapping(address => uint256) private balancesOnHold;
    mapping(uint256 => address) private holdSender;
    mapping(uint256 => address) private holdRecipient;
    mapping(uint256 => uint256) private holdAmount;

    constructor(uint256 _initialAmount) {
        contractOwner = msg.sender;
        balances[msg.sender] = _initialAmount;
        supply = _initialAmount;
    }

    function totalSupply() public override view returns (uint256){
        return supply;
    }

    function balanceOf(address account) public override view returns (uint256 balance) {
        return balances[account];
    }

    function balanceOnHoldOf(address account) public view returns (uint256 balance) {
        return balancesOnHold[account];
    }

    function transfer(address recipient, uint256 amount) public override returns (bool success) {
        require(balances[msg.sender] >= amount, "token balance is lower than the value requested");
        balances[msg.sender] -= amount;
        balances[recipient] += amount;
        emit Transfer(msg.sender, recipient, amount);
        return true;
    }

    function hold(address recipient, uint256 amount, uint256 holdId) public returns (bool success) {
        require(balances[msg.sender] >= amount, "token balance is lower than the value requested");
        require(holdSender[holdId] == address(0), "hold id must be unique");
        balances[msg.sender] -= amount;
        balancesOnHold[recipient] += amount;
        holdSender[holdId] = msg.sender;
        holdRecipient[holdId] = recipient;
        holdAmount[holdId] = amount;
        return true;
    }

    function transferFrom(address sender, address recipient, uint256 amount) public override returns (bool success) {
        uint256 allowanceAmount = allowed[sender][msg.sender];
        require(balances[sender] >= amount && allowanceAmount >= amount, "token balance or allowance is lower than amount requested");
        balances[recipient] += amount;
        balances[sender] -= amount;
        allowed[sender][msg.sender] -= amount;
        emit Transfer(sender, recipient, amount);
        return true;
    }

    function holdFrom(address sender, address recipient, uint256 amount, uint256 holdId) public returns (bool success) {
        uint256 allowanceAmount = allowed[sender][msg.sender];
        require(balances[sender] >= amount && allowanceAmount >= amount, "token balance or allowance is lower than amount requested");
        balancesOnHold[recipient] += amount;
        balances[sender] -= amount;
        allowed[sender][msg.sender] -= amount;
        holdSender[holdId] = sender;
        holdRecipient[holdId] = recipient;
        holdAmount[holdId] = amount;
        return true;
    }

    function approve(address spender, uint256 amount) public override returns (bool success) {
        allowed[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function allowance(address owner, address spender) public override view returns (uint256 remaining) {
        return allowed[owner][spender];
    }

    function executeHold(uint256 holdId) public returns (bool success){
        uint256 amount = holdAmount[holdId];
        address recipient = holdRecipient[holdId];

        require(recipient != address(0), "invalid hold id");
        require(holdSender[holdId] == msg.sender, "it can only be executed by the same address that created the hold.");
        require(amount != 0, "amount on hold already transfered");

        balancesOnHold[recipient] -= amount;
        balances[recipient] += amount;
        holdAmount[holdId] = 0;
        return true;
    }

    function removeHold(uint256 holdId) public returns (bool success){
        uint256 amount = holdAmount[holdId];
        address sender = holdSender[holdId];
        address recipient = holdRecipient[holdId];

        require(sender != address(0), "invalid hold id");
        require(msg.sender == contractOwner, "it can only be executed by the contract owner.");
        require(amount != 0, "amount on hold already transfered");

        balancesOnHold[recipient] -= amount;
        balances[sender] += amount;
        holdAmount[holdId] = 0;
        return true;
    }


}