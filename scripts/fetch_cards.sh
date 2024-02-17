#!/usr/bin/env bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

if [ -d "/tmp/BabelCDB" ]; then
    # if folder exists, fetch and check if there is changes on master
    echo "BabelCDB found, fetching new data..."
    git -C /tmp/BabelCDB/ checkout master
    git -C /tmp/BabelCDB/ fetch
    if [[ $(git -C /tmp/BabelCDB/ diff HEAD...origin/master) ]] || !([ -d "/tmp/ygorec-data.db" ]); then
        # if there is changes or if /tmp/ygorec-data.db is not found, pull them and import the db
        git -C /tmp/BabelCDB/ pull
        echo "Importing new card data..."
        go run ${SCRIPT_DIR}/setup_database.go
        mkdir -p /tmp/ygorec_cards_data/ && cp /tmp/BabelCDB/*.cdb /tmp/ygorec_cards_data/
        ${SCRIPT_DIR}/import_cards_data.py /tmp/ygorec_cards_data/
        rm -f /tmp/cards_data/*.cdb
    else
        # else exit the script
        echo "Card database up-to-date!"
    fi
else
    # folder not found, clone it
    echo "BabelCDB not found, cloning repository..."
    git clone "https://github.com/ProjectIgnis/BabelCDB.git" /tmp/BabelCDB/
    git -C /tmp/BabelCDB/ checkout master
    # upsert the database
    echo "Repository ready, importing card data..."
    go run ${SCRIPT_DIR}/setup_database.go
    mkdir -p /tmp/ygorec_cards_data/ && cp /tmp/BabelCDB/*.cdb /tmp/ygorec_cards_data/
    ${SCRIPT_DIR}/import_cards_data.py /tmp/ygorec_cards_data/
    rm -f /tmp/ygorec_cards_data/*.cdb
fi
