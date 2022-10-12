CREATE TABLE IF NOT EXISTS manager
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR     NOT NULL UNIQUE,
    password   VARCHAR     NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS store
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR     NOT NULL,
    manager_id INTEGER REFERENCES manager (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS banner
(
    id       SERIAL PRIMARY KEY,
    image    VARCHAR NOT NULL,
    store_id INTEGER REFERENCES store (id)
);

CREATE TABLE IF NOT EXISTS category
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR     NOT NULL,
    image      VARCHAR     NOT NULL,
    store_id   INTEGER REFERENCES store (id),
    is_active  BOOL        NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS product
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR     NOT NULL,
    description TEXT,
    price       INTEGER     NOT NULL,
    image       VARCHAR     NOT NULL,
    is_active   BOOL        NOT NULL DEFAULT true,
    category_id INTEGER REFERENCES category (id),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT price_check CHECK ( price > 0 )
);

CREATE TABLE IF NOT EXISTS customer
(
    id          SERIAL PRIMARY KEY,
    external_id BIGINT      NOT NULL,
    username    VARCHAR,
    phone       VARCHAR,
    first_name  VARCHAR,
    last_name   VARCHAR,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TYPE order_payment_method AS ENUM ('CASH', 'CARD');
CREATE TYPE order_state AS ENUM ('PENDING', 'CANCELED', 'ACCEPTED', 'DELIVERY_IN_PROGRESS', 'DELIVERED');
CREATE TYPE order_delivery_method AS ENUM ('BY_COURIER', 'TAKEAWAY');

CREATE TABLE IF NOT EXISTS orders
(
    id                      SERIAL PRIMARY KEY,
    customer_id             INTEGER REFERENCES customer (id),
    store_id                INTEGER REFERENCES store (id),
    address                 VARCHAR               NOT NULL,
    comment                 TEXT,
    allergies_info          TEXT,
    cancellation_reason     TEXT,
    need_kitchen_appliances BOOL                           DEFAULT false,
    state                   order_state           NOT NULL DEFAULT 'PENDING',
    payment_method          order_payment_method  NOT NULL DEFAULT 'CASH',
    delivery_method         order_delivery_method NOT NULL DEFAULT 'BY_COURIER',
    created_at              TIMESTAMPTZ           NOT NULL DEFAULT now(),
    updated_at              TIMESTAMPTZ           NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS order_item
(
    id         SERIAL PRIMARY KEY,
    order_id   INTEGER REFERENCES orders (id),
    product_id INTEGER REFERENCES product (id),
    amount     INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT amount_check CHECK ( amount > 0 )
);
