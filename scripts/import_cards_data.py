#!/usr/bin/env python3

import datetime
import os
import sys
import sqlite3
import subprocess

KV_BABELCDB_LAST_COMMIT = "babelcdb_commit"
to_import = ['cards.cdb']
db_path = '/tmp/ygorec-data.db'
babelcdb_path = '/tmp/BabelCDB/'


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


def update_kv_babelcdb_commit(commit_str):
    con = sqlite3.connect(db_path)
    con.execute(
        "INSERT INTO key_value_stores (key, value) VALUES(:key, :value) ON CONFLICT(key) DO UPDATE SET value=excluded.value;", {
            "key": KV_BABELCDB_LAST_COMMIT, "value": commit_str}
    )
    con.execute(
        "INSERT INTO key_value_stores (key, value) VALUES('babelcdb_version_date', :value) ON CONFLICT(key) DO UPDATE SET value=excluded.value;",
        {"value": datetime.datetime.now(datetime.timezone.utc).isoformat()}
    )
    con.commit()
    con.close()


def git_rev_parse(path) -> str:
    return subprocess.check_output(['git', '-C', path, 'rev-parse', 'HEAD']).decode('ascii').strip()


def main() -> int:
    path = sys.argv[1] if len(sys.argv) > 1 else babelcdb_path

    for filename in to_import:
        full_path = os.path.realpath((path + filename))
        rename_tables(full_path)
        merge_tables(full_path)
    update_kv_babelcdb_commit(git_rev_parse(babelcdb_path))
    return 0


if __name__ == '__main__':
    sys.exit(main())
