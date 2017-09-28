\echo '>>> initializing bundle - users...'

create table users (
    id          char(36)        not null,
    name        varchar(128)    not null,
    login       varchar(64)     not null,
    password    varchar(128)    not null,
    created_at  timestamptz     not null default now(),
    updated_at  timestamptz     not null default now()
);