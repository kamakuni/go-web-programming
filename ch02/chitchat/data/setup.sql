create table users (
    id  serial primary key,
    uuid    varchar(64) not null unique,
    name    varchar(255),
    emal    varchar(255) not null unique,
    password    varchar(255) not null,
    create_at   timestamp not null
);

create table sessions (
    id  serial primary key,
    uuid    varchar(64) not null unique,
    email   varchar(255),
    user_id integer references users(id),
    create_at   timestamp not null
);

create table threads (
    id  serial primary key,
    uuid    varchar(64) not null unique,
    topic   text,
    user_id integer references users(id),
    create_at   timestamp not null
);

create table posts (
    id  serial primary key,
    uuid    varchar(64) not null unique,
    body    text,
    user_id integer references users(id),
    thread_id   integer references threads(id),
    create_at   timestamp not null
);