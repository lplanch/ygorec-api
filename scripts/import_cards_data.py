#!/usr/bin/env python3

import datetime
import os
import sys
import sqlite3
import mysql.connector
import subprocess

KV_BABELCDB_LAST_COMMIT = "babelcdb_commit"
to_import = ['cards.cdb']

db_user = os.environ.get('DB_USER', 'root')
db_password = os.environ.get('DB_PASSWORD', '123456')
db_host = os.environ.get('DB_HOST', '127.0.0.1')
db_port = os.environ.get('DB_PORT', '3306')
db_name = os.environ.get('DB_NAME', 'railway')

babelcdb_path = os.environ.get('BABELCDB_PATH', './data/BabelCDB/')


def rename_tables(name):
    con = sqlite3.connect(name)
    con.execute("ALTER TABLE datas RENAME TO entity_cards;")
    con.execute("ALTER TABLE entity_cards RENAME COLUMN setcode TO set_code")
    con.execute(
        "ALTER TABLE entity_cards ADD COLUMN name TEXT;")
    con.execute(
        "ALTER TABLE entity_cards ADD COLUMN desc TEXT;")
    con.execute(
        "UPDATE entity_cards SET name = (SELECT name FROM texts WHERE id = entity_cards.id);")
    con.execute(
        "UPDATE entity_cards SET desc = (SELECT desc FROM texts WHERE id = entity_cards.id);")
    con.execute(
        "UPDATE entity_cards SET desc = (SELECT desc FROM texts WHERE id = entity_cards.id);")
    con.execute("DROP TABLE texts;")
    con.commit()
    con.close()


def update_kv_babelcdb_commit(commit_str):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    cursor.execute(
        """
            INSERT INTO key_value_stores (`key`, value)
                VALUES (%s, %s)
            ON DUPLICATE KEY UPDATE value=VALUES(value);
        """,
        (KV_BABELCDB_LAST_COMMIT, commit_str)
    )
    cursor.execute(
        "INSERT INTO key_value_stores (`key`, value) VALUES ('babelcdb_version_date', %s) ON DUPLICATE KEY UPDATE value=VALUES(value);",
        (datetime.datetime.now(datetime.timezone.utc).isoformat(),)
    )
    con.commit()
    con.close()


def git_rev_parse(path) -> str:
    return subprocess.check_output(['git', '-C', path, 'rev-parse', 'HEAD']).decode('ascii').strip()


def main() -> int:
    path = sys.argv[1] if len(sys.argv) > 1 else "/tmp/ygorec_cards_data/"

    for filename in to_import:
        full_path = os.path.realpath((path + filename))
        rename_tables(full_path)
    update_kv_babelcdb_commit(git_rev_parse(babelcdb_path))
    return 0


if __name__ == '__main__':
    sys.exit(main())
