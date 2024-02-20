#!/usr/bin/env python3

import datetime
import re
from helpers import Deck, get_kv_lastupdate, update_kv_lastupdate
import os
import sys
import requests

DECKS_ENDPOINT = 'https://ygoprodeck.com/api/decks/getDecks.php'


def get_date_published(url) -> datetime.datetime:
    response = requests.get(url)
    match = re.findall('"datePublished": "(.*)",', response.text)
    return datetime.datetime.strptime(match[0], "%Y-%m-%d %H:%M:%S")


def fetch_deck(offset, last_update: datetime.datetime):
    decks = []
    response = requests.get(DECKS_ENDPOINT, params={
                            'offset': offset, 'sort': 'Date', 'from': last_update.strftime("%Y-%m-%d")})
    raw_decks = response.json()
    for raw_deck in raw_decks:
        raw_deck['date_published'] = get_date_published(
            'https://ygoprodeck.com/deck/' + raw_deck.get('pretty_url'))
        decks.append(Deck.from_ygoprodeck(raw_deck))
    return decks


def main() -> int:
    path = sys.argv[1] if len(sys.argv) > 1 else os.environ.get(
        'DATABASE_PATH', './data/ygorec-data.db')
    last_update = get_kv_lastupdate(path, 'ygoprodeck.com')
    offset = 0
    while True:
        decks = fetch_deck(offset, last_update)
        for deck in decks:
            deck.upsert_in_db(path)
        print('OFFSET: [' + str(offset) +
              '], DECKS SIZE: [' + str(len(decks)) + ']')
        decks[-1].dump()
        update_kv_lastupdate(path, 'ygoprodeck.com',
                             decks[-1].updated_at - datetime.timedelta(days=1))
        if len(decks) < 20:
            break
        offset += len(decks)
    return 0


if __name__ == '__main__':
    sys.exit(main())
