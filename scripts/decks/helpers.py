import datetime
import json
from enum import Enum
import os
import mysql.connector

db_user = os.environ.get('DB_USER', 'root')
db_password = os.environ.get('DB_PASSWORD', '123456')
db_host = os.environ.get('DB_HOST', '127.0.0.1')
db_port = os.environ.get('DB_PORT', '3306')
db_name = os.environ.get('DB_NAME', 'railway')


def safe_load_json(json_string):
    try:
        return json.loads(json_string)
    except ValueError:
        return []


class DeckType(Enum):
    META = 'meta'
    NON_META = 'non_meta'
    STRUCTURE = 'structure'
    FUN = 'fun'
    ANIME = 'anime'
    MASTER_DUEL = 'master_duel'
    UNKNOWN = 'unknown'

    def to_string(self):
        if self == DeckType.META:
            return 'meta'
        elif self == DeckType.NON_META:
            return 'non_meta'
        elif self == DeckType.STRUCTURE:
            return 'structure'
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
        elif deck_type == 'Structure Decks':
            return DeckType.STRUCTURE
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

    def upsert_in_db(self):
        con = mysql.connector.connect(host=db_host,
                                      port=db_port,
                                      user=db_user,
                                      password=db_password,
                                      database=db_name)
        cursor = con.cursor()
        # INSERT DECK IN DB
        cursor.execute(
            """
            INSERT INTO entity_decks (id, source, deck_type, updated_at)
                VALUES(%s, %s, %s, %s)
            ON DUPLICATE KEY UPDATE
                    source=source,
                    deck_type=deck_type,
                    updated_at=updated_at;
            """,
            (self.id, self.source, self.deck_type.to_string(), self.updated_at)
        )
        con.commit()
        # INSERT CARDS IN DB LINKED TO THE DECK
        cursor.execute(
            "DELETE FROM graph_cards_belong_to_decks WHERE deck_id = %s", (self.id,))
        con.commit()
        for card_id in self.main_deck:
            try:
                cursor.execute(
                    """
                    INSERT INTO graph_cards_belong_to_decks(card_id, deck_id, category)
                        VALUES (%s, %s, %s);
                    """, (card_id, self.id, 0)
                )
            except mysql.connector.errors.IntegrityError:
                continue
        for card_id in self.extra_deck:
            try:
                cursor.execute(
                    """
                    INSERT INTO graph_cards_belong_to_decks(card_id, deck_id, category)
                        VALUES (%s, %s, %s);
                    """, (card_id, self.id, 1)
                )
            except mysql.connector.errors.IntegrityError:
                continue
        for card_id in self.side_deck:
            try:
                cursor.execute(
                    """
                    INSERT INTO graph_cards_belong_to_decks(card_id, deck_id, category)
                        VALUES (%s, %s, %s);
                    """, (card_id, self.id, 2)
                )
            except mysql.connector.errors.IntegrityError:
                continue
        con.commit()
        con.close()

    @staticmethod
    def from_ygoprodeck(deck_data: dict):
        return Deck(
            id=str(deck_data['deckNum']),
            source='ygoprodeck.com',
            deck_type=DeckType.from_ygoprodeck(
                deck_data.get('format', 'unknown')),
            main_deck=[int(n) for n in safe_load_json(
                deck_data.get('main_deck', "[]")) if str(n).isdigit()],
            extra_deck=[int(n) for n in safe_load_json(
                deck_data.get('extra_deck', "[]")) if str(n).isdigit()],
            side_deck=[int(n) for n in safe_load_json(
                deck_data.get('side_deck', "[]")) if str(n).isdigit()],
            updated_at=deck_data.get(
                'date_published')
        )


def get_kv_lastupdate(key):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name,
                                  auth_plugin='mysql_native_password')
    cursor = con.cursor()
    cursor.execute(
        "SELECT value FROM key_value_stores WHERE `key` = CONCAT('last_update_', %s) LIMIT 1;",
        (key,)
    )
    result = cursor.fetchall()
    con.commit()
    con.close()
    return datetime.datetime.fromisoformat(result[0][0]) if len(result) > 0 else datetime.datetime.fromisoformat('2016-06-02')


def update_kv_lastupdate(key: str, date: datetime.datetime):
    con = mysql.connector.connect(host=db_host,
                                  port=db_port,
                                  user=db_user,
                                  password=db_password,
                                  database=db_name,
                                  auth_plugin='mysql_native_password')
    cursor = con.cursor()
    cursor.execute(
        "INSERT INTO key_value_stores (`key`, value) VALUES(CONCAT('last_update_', %s), %s) ON DUPLICATE KEY UPDATE value=VALUES(value);",
        (key, date.isoformat())
    )
    con.commit()
    con.close()
