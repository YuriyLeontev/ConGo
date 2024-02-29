import json
import zipfile
import sqlite3
import datetime
import sys
import re

import psycopg2

DATA_PATH = './data.zip'


def main():
    start = datetime.datetime.now()

    conn = psycopg2.connect(dbname="congo", host="0.0.0.0", user="postgres", password="postgres", port="5432")
    print("Подключение установлено")
    cursor = conn.cursor()

    create_tables(cursor)

    with zipfile.ZipFile(DATA_PATH, 'r') as zip_ref:
        for filename in zip_ref.namelist():
            if re.match(r'accounts_\d+.json$', filename):
                content_str = zip_ref.read(filename).decode('utf-8')
                content_data = json.loads(content_str).get('accounts', [])

                load_accounts_data(content_data, cursor, conn)

    finish = datetime.datetime.now()
    delta = finish - start
    print(delta.total_seconds())


def create_tables(cursor):

    cursor.execute("""CREATE TABLE accounts (
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

                                    CHECK (birth  > '1950-01-01 00:00' AND birth < '2005-01-01 00:00'),  /* Добавить ограничения от 1950 до 2005 */
                                    CHECK (joined > '2011-01-01 00:00' AND joined < '2018-01-01 00:00'), /* Добавить ограничения от 2011 до 2018 */
                                    CHECK (sex = 'm' OR sex = 'f')
                                )
            	               """)


    cursor.execute("""CREATE TABLE accounts_interest (
                                    id             serial       not null unique,
                                    account_id     int references accounts (id) not null,
                                    interest       varchar(255) not null
                                )
            	               """)

    
    cursor.execute("""CREATE TABLE accounts_like (
                                    id           serial       not null unique,
                                    user_id      int references accounts (id) not null,
                                    account_id   int references accounts (id) not null,
                                    ts           timestamp not null
                                    )
            	               """)    
    
    cursor.execute("""CREATE TABLE accounts_premium (
                                    id             serial       not null unique,
                                    user_id        int references accounts (id) not null,
                                    start_premium  timestamp not null,
                                    stop_premium   timestamp not null,
                                    CHECK (start_premium > '2018-01-01 00:00'),
                                    CHECK (stop_premium  > '2018-01-01 00:00')
                                    )
                	               """)    


def load_accounts_data(content_data, cursor, conn):
    account_cache = []
    interest_cache = []
    like_cache = []
    premium_cache = []

    for obj in content_data:
        birth  = datetime.datetime.fromtimestamp(obj.get('birth'))
        joined = datetime.datetime.fromtimestamp(obj.get('joined'))

        account_cache.append((obj.get('email'), obj.get('fname'), obj.get('sname'), obj.get('phone'),
                              obj.get('sex'), birth, joined,
                              obj.get('country'), obj.get('city'), obj.get('status')
                             ))

        for interest in obj.get('interests', []):
            interest_cache.append((obj.get('id'), interest))

        for like in obj.get('likes', []):
            ts = datetime.datetime.fromtimestamp(like.get('ts'))
            like_cache.append((obj.get('id'), like.get('id'), ts))

        premium = obj.get('premium', {})
        if premium:
            start  = datetime.datetime.fromtimestamp(premium.get('start'))
            finish = datetime.datetime.fromtimestamp(premium.get('finish'))            
            premium_cache.append((obj.get('id'), start, finish))

        if len(account_cache) == 10000:
            load_slice(account_cache, interest_cache, like_cache, premium_cache, cursor, conn)

            account_cache = []
            interest_cache = []
            like_cache = []
            premium_cache = []


def load_slice(account_cache, interest_cache, like_cache, premium_cache, cursor, conn):
    try:
        cursor.executemany("""INSERT INTO accounts
                            (email, fname, sname, phone, sex, birth, joined, country, city, status_user)
                            VALUES (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)""", account_cache)
        conn.commit()
        cursor.executemany("INSERT INTO accounts_interest (account_id, interest) VALUES (%s,%s)", interest_cache)
        cursor.executemany("INSERT INTO accounts_like (user_id, account_id, ts)  VALUES (%s,%s,%s)", like_cache)
        cursor.executemany("INSERT INTO accounts_premium (user_id, start_premium, stop_premium) VALUES (%s,%s,%s)", premium_cache)
        conn.commit()
    except sqlite3.Error:
        sys.stdout.write('E')
        sys.stdout.flush()
    sys.stdout.write('+')
    sys.stdout.flush()


if __name__ == '__main__':
    main()
