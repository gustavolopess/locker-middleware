CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE topics (
        id text PRIMARY KEY default uuid_generate_v4(),
        name text
);

CREATE TABLE rooms (
       id text PRIMARY KEY default uuid_generate_v4(),
       topic_id text NOT NULL,
       FOREIGN KEY (topic_id) REFERENCES topics (id)
);

CREATE SEQUENCE messages_index_seq;

CREATE TABLE messages (
      id text PRIMARY KEY default uuid_generate_v4(),
      room_id text NOT NULL,
      payload text,
      index NOT NULL DEFAULT nextval('messages_index_seq')
);