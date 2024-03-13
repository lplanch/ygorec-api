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


def get_ungenerated_decks() -> list[str]:
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    cursor.execute('''
                   SELECT id
                   FROM entity_decks
                   WHERE NOT EXISTS (SELECT * FROM mv_deck_archetypes WHERE deck_id = id);
                   ''')
    result_set = cursor.fetchall()
    con.close()
    return list(map(lambda set: set[0], result_set))


def generate_archetypes_mv(ungenerated_deck_ids: list[str]):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    print('Generating missing deck archetypes...')
    print('To generate: %s' % len(ungenerated_deck_ids))
    for deck_id in ungenerated_deck_ids:
        cursor.execute('CALL refresh_mv_deck_archetype(%s)', (deck_id,))
    print('Missing deck archetypes generated!')
    con.commit()
    con.close()


def main() -> int:
    ungenerated_deck_ids = get_ungenerated_decks()
    generate_archetypes_mv(ungenerated_deck_ids)


if __name__ == '__main__':
    sys.exit(main())
