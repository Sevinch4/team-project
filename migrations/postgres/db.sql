CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE tarif_type_enum AS ENUM ('percent', 'fixed');

CREATE TYPE staff_type_enum AS ENUM ('shop_assistant', 'cashier');

CREATE TABLE branches(
    id         UUID PRIMARY KEY,
    name       VARCHAR(30),
    address    VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT null
);

CREATE TABLE staff_tarifs (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(30) UNIQUE NOT NULL,
  tarif_type tarif_type_enum NOT NULL,
  amount_for_cash INT,
  amount_for_card INT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE staffs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
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
