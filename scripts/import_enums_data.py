#!/usr/bin/env python3

import collections
import datetime
import os
import subprocess
import sys
import sqlite3
import requests
import mysql.connector

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

db_user = os.environ.get('DB_USER', 'root')
db_password = os.environ.get('DB_PASSWORD', '123456')
db_host = os.environ.get('DB_HOST', '127.0.0.1')
db_port = os.environ.get('DB_PORT', '3306')
db_name = os.environ.get('DB_NAME', 'railway')


def update_enums(parsed_data):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    amount = 0

    for key in parsed_data:
        for kv in parsed_data[key]:
            cursor.execute(
                "INSERT INTO {} (id, value) VALUES(%s, %s) ON DUPLICATE KEY UPDATE value=value;".format(
                    key),
                (kv["id"], kv["value"])
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
            if int_id > 0:
                parsed_data[current_key].append(
                    {'id': int(id, 16), 'value': value.strip()})
    return parsed_data


def update_kv_dataeditorx_commit(commit_str):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    cursor.execute(
        "INSERT INTO key_value_stores (`key`, value) VALUES(%s, %s) ON DUPLICATE KEY UPDATE value=value;",
        (KV_ENUMS_LAST_COMMIT, commit_str)
    )
    cursor.execute(
        "INSERT INTO key_value_stores (`key`, value) VALUES('enums_version_date', %s) ON DUPLICATE KEY UPDATE value=value;",
        (datetime.datetime.now(datetime.timezone.utc).isoformat(),)
    )
    con.commit()
    con.close()


def get_last_commit() -> str:
    return subprocess.check_output(['git', 'ls-remote', git_repo, '|', 'grep', 'refs/heads/master']).decode('ascii').strip().split()[0]


def main() -> int:
    print('Importing enums from the web...')

    parsed_data = get_cards_info()
    update_enums(parsed_data)
    update_kv_dataeditorx_commit(get_last_commit())
    return 0


if __name__ == '__main__':
    sys.exit(main())
