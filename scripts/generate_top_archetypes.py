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


def refresh_top_archetypes():
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    print('Refreshing top archetypes...')
    cursor.execute('CALL refresh_mv_top_archetypes()')
    con.commit()
    con.close()


def main() -> int:
    refresh_top_archetypes()


if __name__ == '__main__':
    sys.exit(main())
