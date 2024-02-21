#!/usr/bin/env bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

if [ -d "$BABELCDB_PATH" ]; then
    # if folder exists, fetch and check if there is changes on master
    echo "BabelCDB found, fetching new data..."
    git -C "$BABELCDB_PATH" checkout master
    # if there is changes or if DATABASE_URI file is not found, pull them and import the db
    git -C "$BABELCDB_PATH" pull
    echo "Importing new card data..."
    go run ${SCRIPT_DIR}/setup_database.go
    mkdir -p /tmp/ygorec_cards_data/ && cp $BABELCDB_PATH/*.cdb /tmp/ygorec_cards_data/
    ${SCRIPT_DIR}/import_cards_data.py /tmp/ygorec_cards_data/
    sqlite3mysql -f /tmp/ygorec_cards_data/cards.cdb -i UPDATE -u "$DB_USER" --mysql-password "$DB_PASSWORD" -h "$DB_HOST" -P "$DB_PORT" -d "$DB_NAME"
    rm -f /tmp/ygorec_cards_data/*.cdb
else
    # folder not found, clone it
    echo "BabelCDB not found, cloning repository..."
    git clone "https://github.com/ProjectIgnis/BabelCDB.git" "$BABELCDB_PATH"
    git -C "$BABELCDB_PATH" checkout master
    # upsert the database
    echo "Repository ready, importing card data..."
    go run ${SCRIPT_DIR}/setup_database.go
    mkdir -p /tmp/ygorec_cards_data/ && cp $BABELCDB_PATH/*.cdb /tmp/ygorec_cards_data/
    ${SCRIPT_DIR}/import_cards_data.py /tmp/ygorec_cards_data/
    sqlite3mysql -f /tmp/ygorec_cards_data/cards.cdb -i UPDATE -u "$DB_USER" --mysql-password "$DB_PASSWORD" -h "$DB_HOST" -P "$DB_PORT" -d "$DB_NAME"
    rm -f /tmp/ygorec_cards_data/*.cdb
fi
