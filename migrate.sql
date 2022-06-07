CREATE TABLE events
(
    id       serial PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    name     VARCHAR(255) NOT NULL,
    date     TIMESTAMP    NOT NULL
);