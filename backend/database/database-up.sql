CREATE TABLE
    balances(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        type VARCHAR(256) NOT NULL,
        kind VARCHAR(256) NULL,
        parent VARCHAR(256) NULL,
        child VARCHAR(256)  NULL,
        accounts VARCHAR(256) NOT NULL
);

CREATE TABLE
    balances_by_year(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        balance_id VARCHAR(64) NOT NULL,
        amount VARCHAR(64) NOT NULL,
        year VARCHAR(4) NOT NULL,
        FOREIGN KEY (balance_id) REFERENCES balances(id)
);

CREATE TABLE
    accounts(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        name VARCHAR(256) NOT NULL,
        level VARCHAR(64) NOT NULL,
        balance_id VARCHAR(64) NULL
    );

CREATE TABLE
    business(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        name VARCHAR(256) NOT NULL,
        social_denomination VARCHAR(256) NOT NULL,
        registered_office VARCHAR(256) NOT NULL,
        cif VARCHAR(9) NOT NULL,
        logo TEXT NOT NULL,
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

LOAD DATA INFILE '/var/lib/mysql-files/accounts.csv'
INTO TABLE accounts
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 0 ROWS;

LOAD DATA INFILE '/var/lib/mysql-files/balances.csv'
INTO TABLE balances
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 0 ROWS;

LOAD DATA INFILE '/var/lib/mysql-files/business.csv'
INTO TABLE business
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 0 ROWS;
