#!/usr/bin/env python3

import collections
import os
import sys
import sqlite3
import requests

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
            parsed_data[current_key].append(
                {'id': int(id, 16), 'value': value})
    return parsed_data


def main() -> int:
    print('Importing enums from the web...')
    path = sys.argv[1] if len(sys.argv) > 1 else "/../ygorec-data.db"
    db_path = os.path.realpath(os.path.dirname(
        os.path.realpath(__file__)) + path)

    parsed_data = get_cards_info()
    update_enums(db_path, parsed_data)
    return 0


if __name__ == '__main__':
    sys.exit(main())
