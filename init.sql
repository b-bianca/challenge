CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    ID uuid DEFAULT uuid_generate_v4 (),
    CPF VARCHAR NOT NULL,
    NOTIFICATION BOOL NOT NULL,
    CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (ID)
);

CREATE TABLE notify (
    ID uuid DEFAULT uuid_generate_v4 (),
    USER_ID uuid,
    DATE_TIME  timestamptz NOT NULL,
    MESSAGE TEXT NOT NULL,
    CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (ID),
    CONSTRAINT fk_user FOREIGN KEY(USER_ID) REFERENCES users(ID)
);

CREATE TABLE message (
  ID uuid DEFAULT uuid_generate_v4 (),
  MESSAGE TEXT NOT NULL,
  NOTIFY_ID uuid,
  CREATED_AT timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (ID),
  CONSTRAINT fk_message FOREIGN KEY(NOTIFY_ID) REFERENCES notify(ID)
);