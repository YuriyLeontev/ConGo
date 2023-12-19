CREATE TABLE country
(
    id          serial       not null unique,
    country     varchar(255) not null
);

CREATE TABLE city
(
    id          serial       not null unique,
    city        varchar(255) not null,
    country_id    int references country (id) on delete cascade not null
);

CREATE TABLE statusUser
(
    id             serial       not null unique,
    nameStatus     varchar(255) not null
);


CREATE TABLE account
(
    id            serial       not null unique,
    email         varchar(100) not null unique,
    fname         varchar(50)  not null,
    sname         varchar(50)  not null,
    phone         varchar(16)  unique,
    sex           varchar(1)   not null,
    birth         timestamp    not null,    /* Добавить ограничения от 1950 до 2005 */
    joned         timestamp    not null,    /* Добавить ограничения от 2011 до 2018 */
    country_id    int references country (id) on delete cascade,
    city_id       int references city (id) on delete cascade,
    status_id     int references statusUser (id) not null
);