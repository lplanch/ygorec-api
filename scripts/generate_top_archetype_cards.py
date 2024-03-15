#!/usr/bin/env python3

import os
import sys
import mysql.connector

db_user = os.environ.get('DB_USER', 'root')
db_password = os.environ.get('DB_PASSWORD', '123456')
db_host = os.environ.get('DB_HOST', '127.0.0.1')
db_port = os.environ.get('DB_PORT', '3306')
db_name = os.environ.get('DB_NAME', 'railway')


def get_all_archetypes_id() -> list[str]:
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    cursor.execute('''
                   SELECT id
                   FROM enum_archetypes;
                   ''')
    result_set = cursor.fetchall()
    con.close()
    return list(map(lambda set: set[0], result_set))


def generate_archetypes_cards_mv(archetype_ids: list[str]):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    print('Refreshing archetype cards...')
    print('To refresh: %s' % len(archetype_ids))
    for index, archetype_id in enumerate(archetype_ids):
        if index % round(len(archetype_ids) / 40) == 0:
            print('#', end='', flush=True)
        cursor.execute(
            'CALL refresh_mv_top_archetype_cards(%s)', (archetype_id,))
    print('\nArchetype cards refreshed!')
    con.commit()
    con.close()


def main() -> int:
    archetype_ids = get_all_archetypes_id()
    generate_archetypes_cards_mv(archetype_ids)


if __name__ == '__main__':
    sys.exit(main())
