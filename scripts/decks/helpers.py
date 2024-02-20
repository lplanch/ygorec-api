import datetime
import json
from enum import Enum
import sqlite3


class DeckType(Enum):
    META = 'meta'
    NON_META = 'non_meta'
    FUN = 'fun'
    ANIME = 'anime'
    MASTER_DUEL = 'master_duel'
    UNKNOWN = 'unknown'

    def to_string(self):
        if self == DeckType.META:
            return 'meta'
        elif self == DeckType.NON_META:
            return 'non_meta'
        elif self == DeckType.FUN:
            return 'fun'
        elif self == DeckType.ANIME:
            return 'anime'
        elif self == DeckType.MASTER_DUEL:
            return 'master_duel'
        else:
            return 'unknown'

    @staticmethod
    def from_ygoprodeck(deck_type: str):
        if deck_type in ['Meta Decks', 'Tournament Meta Decks', 'World Championship Decks']:
            return DeckType.META
        elif deck_type in ['Non-Meta Decks', 'Theorycrafting Decks']:
            return DeckType.NON_META
        elif deck_type in ['Fun/Casual Decks']:
            return DeckType.FUN
        elif deck_type in ['Anime Decks']:
            return DeckType.ANIME
        elif deck_type == 'Master Duel Decks':
            return DeckType.MASTER_DUEL
        else:
            print('\n\n\nDECK TYPE NOT FOUND:\n\n\n', deck_type)
            return DeckType.UNKNOWN


class Deck:
    updated_at: datetime.datetime

    def __init__(self, id: str, source: str, deck_type: DeckType, main_deck: list, extra_deck: list, side_deck: list, updated_at: datetime.datetime):
        self.id = source + '/' + id
        self.source = source
        self.deck_type = deck_type
        self.main_deck = main_deck
        self.extra_deck = extra_deck
        self.side_deck = side_deck
        self.updated_at = updated_at

    def dump(self):
        print("{\n\t\"id\": \"%s\",\n\t\"source\": \"%s\",\n\t\"deck_type\": %s,\n\t\"main_deck\": %a,\n\t\"extra_deck\": %a,\n\t\"side_deck\": %a,\n\t\"updated_at\": %s\n}" % (
            self.id, self.source, self.deck_type, self.main_deck, self.extra_deck, self.side_deck, self.updated_at.strftime("%Y-%m-%d")))

    def upsert_in_db(self, db_path):
        con = sqlite3.connect(db_path)
        new_deck = con.execute(
            """
            INSERT INTO entity_decks (id, source, deck_type, updated_at)
                VALUES(:id, :source, :deck_type, :updated_at)
            ON CONFLICT(id) DO UPDATE
                SET
                    source=excluded.source,
                    deck_type=excluded.deck_type,
                    updated_at=excluded.updated_at
                RETURNING id;
            """,
            {"id": self.id, "source": self.source,
                "deck_type": self.deck_type.to_string(), "updated_at": self.updated_at}
        ).fetchall()
        con.commit()
        new_deck_id = new_deck[0][0] if len(new_deck) > 0 else None
        con.close()

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
                deck_data.get('side_deck', "[]"))],
            updated_at=deck_data.get(
                'date_published')
        )


def get_kv_lastupdate(db_path, key):
    con = sqlite3.connect(db_path)
    result = con.execute(
        "SELECT value FROM key_value_stores WHERE key = 'last_update_' || :key",
        {"key": key}
    ).fetchall()
    con.commit()
    con.close()
    return datetime.datetime.fromisoformat(result[0][0]) if len(result) > 0 else datetime.datetime.fromisoformat('2016-06-02')


def update_kv_lastupdate(db_path: str, key: str, date: datetime.datetime):
    con = sqlite3.connect(db_path)
    con.execute(
        "INSERT INTO key_value_stores (key, value) VALUES('last_update_' || :key, :value) ON CONFLICT(key) DO UPDATE SET value=excluded.value;",
        {"key": key, "value": date.isoformat()}
    )
    con.commit()
    con.close()
