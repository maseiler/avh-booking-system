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

CREATE TABLE IF NOT EXISTS role
(
    role_id SERIAL PRIMARY KEY,
    name    VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS rights
(
    rights_id  SERIAL PRIMARY KEY,
    role       INTEGER     NOT NULL,
    permission VARCHAR(40) NOT NULL,
    allowed    BOOL        NOT NULL,
    FOREIGN KEY (role) REFERENCES role (role_id)
);

CREATE TABLE IF NOT EXISTS "user"
(
    user_id  SERIAL PRIMARY KEY,
    name     VARCHAR(40) NOT NULL,
    role     INTEGER     NOT NULL,
    password VARCHAR(200) NOT NULL,
    FOREIGN KEY (role) REFERENCES role (role_id)
);

CREATE TABLE IF NOT EXISTS user_option
(
    user_id INTEGER     NOT NULL,
    key     VARCHAR(40) NOT NULL,
    value   VARCHAR(200) NOT NULL,
    PRIMARY KEY (user_id, key),
    FOREIGN KEY (user_id) REFERENCES "user" (user_id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS settings_frontend
(
    key   VARCHAR(40) PRIMARY KEY,
    value VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS settings_payment
(
    key   VARCHAR(40) PRIMARY KEY,
    value VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS settings_email
(
    key   VARCHAR(40) PRIMARY KEY,
    value VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS favorites
(
    product INTEGER NOT NULL,
    account INTEGER NOT NULL,
    count   INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (product, account),
    FOREIGN KEY (product) REFERENCES product (product_id),
    FOREIGN KEY (account) REFERENCES account (account_id)
);

CREATE TABLE IF NOT EXISTS "order"
(
    order_id   SERIAL PRIMARY KEY,
    account    INTEGER   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (account) REFERENCES account (account_id)
);

CREATE TABLE IF NOT EXISTS product_order
(
    order_id      INTEGER NOT NULL,
    product       INTEGER NOT NULL,
    amount        INTEGER NOT NULL,
    product_price INTEGER NOT NULL,
    PRIMARY KEY (order_id, product),
    FOREIGN KEY (order_id) REFERENCES "order" (order_id),
    FOREIGN KEY (product) REFERENCES product (product_id)
);
