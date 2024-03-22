#!/usr/bin/env python3

import asyncio
import os
from mysql.connector.aio import connect


db_user = os.environ.get('DB_USER', 'root')
db_password = os.environ.get('DB_PASSWORD', '123456')
db_host = os.environ.get('DB_HOST', '127.0.0.1')
db_port = os.environ.get('DB_PORT', '3306')
db_name = os.environ.get('DB_NAME', 'railway')


async def get_all_cards_id() -> list[str]:
    async with await connect(host=db_host,
                             port=int(db_port),
                             user=db_user,
                             password=db_password,
                             database=db_name) as cnx:
        async with await cnx.cursor() as cur:
            await cur.execute('''
                           SELECT id
                           FROM entity_cards
                           WHERE alias = 0;
                           ''')
            result_set = await cur.fetchall()
            return list(map(lambda set: set[0], result_set))


async def generate_archetypes_cards_mv(card_id: str):
    async with await connect(host=db_host,
                             port=int(db_port),
                             user=db_user,
                             password=db_password,
                             database=db_name) as cnx:
        async with await cnx.cursor() as cur:
            await cur.execute('DELETE FROM mv_top_related_cards mv WHERE mv.from_card_id = %s;', (card_id,))
            await cnx.commit()
            await cur.execute('''
                              INSERT INTO mv_top_related_cards
                              SELECT
                                  %s AS from_card_id,
                                  a.card_id AS to_card_id,
                                  COUNT(a.amount) AS deck_amount,
                                  SUM(a.amount) AS card_amount
                              FROM (
                                  SELECT
                                      (CASE WHEN ISNULL(c.id) THEN g.card_id ELSE c.alias END) AS card_id,
                                      COUNT(*) AS amount
                                  FROM graph_cards_belong_to_decks g
                                      LEFT OUTER JOIN entity_cards c ON c.id = g.card_id AND c.alias != 0
                                      WHERE card_id != %s AND EXISTS(
                                          SELECT c_d.deck_id FROM graph_cards_belong_to_decks c_d
                                          LEFT OUTER JOIN entity_cards ec ON ec.alias = c_d.card_id
                                          WHERE (c_d.card_id = %s OR ec.id = %s) AND c_d.deck_id = g.deck_id
                                          GROUP BY c_d.deck_id, c_d.card_id
                                      )
                                      GROUP BY card_id, g.deck_id
                              ) a
                                  GROUP BY a.card_id
                                  ORDER BY deck_amount DESC, card_amount DESC, a.card_id ASC
                                  LIMIT 200;
                              ''',
                              (card_id, card_id, card_id, card_id,))
            await cnx.commit()


async def main() -> int:
    card_ids = await get_all_cards_id()
    print('Refreshing top related cards...')
    print('To refresh: %s' % len(card_ids))

    iterables = [card_ids[n::1] for n in range(1)]
    for selected_card_ids in zip(*iterables):
        print(selected_card_ids)
        tasks = []
        for card_id in selected_card_ids:
            tasks.append(generate_archetypes_cards_mv(card_id))
        await asyncio.gather(*tasks)

    print('\nTop related cards refreshed!')


if __name__ == '__main__':
    asyncio.run(main())
