CREATE TABLE IF NOT EXISTS category
(
    category_id SERIAL PRIMARY KEY,
    name        VARCHAR(20) NOT NULL,
    enabled     BOOL        NOT NULL DEFAULT TRUE,
    icon        VARCHAR(20),
    type        VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS account
(
    account_id SERIAL PRIMARY KEY,
    first_name VARCHAR(20),
    nickname   VARCHAR(20),
    last_name  VARCHAR(20),
    email      VARCHAR(40),
    phone      VARCHAR(20),
    balance    INTEGER   NOT NULL,
    max_debt   INTEGER   NOT NULL,
    category   INTEGER   NOT NULL,
    enabled    BOOL               DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (category) REFERENCES category (category_id)
        ON DELETE NO ACTION
);


CREATE TABLE IF NOT EXISTS account_option
(
    account INTEGER     NOT NULL,
    key     VARCHAR(40) NOT NULL,
    value   VARCHAR(40) NOT NULL,
    PRIMARY KEY (account, key),
    FOREIGN KEY (account) REFERENCES account (account_id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS unit
(
    unit_id SERIAL PRIMARY KEY,
    name    VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS location
(
    location_id SERIAL PRIMARY KEY,
    name        VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS product_group
(
    product_group_id SERIAL PRIMARY KEY,
    name             VARCHAR(20) NOT NULL,
    parent           INTEGER,
    FOREIGN KEY (parent) REFERENCES product_group (product_group_id)
);

CREATE TABLE IF NOT EXISTS vat
(
    vat_id SERIAL PRIMARY KEY,
    rate   INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS product
(
    product_id    SERIAL PRIMARY KEY,
    name          VARCHAR(40) NOT NULL,
    price         INTEGER     NOT NULL,
    vat           INTEGER     NOT NULL,
    product_group INTEGER     NOT NULL,
    size          INTEGER     NOT NULL,
    unit          INTEGER     NOT NULL,
    category      INTEGER     NOT NULL,
    created_at    TIMESTAMP   NOT NULL DEFAULT NOW(),
    FOREIGN KEY (vat) REFERENCES vat (vat_id),
    FOREIGN KEY (product_group) REFERENCES product_group (product_group_id),
    FOREIGN KEY (unit) REFERENCES unit (unit_id),
    FOREIGN KEY (category) REFERENCES category (category_id)
);

CREATE TABLE IF NOT EXISTS product_visibility
(
    product_visibility_id SERIAL PRIMARY KEY,
    category              INTEGER NOT NULL,
    location              INTEGER NOT NULL,
    product               INTEGER NOT NULL,
    FOREIGN KEY (category) REFERENCES category (category_id),
    FOREIGN KEY (location) REFERENCES location (location_id),
    FOREIGN KEY (product) REFERENCES product (product_id)
);
