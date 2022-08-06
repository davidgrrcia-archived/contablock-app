CREATE TABLE
    business(
        id VARCHAR(64) NOT NULL,
        name VARCHAR(64) NOT NULL,
        short_description VARCHAR(300) NOT NULL,
        long_description TEXT NOT NULL,
        logo TEXT NOT NULL,
        stock_price DECIMAL,
        url_website VARCHAR(256),
        ceo_name VARCHAR(64),
        number_of_employees INTEGER,
        country VARCHAR(64),
        nif VARCHAR(64) NOT NULL,
        fiscal_name VARCHAR(64) NOT NULL,
        address TEXT NOT NULL,
        phone VARCHAR(64) NOT NULL,
        email VARCHAR(64) NOT NULL,
        state VARCHAR(64) NOT NULL,
        creation_date DATETIME NOT NULL,
        last_update DATETIME NOT NULL,
        PRIMARY KEY (id)
    );

CREATE TABLE
    users(
        id VARCHAR(64) NOT NULL,
        business_id VARCHAR(64) REFERENCES business,
        name VARCHAR(64),
        email VARCHAR(64) NOT NULL,
        password VARCHAR(64) NOT NULL,
        phone VARCHAR(64),
        role VARCHAR(64) NOT NULL,
        state VARCHAR(64) NOT NULL,
        creation_date DATETIME NOT NULL,
        last_update DATETIME NOT NULL,
        PRIMARY KEY (id, business_id),
        FOREIGN KEY (business_id) REFERENCES business(id)
    );

CREATE TABLE
    accounts(
        id VARCHAR(64) NOT NULL,
        name VARCHAR(256) NOT NULL,
        type VARCHAR (64) NOT NULL,
        kind VARCHAR(20),
        PRIMARY KEY (id)
    );

CREATE TABLE
    entries(
        id VARCHAR(64) NOT NULL,
        business_id VARCHAR(64) NOT NULL,
        currency VARCHAR(3) NOT NULL,
        description TEXT NOT NULL,
        participant VARCHAR(64),
        file_upload TEXT,
        date DATETIME NOT NULL,
        type VARCHAR (64) NOT NULL,
        debit_balance DECIMAL NOT NULL,
        credit_balance DECIMAL NOT NULL,
        total_balance DECIMAL NOT NULL,
        state VARCHAR(64) NOT NULL,
        creation_date DATETIME NOT NULL,
        last_update DATETIME NOT NULL,
        PRIMARY KEY (id, business_id),
        FOREIGN KEY (business_id) REFERENCES business(id)
    );

CREATE TABLE
    transactions(
        id VARCHAR(64) NOT NULL,
        business_id VARCHAR(64) NOT NULL,
        entry_id VARCHAR(20) NOT NULL,
        account_id VARCHAR(20) NOT NULL,
        type VARCHAR (64) NOT NULL,
        concept TEXT NOT NULL,
        debit_amount DECIMAL NOT NULL,
        credit_amount DECIMAL NOT NULL,
        total_amount DECIMAL NOT NULL,
        state VARCHAR(64) NOT NULL,
        creation_date DATETIME NOT NULL,
        last_update DATETIME NOT NULL,
        PRIMARY KEY (id, business_id),
        FOREIGN KEY (business_id) REFERENCES business(id),
        FOREIGN KEY (entry_id) REFERENCES entries(id),
        FOREIGN KEY (account_id) REFERENCES accounts(id)
    );