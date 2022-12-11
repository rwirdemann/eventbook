CREATE TABLE events
(
    id       serial PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    name     VARCHAR(255) NOT NULL,
    date     TIMESTAMP    NOT NULL
);

insert into events(location, name, date)
values ('Heiligenhafen', 'Wingsurfing', '2017-03-14')

CREATE TABLE locations
(
    id   serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

insert into locations(name) values ('Altenteil');
insert into locations(name) values ('Heiligenhafen');
insert into locations(name) values ('Hvide Sande');
insert into locations(name) values ('Nørre Voropør');
insert into locations(name) values ('Gammelmark');
insert into locations(name) values ('Grossenbrode');
insert into locations(name) values ('Klitmøller');
insert into locations(name) values ('Hanstholm');
insert into locations(name) values ('El Medano');
insert into locations(name) values ('Kellenhusen');

alter table events add column distance integer;

alter table events drop column distance;
alter table events add column distance double precision;

alter table events add column maxspeed double precision;
alter table events add column duration double precision;

alter table events add column location_id integer;

select * from events;
select * from locations;