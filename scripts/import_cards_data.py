#!/usr/bin/env python3

import os
import sys
import sqlite3

to_import = ['cards.cdb']
db_path = os.path.realpath(os.path.dirname(
    os.path.realpath(__file__)) + '/../ygorec-data.db')


def rename_tables(name):
    con = sqlite3.connect(name)
    con.execute("ALTER TABLE datas RENAME TO entity_cards;")
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


def merge_tables(full_path):
    con = sqlite3.connect(db_path)
    con.execute("ATTACH ? AS toMerge;", (full_path,))
    con.execute(
        "INSERT OR REPLACE INTO entity_cards SELECT * FROM toMerge.entity_cards;")
    con.commit()
    con.execute("DETACH toMerge;")
    con.close()


def main() -> int:
    path = sys.argv[1] if len(sys.argv) > 1 else "BabelCDB/"

    for filename in to_import:
        full_path = os.path.realpath((path + filename))
        rename_tables(full_path)
        merge_tables(full_path)
    return 0


if __name__ == '__main__':
    sys.exit(main())
