CREATE TYPE payment_type_enum AS ENUM ('card', 'cash');
CREATE TYPE status_enum AS ENUM ('in_process', 'success', 'cancel');
CREATE TYPE transaction_type_enum AS ENUM ('withdraw', 'topup');
CREATE TYPE source_type_enum AS ENUM ('bonus', 'sales');
CREATE TYPE tarif_type_enum AS ENUM ('percent', 'fixed');
CREATE TYPE staff_type_enum AS ENUM ('shop_assistant', 'cashier');
create type repostitory_transaction_type_enum as enum ('minus', 'plus');

create table categories(
    id varchar(40) primary key not null ,
    name varchar(30),
    parent_id varchar(40) references categories(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

create table products(
    id uuid primary key not null ,
    name varchar(30),
    price int ,
    barcode int,
    category_id timestamp default now(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE branches (
        id uuid PRIMARY KEY NOT NULL,
        name varchar(30),
        address varchar(100),
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP DEFAULT NOW(),
        deleted_at TIMESTAMP DEFAULT NULL
);

create table repositories(
    id uuid primary key not null ,
    product_id uuid references products(id),
    branch_id uuid references branches(id),
    count int,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE sales (
        id UUID PRIMARY KEY NOT NULL,
        branch_id UUID REFERENCES branches (id),
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

create table baskets(
    id uuid primary key ,
    sale_id uuid references sales(id),
    product_id uuid references products(id),
    quantity int,
    price int,
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

CREATE TABLE transactions (
        id uuid PRIMARY KEY NOT NULL,
        sale_id uuid REFERENCES sales (id),
        staff_id uuid REFERENCES staffs (id),
        transaction_type transaction_type_enum,
        source_type source_type_enum,
        from_amount numeric,
        to_amount numeric,
        description text,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP DEFAULT NOW(),
        deleted_at TIMESTAMP DEFAULT NULL
);

create table repository_transactions(
    id uuid primary key ,
    branch_id uuid references branches(id),
    staff_id uuid references staffs(id),
    product_id uuid references products(id),
    repository_transaction_type repostitory_transaction_type_enum,
    price int,
    quantity int
);