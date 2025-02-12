CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    username VARCHAR(20) NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY NOT NULL,
    author INTEGER REFERENCES users (id),
    title VARCHAR(255) NOT NULL,
    text VARCHAR(2000) NOT NULL,
    isCommented boolean NOT NULL
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY NOT NULL,
    author INTEGER REFERENCES users (id) NOT NULL,
    text VARCHAR(200) NOT NULL,
    post INTEGER REFERENCES posts (id) NOT NULL,
    parent INTEGER REFERENCES comments (id)
);