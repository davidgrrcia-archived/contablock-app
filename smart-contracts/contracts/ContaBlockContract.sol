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

    mapping(string => Entry) public entries;

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
        entries[_id] = Entry(
            _id,
            _accountId,
            _multipleEntryId,
            _currency,
            _concept,
            _date,
            _debitBalance,
            _creditBalance
        );
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
}
