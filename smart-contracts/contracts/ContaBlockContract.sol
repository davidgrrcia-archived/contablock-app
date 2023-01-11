// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

contract ContaBlockContract {
    struct Entry {
        string id;
        string accountId;
        string multipleEntryId;
        string currency;
        string concept;
        string date;
        uint256 debitBalance;
        uint256 creditBalance;
    }

    Entry[] entries;
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
        require(entryExist[_id] == false, 'Entry already exists');

        entries.push(
            Entry(
                _id,
                _accountId,
                _multipleEntryId,
                _currency,
                _concept,
                _date,
                _debitBalance,
                _creditBalance
            )
        );

        entryExist[_id] = true;

        if (!multipleEntryIdExist[_multipleEntryId]) {
            multipleEntryIds.push(_multipleEntryId);
            multipleEntryIdExist[_multipleEntryId] = true;
        }

        if (!accountIdExist[_accountId]) {
            accountIds.push(_accountId);
            accountIdExist[_accountId] = true;
        }

        emit EntryCreated(
            _id,
            _accountId,
            _multipleEntryId,
            _currency,
            _concept,
            _date,
            _debitBalance,
            _creditBalance
        );
    }

    event EntryCreated(
        string id,
        string accountId,
        string multipleEntryId,
        string currency,
        string concept,
        string date,
        uint256 debitBalance,
        uint256 creditBalance
    );

    function listMultipleEntryIds() public view returns (string[] memory) {
        return multipleEntryIds;
    }

    function listAccountIds() public view returns (string[] memory) {
        return accountIds;
    }
}
