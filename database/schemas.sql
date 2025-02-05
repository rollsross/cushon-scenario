CREATE TABLE users(
    id TEXT PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE account_types(
    id TEXT PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE funds(
    id TEXT PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE accounts(
    id TEXT PRIMARY KEY,
    account_types_id TEXT NOT NULL,
    users_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (account_types_id) REFERENCES account_types(id),
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE accounts_funds(
    balance INTEGER,
    funds_id TEXT NOT NULL,
    accounts_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (funds_id) REFERENCES funds(id),
    FOREIGN KEY (accounts_id) REFERENCES accounts(id),
    PRIMARY KEY (funds_id, accounts_id)
);