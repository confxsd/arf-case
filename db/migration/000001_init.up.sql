CREATE TABLE users (
    id UUID NOT NULL DEFAULT (uuid_generate_v4()),
    username text NOT NULL,
    password text NOT NULL
);