CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email varchar(128) UNIQUE NOT NULL,
    password varchar(256) NOT NULL,
    username varchar(64) NOT NULL,
    phone varchar(16),
    role varchar(16),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    is_active boolean DEFAULT true
);

CREATE TABLE read_books (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    title varchar(256) NOT NULL,
    author varchar(128) NOT NULL,
    notes text,
    is_finished bool DEFAULT false,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);