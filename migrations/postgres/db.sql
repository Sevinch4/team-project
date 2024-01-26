create table branches(
    id         uuid primary key,
    name       varchar(30),
    address    varchar(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT null
);

create table sales(
    id                uuid primary key,
    branch_id         uuid references branches (id),
    shop_assistant_id varchar(30),
    cashier_id        uuid,
    payment_type      payment_type_enum,
    price             numeric(1000, 100),
    status            status_enum,
    client_name       varchar(30),
    created_at        TIMESTAMP DEFAULT NOW(),
    updated_at        TIMESTAMP DEFAULT NOW(),
    deleted_at        TIMESTAMP DEFAULT null
);

create table transactions(
    sale_id          uuid references sales (id),
    staff_id         uuid references staff (id),
    transaction_type transaction_type_enum,
    source_type      source_type_enum,
    amount           int,
    description      text,
    created_at       TIMESTAMP DEFAULT NOW(),
    updated_at       TIMESTAMP DEFAULT NOW(),
    deleted_at       TIMESTAMP DEFAULT null
);


create payment_type_enum as enum ('card', 'cash')

create status_enum as enum ('in_process', 'success', 'cancel')

create transaction_type_enum as enum ('withdraw', 'topup')

create source_type_enum('bonus', 'sales')

