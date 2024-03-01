#!/usr/bin/env python3

import datetime
import re
import requests
import aiohttp
import asyncio

from helpers import Deck, get_kv_lastupdate, update_kv_lastupdate

DECKS_ENDPOINT = 'https://ygoprodeck.com/api/decks/getDecks.php'


async def fetch(session, url):
    async with session.get(url) as response:
        return await response.text()


async def get_date_published(raw_deck, default_date):
    async with aiohttp.ClientSession() as session:
        response = await fetch(session, 'https://ygoprodeck.com/deck/' + raw_deck.get('pretty_url'))
        match = re.findall('"datePublished": "(.*)",', response)
        print('response of ' + raw_deck.get('pretty_url'))
        raw_deck['date_published'] = datetime.datetime.strptime(
            match[0], "%Y-%m-%d %H:%M:%S") if len(match) > 0 else default_date


async def fetch_deck(offset, last_update: datetime.datetime, default_date: datetime.datetime):
    decks = []
    response = requests.get(DECKS_ENDPOINT, params={
                            'offset': offset, 'sort': 'Date', 'from': last_update.strftime("%Y-%m-%d")})
    raw_decks = response.json()

    tasks = []
    for raw_deck in raw_decks:
        task = asyncio.create_task(get_date_published(raw_deck, default_date))
        tasks.append(task)
    await asyncio.wait(tasks)

    for raw_deck in raw_decks:
        decks.append(Deck.from_ygoprodeck(raw_deck))
    return decks


async def main() -> int:
    last_update = get_kv_lastupdate('ygoprodeck.com')
    default_date = None
    offset = 0
    while True:
        decks = await fetch_deck(offset, last_update, default_date)
        for deck in decks:
            deck.upsert_in_db()
        print('OFFSET: [' + str(offset) +
              '], DECKS SIZE: [' + str(len(decks)) + ']')
        decks[-1].dump()
        default_date = decks[-1].updated_at
        update_kv_lastupdate('ygoprodeck.com',
                             decks[-1].updated_at - datetime.timedelta(days=1))
        if len(decks) < 20:
            break
        offset += len(decks)
    return 0


if __name__ == '__main__':
    asyncio.run(main())
