
CREATE TABLE accounts
(
    id            serial       not null unique,
    email         varchar(100) not null unique,
    fname         varchar(50),
    sname         varchar(50),
    phone         varchar(16)  unique,
    sex           varchar(1)   not null,
    birth         timestamp    not null,    
    joined        timestamp    not null,    
    country       varchar(50),
    city          varchar(50),
    status_user   varchar(50), 

    CHECK (birth > '1950-01-01 00:00' AND birth < '2005-01-01 00:00'), /* Добавить ограничения от 1950 до 2005 */
    CHECK (joned > '2011-01-01 00:00' AND joned < '2018-01-01 00:00'), /* Добавить ограничения от 2011 до 2018 */
    CHECK (sex = 'm' OR sex = 'f')
);

CREATE TABLE accounts_interest
(
    id             serial       not null unique,
    account_id     int references accounts (id) not null,
    interest       varchar(255) not null
);


CREATE TABLE accounts_like
(
    id           serial       not null unique,
    user_id      int references accounts (id) not null,
    account_id   int references accounts (id) not null,
    ts           timestamp not null
);

CREATE TABLE accounts_premium
(
    id             serial       not null unique,
    user_id        int references accounts (id) not null,
    start_premium  timestamp not null,
    stop_premium   timestamp not null,
    CHECK (start_premium > '2018-01-01 00:00'),
    CHECK (stop_premium  > '2018-01-01 00:00')
);
