// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

contract ContaBlockContract {
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

    mapping(string => Entry) public entries;
    mapping(string => bool) public entryExist;
    string[] multipleEntryIds;
    mapping(string => bool) multipleEntryIdExist;
    string[] accountIds;
    mapping(string => bool) accountIdExist;

    function createEntry(
        string memory _id,
        string memory _accountId,
        string memory _multipleEntryId,
        string memory _currency,
        string memory _concept,
        string memory _date,
        uint256 _debitBalance,
        uint256 _creditBalance
    ) public {
        require(!entryExist[_id], 'Entry already exists');

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
    ) public view returns (Entry memory) {
        return entries[_id];
    }

    function listMultipleEntryIds() public view returns (string[] memory) {
        return multipleEntryIds;
    }

    function listAccountIds() public view returns (string[] memory) {
        return accountIds;
    }

    function publishEntry(string memory _id) public {
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
}
