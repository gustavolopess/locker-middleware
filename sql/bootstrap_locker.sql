CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE identities (
    id text PRIMARY KEY default uuid_generate_v4(),
    fingerprint text NOT NULL
);

CREATE TABLE lockers (
    id text PRIMARY KEY default uuid_generate_v4()
);