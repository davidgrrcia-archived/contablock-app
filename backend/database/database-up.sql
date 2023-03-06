CREATE TABLE balances(
    project VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    type VARCHAR(256) NOT NULL,
    kind VARCHAR(256) NULL,
    parent VARCHAR(256) NULL,
    child VARCHAR(256) NULL,
    accounts VARCHAR(256) NOT NULL,
    PRIMARY KEY (id, project)
);
CREATE TABLE balances_by_year(
    project VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    balance_id VARCHAR(64) NOT NULL,
    amount VARCHAR(64) NOT NULL,
    year VARCHAR(4) NOT NULL,
    PRIMARY KEY (id, project)
);
CREATE TABLE accounts(
    project VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    level VARCHAR(64) NOT NULL,
    balance_id VARCHAR(64) NULL,
    PRIMARY KEY (id, project)
);
CREATE TABLE business(
    project VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    social_denomination VARCHAR(256) NOT NULL,
    registered_office VARCHAR(256) NOT NULL,
    cif VARCHAR(9) NOT NULL,
    logo TEXT NOT NULL,
    PRIMARY KEY (id, project)
);
CREATE TABLE contracts(
    project VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    description TEXT NOT NULL,
    eth_address VARCHAR(42) NOT NULL,
    network VARCHAR(64) NOT NULL,
    admin_id VARCHAR(64) NOT NULL,
    creation_date VARCHAR(64) NOT NULL,
    last_update VARCHAR(64) NOT NULL,
    PRIMARY KEY (id, project)
);
CREATE TABLE entries(
    project VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    contract_id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    account_id VARCHAR(64) NOT NULL,
    multiple_entry_id VARCHAR(64) NOT NULL,
    balance_by_year_id VARCHAR(64) NULL,
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
    UNIQUE KEY (project, id, multiple_entry_id),
    PRIMARY KEY (id, project)
);
CREATE TABLE users(
    project VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    contract_id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    email VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    role VARCHAR(64) NOT NULL,
    eth_address VARCHAR(42) NOT NULL,
    private_key VARCHAR(256) NOT NULL,
    default_password_changed BOOLEAN NOT NULL,
    network VARCHAR(64) NOT NULL,
    business_id VARCHAR(64) NOT NULL,
    creation_date VARCHAR(64) NOT NULL,
    last_update VARCHAR(64) NOT NULL,
    UNIQUE KEY (project, email),
    PRIMARY KEY (id, project)
);