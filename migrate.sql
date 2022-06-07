CREATE TABLE events
(
    id       serial PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    name     VARCHAR(255) NOT NULL,
    date     TIMESTAMP    NOT NULL
);

insert into events(location, name, date) values ('Heiligenhafen', 'Wingsurfing', '2017-03-14')
