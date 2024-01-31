CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE tarif_type_enum AS ENUM ('percent', 'fixed');

CREATE TYPE staff_type_enum AS ENUM ('shop_assistant', 'cashier');

CREATE TYPE payment_type_enum AS ENUM ('card', 'cash');

CREATE TYPE storage_trunsaction_type_enum AS ENUM ('minus', 'plus');

CREATE TYPE transaction_type_enum AS ENUM ('withdraw', 'topup');

CREATE TABLE branches(
    id         UUID PRIMARY KEY,
    name       VARCHAR(30),
    address    VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT null
);

CREATE TABLE categories(
    id UUID PRIMARY KEY,
    name varchar(30),
    parent_id UUID REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE products(
    id UUID PRIMARY KEY,
    name varchar(30),
    price int ,
    barcode int,
    category_id UUID references categories(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE repositories (
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES products(id),
    branch_id UUID REFERENCES branches(id),
    count INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE transactions (
    id uuid PRIMARY KEY NOT NULL,
    sale_id uuid REFERENCES sales (id),
    staff_id uuid REFERENCES staffs (id),
    transaction_type transaction_type_enum,
    source_type_enum,
    amount numeric,
    description text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE staff_tarifs (
  id UUID PRIMARY KEY,
  name VARCHAR(30) UNIQUE NOT NULL,
  tarif_type tarif_type_enum NOT NULL,
  amount_for_cash INT,
  amount_for_card INT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE staffs (
    id UUID PRIMARY KEY,
    branch_id UUID REFERENCES branches(id),
    tariff_id UUID REFERENCES staff_tarifs(id),
    staff_type staff_type_enum NOT NULL,
    name VARCHAR(30),
    balance INT DEFAULT 0,
    age INT,
    birth_date DATE,
    login VARCHAR(15),
    password VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE sales (
    id UUID PRIMARY KEY NOT NULL,
    branch_id UUID REFERENCES branches(id),
    shop_assistant_id varchar(80),
    cashier_id UUID,
    payment_type payment_type_enum,
    price numeric,
    status status_enum,
    client_name varchar(30),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE baskets (
    id UUID PRIMARY KEY NOT NULL,
    sale_id UUID REFERENCES sales(id),
    product_id UUID REFERENCES products(id),
    quantity int DEFAULT 0,
    price int DEFAULT 0, 
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE repository_transactions (
    id UUID PRIMARY KEY NOT NULL,
    staff_id UUID REFERENCES staffs(id),
    product_id UUID REFERENCES products(id),
    storage_transaction_type storage_transaction_type_enum,
    price int, 
    quantity int DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);
