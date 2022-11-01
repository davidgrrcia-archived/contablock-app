CREATE TABLE
    accounts(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        name VARCHAR(256) NOT NULL,
        group INTEGER NOT NULL,
        category VARCHAR(64)
    );

CREATE TABLE
    business(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        name VARCHAR(256) NOT NULL,
        social_denomination VARCHAR(256) NOT NULL,
        cif VARCHAR(9) NOT NULL,
        logo BLOB
    );

CREATE TABLE
    entries(
        id VARCHAR(64) NOT NULL PRIMARY KEY,
        account_id VARCHAR(64) NOT NULL,
        double_entry_code VARCHAR(64) NOT NULL,
        currency VARCHAR(3) NOT NULL,
        concept TEXT NOT NULL,
        date DATETIME NOT NULL,
        quarter INTEGER NOT NULL,
        debit_balance DECIMAL NOT NULL,
        credit_balance DECIMAL NOT NULL,
        last_total_balance DECIMAL NOT NULL,
        state VARCHAR(64) NOT NULL,
        creation_date DATETIME NOT NULL,
        last_update DATETIME NOT NULL,
        FOREIGN KEY (account_id) REFERENCES accounts(id)
    );
