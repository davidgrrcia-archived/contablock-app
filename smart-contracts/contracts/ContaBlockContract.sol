// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

contract ContaBlockContract {
    address public owner;
    mapping(address => bool) private userAccess;

    string public constant CREATED_STATE = 'CREATED';
    string public constant PUBLISHED_STATE = 'PUBLISHED';

    struct Entry {
        string id;
        string contractId;
        string userId;
        string accountId;
        string multipleEntryId;
        string currency;
        string concept;
        uint256 date;
        uint256 debitBalance;
        uint256 creditBalance;
        string state;
    }

    mapping(string => Entry) entries;
    mapping(string => bool) entryExist;
    string[] multipleEntryIds;
    mapping(string => bool) multipleEntryIdExist;
    string[] accountIds;
    mapping(string => bool) accountIdExist;

    modifier onlyOwner() {
        require(
            msg.sender == owner,
            'Only the contract owner can perform this action'
        );
        _;
    }

    modifier hasAccess() {
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
        userAccess[user] = true;
    }

    function revokeAccess(address user) public onlyOwner {
        userAccess[user] = false;
    }

    function createEntry(
        string memory _id,
        string memory _contractId,
        string memory _userId,
        string memory _accountId,
        string memory _multipleEntryId,
        string memory _currency,
        string memory _concept,
        uint256 _date,
        uint256 _debitBalance,
        uint256 _creditBalance
    ) public hasAccess {
        require(!entryExist[_id], 'Entry already exists');

        Entry memory _entry = Entry(
            _id,
            _contractId,
            _userId,
            _accountId,
            _multipleEntryId,
            _currency,
            _concept,
            _date,
            _debitBalance,
            _creditBalance,
            CREATED_STATE
        );

        entryExist[_id] = true;
        entries[_id] = _entry;

        if (!multipleEntryIdExist[_multipleEntryId]) {
            multipleEntryIds.push(_multipleEntryId);
            multipleEntryIdExist[_multipleEntryId] = true;
        }

        if (!accountIdExist[_accountId]) {
            accountIds.push(_accountId);
            accountIdExist[_accountId] = true;
        }

        emit EntryCreated(_entry);
    }

    event EntryCreated(Entry entry);

    function getEntryData(
        string memory _id
    ) public view hasAccess returns (Entry memory) {
        return entries[_id];
    }

    function listMultipleEntryIds()
        public
        view
        hasAccess
        returns (string[] memory)
    {
        return multipleEntryIds;
    }

    function listAccountIds() public view hasAccess returns (string[] memory) {
        return accountIds;
    }

    function publishEntry(string memory _id) public hasAccess {
        require(entryExist[_id], 'Entry does not exists');

        Entry memory _entry = entries[_id];

        //compare state strings
        require(
            keccak256(abi.encodePacked(_entry.state)) ==
                keccak256(abi.encodePacked(CREATED_STATE)),
            'Entry should be in created state'
        );

        _entry.state = PUBLISHED_STATE;

        entries[_id] = _entry;

        emit EntryPublished(_entry);
    }

    event EntryPublished(Entry entry);

    function removeEntry(string memory _id) public hasAccess {
        require(entryExist[_id], 'Entry does not exists');

        Entry memory _entry = entries[_id];

        //compare state strings
        require(
            keccak256(abi.encodePacked(_entry.state)) ==
                keccak256(abi.encodePacked(CREATED_STATE)),
            'Entry should be in created state'
        );

        entryExist[_id] = false;

        emit EntryRemoved(_entry);
    }

    event EntryRemoved(Entry entry);
}
