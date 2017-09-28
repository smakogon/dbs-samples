\set client_min_messages = warning

create extension if not exists "pgcrypto";

\i schema/init-bundle-users.sql