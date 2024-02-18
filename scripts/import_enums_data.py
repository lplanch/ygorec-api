#!/usr/bin/env python3

import collections
import datetime
import os
import subprocess
import sys
import sqlite3
import requests

KV_ENUMS_LAST_COMMIT = "enums_commit"
git_repo = 'https://github.com/NaimSantos/DataEditorX.git'
cardinfo_uri = 'https://raw.githubusercontent.com/NaimSantos/DataEditorX/master/DataEditorX/data/cardinfo_english.txt'
dict_enums = {
    '##rule': 'enum_rules',
    '##attribute': 'enum_attributes',
    '##level': 'enum_levels',
    '##category': 'enum_categories',
    '##race': 'enum_races',
    '##type': 'enum_types',
    '##setname': 'enum_archetypes',
}


def update_enums(db_path, parsed_data):
    con = sqlite3.connect(db_path)
    amount = 0

    for key in parsed_data:
        for kv in parsed_data[key]:
            con.execute(
                "INSERT INTO %s (id, value) VALUES(:id, :value) ON CONFLICT(id) DO UPDATE SET value=excluded.value;" % key,
                {"id": kv["id"], "value": kv["value"]}
            )
            amount += 1
    con.commit()
    con.close()
    print(str(amount) + ' enums upserted!')


def get_cards_info():
    response = requests.get(cardinfo_uri)
    if response.status_code != 200:
        print('Enums not found online')
        exit(1)

    parsed_data = collections.defaultdict(list)
    current_key = None

    for line in response.content.decode().splitlines():
        if (line == '#end'):
            break
        if line.startswith('##'):
            if line in dict_enums:
                current_key = dict_enums[line]
            else:
                current_key = None
        elif current_key != None:
            [id, value] = line.split(None, 1)
            int_id = int(id, 16)
            if int_id != 0:
                parsed_data[current_key].append(
                    {'id': int(id, 16), 'value': value.strip()})
    return parsed_data


def update_kv_babelcdb_commit(commit_str, db_path):
    con = sqlite3.connect(db_path)
    con.execute(
        "INSERT INTO key_value_stores (key, value) VALUES(:key, :value) ON CONFLICT(key) DO UPDATE SET value=excluded.value;", {
            "key": KV_ENUMS_LAST_COMMIT, "value": commit_str}
    )
    con.execute(
        "INSERT INTO key_value_stores (key, value) VALUES('enums_version_date', :value) ON CONFLICT(key) DO UPDATE SET value=excluded.value;",
        {"value": datetime.datetime.now(datetime.timezone.utc).isoformat()}
    )
    con.commit()
    con.close()


def get_last_commit() -> str:
    return subprocess.check_output(['git', 'ls-remote', git_repo, '|', 'grep', 'refs/heads/master']).decode('ascii').strip().split()[0]


def main() -> int:
    print('Importing enums from the web...')
    path = sys.argv[1] if len(sys.argv) > 1 else os.path.realpath(
        os.path.dirname(os.path.realpath(__file__)) + '/../ygorec-data.db')

    parsed_data = get_cards_info()
    update_enums(path, parsed_data)
    update_kv_babelcdb_commit(get_last_commit(), path)
    return 0


if __name__ == '__main__':
    sys.exit(main())
