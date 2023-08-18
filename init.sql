CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    ID uuid DEFAULT uuid_generate_v4 (),
    CPF VARCHAR NOT NULL,
    CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (ID)
);

CREATE TABLE notify (
    ID uuid DEFAULT uuid_generate_v4 (),
    USER_ID uuid,
    datetime time not null,
    MESSAGE TEXT NOT NULL,
    CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (ID),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);
