// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "hardhat/console.sol";

contract ContaBlockContract {
    address public owner;
    mapping(address => bool) private userAccess;

    string public constant CREATED_STATE = 'CREATED';
    string public constant PUBLISHED_STATE = 'PUBLISHED';

    struct Entry {
        string id;
        string accountId;
        string multipleEntryId;
        string currency;
        string concept;
        string date;
        uint256 debitBalance;
        uint256 creditBalance;
        string state;
    }

    mapping(address => mapping(string => Entry)) entries;
    mapping(address => mapping(string => bool)) entryExist;
    mapping(address => string[]) multipleEntryIds;
    mapping(address => mapping(string => bool)) multipleEntryIdExist;
    mapping(address => string[]) accountIds;
    mapping(address => mapping(string => bool)) accountIdExist;

    modifier onlyOwner() {
        require(
            msg.sender == owner,
            'Only the contract owner can perform this action'
        );
        _;
    }

    modifier hasAccess() {
        console.log('aaaaa: ',msg.sender);
        console.log('bbbbb: ',userAccess[msg.sender]);
        require(
            userAccess[msg.sender] == true,
            'You do not have access to this information'
        );
        _;
    }

    constructor() {
        owner = msg.sender;
        userAccess[msg.sender] = true;
    }

    function grantAccess(address user) public onlyOwner {
        console.log('grantAccess: ',user);
        userAccess[user] = true;
        console.log('grantAccess: ',userAccess[user]);
    }

    function revokeAccess(address user) public onlyOwner {
        userAccess[user] = false;
    }

    function createEntry(
        string memory _id,
        string memory _accountId,
        string memory _multipleEntryId,
        string memory _currency,
        string memory _concept,
        string memory _date,
        uint256 _debitBalance,
        uint256 _creditBalance
    ) public hasAccess {
        require(!entryExist[msg.sender][_id], 'Entry already exists');

        Entry memory _entry = Entry(
            _id,
            _accountId,
            _multipleEntryId,
            _currency,
            _concept,
            _date,
            _debitBalance,
            _creditBalance,
            CREATED_STATE
        );

        entryExist[msg.sender][_id] = true;
        entries[msg.sender][_id] = _entry;

        if (!multipleEntryIdExist[msg.sender][_multipleEntryId]) {
            multipleEntryIds[msg.sender].push(_multipleEntryId);
            multipleEntryIdExist[msg.sender][_multipleEntryId] = true;
        }

        if (!accountIdExist[msg.sender][_accountId]) {
            accountIds[msg.sender].push(_accountId);
            accountIdExist[msg.sender][_accountId] = true;
        }

        emit EntryCreated(_entry);
    }

    event EntryCreated(Entry entry);

    function getEntryData(string memory _id)
        public
        view
        hasAccess
        returns (Entry memory)
    {
        return entries[msg.sender][_id];
    }

    function listMultipleEntryIds()
        public
        view
        hasAccess
        returns (string[] memory)
    {
        return multipleEntryIds[msg.sender];
    }

    function listAccountIds() public view hasAccess returns (string[] memory) {
        return accountIds[msg.sender];
    }

    function publishEntry(string memory _id) public hasAccess {
        require(entryExist[msg.sender][_id], 'Entry does not exists');

        Entry memory _entry = entries[msg.sender][_id];

        //compare state strings
        require(
            keccak256(abi.encodePacked(_entry.state)) ==
                keccak256(abi.encodePacked(CREATED_STATE)),
            'Entry should be in created state'
        );

        _entry.state = PUBLISHED_STATE;

        entries[msg.sender][_id] = _entry;

        emit EntryPublished(_entry);
    }

    event EntryPublished(Entry entry);

    function removeEntry(string memory _id) public hasAccess {
        require(entryExist[msg.sender][_id], 'Entry does not exists');

        Entry memory _entry = entries[msg.sender][_id];

        //compare state strings
        require(
            keccak256(abi.encodePacked(_entry.state)) ==
                keccak256(abi.encodePacked(CREATED_STATE)),
            'Entry should be in created state'
        );

        entryExist[msg.sender][_id] = false;

        emit EntryRemoved(_entry);
    }

    event EntryRemoved(Entry entry);
}
