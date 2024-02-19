#!/usr/bin/env python3

import sys
import json
import requests
from enum import Enum


class DeckType(Enum):
    META = 'meta'
    NON_META = 'non_meta'
    FUN = 'fun'
    MASTER_DUEL = 'master_duel'
    UNKNOWN = 'unknown'

    @staticmethod
    def from_ygoprodeck(deck_type: str):
        if deck_type == 'Meta Decks':
            return DeckType.META
        elif deck_type == 'Non-Meta Decks':
            return DeckType.NON_META
        elif deck_type == 'Fun/Casual Decks':
            return DeckType.FUN
        elif deck_type == 'Master Duel Decks':
            return DeckType.MASTER_DUEL
        else:
            print('not found: ', deck_type)
            return DeckType.UNKNOWN


DECKS_ENDPOINT = 'https://ygoprodeck.com/api/decks/getDecks.php'


class Deck:
    def __init__(self, id: str, source: str, deck_type: DeckType, main_deck: list, extra_deck: list, side_deck: list):
        self.id = source + '/' + id
        self.source = source
        self.deck_type = deck_type
        self.main_deck = main_deck
        self.extra_deck = extra_deck
        self.side_deck = side_deck

    def dump(self):
        print("{\n\t\"id\": \"%s\",\n\t\"source\": \"%s\",\n\t\"deck_type\": %s,\n\t\"main_deck\": %a,\n\t\"extra_deck\": %a,\n\t\"side_deck\": %a\n}" % (
            self.id, self.source, self.deck_type, self.main_deck, self.extra_deck, self.side_deck))

    @staticmethod
    def from_ygoprodeck(deck_data: dict):
        return Deck(
            id=str(deck_data['deckNum']),
            source='ygoprodeck.com',
            deck_type=DeckType.from_ygoprodeck(
                deck_data.get('format', 'unknown')),
            main_deck=[int(n) for n in json.loads(
                deck_data.get('main_deck', "[]"))],
            extra_deck=[int(n) for n in json.loads(
                deck_data.get('extra_deck', "[]"))],
            side_deck=[int(n) for n in json.loads(
                deck_data.get('side_deck', "[]"))]
        )


def fetch_deck(offset):
    response = requests.get(DECKS_ENDPOINT, params={'offset': offset})
    data = response.json()
    deck1 = Deck.from_ygoprodeck(data[0])
    deck1.dump()


def main() -> int:
    fetch_deck(0)
    return 0


if __name__ == '__main__':
    sys.exit(main())
