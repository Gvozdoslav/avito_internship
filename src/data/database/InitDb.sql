create table if not exists users
(
    Id serial primary key,
    name varchar not null,
    username varchar not null,
    password varchar not null,
    account_id int not null
);

create table if not exists accounts
(
    Id serial primary key,
    Balance decimal default 0
);

create table if not exists transactions
(
    Id serial primary key,
    amount int not null,
    service_id int not null,
    from_id int not null,
    to_id int not null,
    status varchar not null default 'reserved'
);