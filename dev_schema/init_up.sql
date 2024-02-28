CREATE TABLE country
(
    id          serial       not null unique,
    country     varchar(50)  not null
);

CREATE TABLE city
(
    id          serial       not null unique,
    city        varchar(50)  not null,
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
    fname         varchar(50),
    sname         varchar(50),
    phone         varchar(16)  unique,
    sex           varchar(1)   not null,
    birth         timestamp    not null,    
    joned         timestamp    not null,    
    country_id    int references country (id),
    city_id       int references city (id),
    status_id     int references statusUser (id) not null,

    CHECK (birth > '1950-01-01 00:00' AND birth < '2005-01-01 00:00'), /* Добавить ограничения от 1950 до 2005 */
    CHECK (joned > '2011-01-01 00:00' AND joned < '2018-01-01 00:00'), /* Добавить ограничения от 2011 до 2018 */
    CHECK (sex = 'm' OR sex = 'f')
);

CREATE TABLE interests
(
    id             serial       not null unique,
    account_id     int references account (id) not null,
    interest       varchar(255) not null
);

CREATE TABLE interestsUser
(
    id             serial       not null unique,
    interests_id   int references interests (id) not null,
    account_id     int references account (id) not null
);

CREATE TABLE likes
(
    id           serial       not null unique,
    user_id      int references account (id) not null,
    account_id   int references account (id) not null,
    ts           timestamp not null
);

CREATE TABLE premium
(
    id             serial       not null unique,
    user_id        int references account (id) not null,
    start_premium  timestamp not null,
    stop_premium   timestamp not null,
    CHECK (start_premium > '2018-01-01 00:00'),
    CHECK (stop_premium  > '2018-01-01 00:00')
);


INSERT INTO statusUser (nameStatus)
VALUES ('свободны');

INSERT INTO statusUser (nameStatus)
VALUES ('заняты');

INSERT INTO statusUser (nameStatus)
VALUES ('всё сложно');

-- INSERT INTO account (email,fname,sname,phone,sex,birth,joned, status_id)
-- VALUES ('test@test.ru','ivan','ivanovich', '+7 111 111 11 11', 'm', '1951-01-01 00:00', '2012-01-01 00:00', 1 );


-- INSERT INTO country (country) VALUES ('Russia');

-- INSERT INTO city (city, country_id) VALUES ('Moscow',1);

-- INSERT INTO statususer (namestatus) VALUES ('Free');

INSERT INTO account (email, fname, sname, phone, sex, birth, joned, country_id, city_id, status_id) 
VALUES ('esop@ya.ru', 'Kyrva', 'Perdole', '+1124456745', 'm', '1997-04-27 00:00', '2013-04-27 00:00', 1,1,1);