CREATE TABLE
    accounts(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        name VARCHAR(256) NOT NULL,
        level VARCHAR(64) NOT NULL,
        type VARCHAR(64)
    );

CREATE TABLE
    business(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        name VARCHAR(256) NOT NULL,
        social_denomination VARCHAR(256) NOT NULL,
        registered_office VARCHAR(256) NOT NULL,
        cif VARCHAR(9) NOT NULL,
        logo MEDIUMBLOB NOT NULL,
        eth_address VARCHAR(42) NOT NULL
    );

CREATE TABLE
    entries(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        account_id VARCHAR(64) NOT NULL,
        multiple_entry_id VARCHAR(64) NOT NULL,
        currency VARCHAR(3) NOT NULL,
        concept VARCHAR(256),
        date VARCHAR(64) NOT NULL,
        quarter VARCHAR(1) NOT NULL,
        debit_balance VARCHAR(64) NOT NULL,
        credit_balance VARCHAR(64) NOT NULL,
        total_balance VARCHAR(64) NOT NULL,
        state VARCHAR(64) NOT NULL,
        state_reason VARCHAR(256),
        creation_date VARCHAR(64) NOT NULL,
        last_update VARCHAR(64) NOT NULL,
        FOREIGN KEY (account_id) REFERENCES accounts(id)
    );
