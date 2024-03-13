#!/usr/bin/env python3

import os
import sys
import sqlite3
import mysql.connector

db_user = os.environ.get('DB_USER', 'root')
db_password = os.environ.get('DB_PASSWORD', '123456')
db_host = os.environ.get('DB_HOST', '127.0.0.1')
db_port = os.environ.get('DB_PORT', '3306')
db_name = os.environ.get('DB_NAME', 'railway')

babelcdb_path = os.environ.get('BABELCDB_PATH', './data/BabelCDB/')


def get_last_banlist() -> str:
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    cursor.execute('''
                   SELECT id
                   FROM entity_banlists
                   WHERE ot = 2
                   ORDER BY id DESC;
                   ''')
    result_set = cursor.fetchall()
    con.close()
    return result_set[0][0]


def refresh_top_cards(banlist: str | None):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    if (banlist == None):
        print('Refreshing top cards all...')
        cursor.execute('CALL refresh_mv_top_cards(%s)', (banlist,))
    else:
        print('Refreshing top cards from %s...' % banlist)
        cursor.execute('CALL refresh_mv_top_cards(%s)', (banlist,))
    con.commit()
    con.close()


def main() -> int:
    last_banlist = get_last_banlist()
    refresh_top_cards(None)
    refresh_top_cards(last_banlist)


if __name__ == '__main__':
    sys.exit(main())
