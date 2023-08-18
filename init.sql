CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    ID uuid DEFAULT uuid_generate_v4 (),
    NAME VARCHAR(250) NOT NULL,
    CPF VARCHAR NOT NULL,
    EMAIL VARCHAR NOT NULL,
    PASSWORD VARCHAR NOT NULL,
    NOTIFICATION BOOLEAN NOT null,
    CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (ID)
);