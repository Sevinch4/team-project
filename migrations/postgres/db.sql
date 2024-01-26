create table branches(
    id         uuid primary key not null ,
    name       varchar(30),
    address    varchar(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT null
);

create table sales(
    id                uuid primary key not null ,
    branch_id         uuid references branches (id),
    shop_assistant_id varchar(30),
    cashier_id        uuid,
    payment_type      payment_type_enum,
    price             numeric(1000, 1000),
    status            status_enum,
    client_name       varchar(30),
    created_at        TIMESTAMP DEFAULT NOW(),
    updated_at        TIMESTAMP DEFAULT NOW(),
    deleted_at        TIMESTAMP DEFAULT null
);

create table transactions(
    id uuid primary key not null ,
    sale_id          uuid references sales (id),
    staff_id         uuid references staff (id),
    transaction_type transaction_type_enum,
    source_type      source_type_enum,
    amount           numeric(1000, 100),
    description      text,
    created_at       TIMESTAMP DEFAULT NOW(),
    updated_at       TIMESTAMP DEFAULT NOW(),
    deleted_at       TIMESTAMP DEFAULT null
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
create payment_type_enum as enum ('card', 'cash')

create status_enum as enum ('in_process', 'success', 'cancel')

create transaction_type_enum as enum ('withdraw', 'topup')

create source_type_enum('bonus', 'sales')

CREATE TYPE tarif_type_enum AS ENUM ('percent', 'fixed');

CREATE TYPE staff_type_enum AS ENUM ('shop_assistant', 'cashier');