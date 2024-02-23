#!/usr/bin/env python3

import asyncio
import datetime
import os
import sys
import aiohttp
import requests
import mysql.connector

KV_BANLISTS_VERSION_DATE = "banlist_version_date"
banlist_dates_url = 'https://ygoprodeck.com/api/banlist/getBanListDates.php'

banlist_data_url = 'https://ygoprodeck.com/api/banlist/getBanList.php?'

db_user = os.environ.get('DB_USER', 'root')
db_password = os.environ.get('DB_PASSWORD', '123456')
db_host = os.environ.get('DB_HOST', '127.0.0.1')
db_port = os.environ.get('DB_PORT', '3306')
db_name = os.environ.get('DB_NAME', 'railway')


def get_ot_from_string_format(format):
    return ["", "OCG", "TCG"].index(format)


def get_string_from_ot_format(ot):
    return ["", "OCG", "TCG"][ot]


def get_status_from_string_status_text(status_text):
    return ["Banned", "Limited", "Semi-Limited"].index(status_text)


async def fetch(session: aiohttp.ClientSession, url):
    async with session.get(url) as response:
        return await response.json()


async def fetch_specific_banlist(cursor, banlist):
    async with aiohttp.ClientSession() as session:
        url = banlist_data_url + 'date=' + \
            banlist[0] + '&list=' + get_string_from_ot_format(banlist[1])
        response = await fetch(session, url)
        for banned_card in response:
            cursor.execute(
                "INSERT INTO graph_cards_belong_to_banlists (card_id, banlist_id, status) VALUES(%s, %s, %s) ON DUPLICATE KEY UPDATE banlist_id=banlist_id;",
                (banned_card['id'], banlist[0], get_status_from_string_status_text(
                    banned_card['status_text']))
            )


async def upsert_banned_cards():
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    cursor.execute(
        "SELECT * from entity_banlists;"
    )
    banlists = cursor.fetchall()

    tasks = []
    for banlist in banlists:
        print('Fetching banlist from ' + banlist[0] + '...')
        task = asyncio.create_task(fetch_specific_banlist(cursor, banlist))
        tasks.append(task)
    await asyncio.wait(tasks)
    con.commit()
    con.close()


def upsert_banlists(banlist_data):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    amount = 0

    for banlist in banlist_data:
        cursor.execute(
            "INSERT INTO entity_banlists (id, ot) VALUES(%s, %s) ON DUPLICATE KEY UPDATE ot=ot;",
            (banlist['date'], get_ot_from_string_format(banlist['type']))
        )
        amount += 1

    con.commit()
    con.close()
    print(str(amount) + ' banlists upserted!')


def get_banlists_data(formats):
    response = requests.get(banlist_dates_url)
    if response.status_code != 200:
        print('Banlists not found online')
        exit(1)

    data = [bl for bl in response.json() if bl['type'] in formats]
    return data


def update_kv_banlists_version_date():
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name)
    cursor = con.cursor()
    cursor.execute(
        "INSERT INTO key_value_stores (`key`, value) VALUES(%s, %s) ON DUPLICATE KEY UPDATE value=value;",
        (KV_BANLISTS_VERSION_DATE, datetime.datetime.now(
            datetime.timezone.utc).isoformat(),)
    )
    con.commit()
    con.close()


async def main() -> int:
    print('Importing banlists from the web...')

    banlist_data = get_banlists_data({"TCG"})
    upsert_banlists(banlist_data)
    await upsert_banned_cards()
    update_kv_banlists_version_date()
    return 0


if __name__ == '__main__':
    asyncio.run(main())
