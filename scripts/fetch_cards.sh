#!/usr/bin/env bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

if [ -d "${SCRIPT_DIR}/BabelCDB" ]; then
    # if folder exists, fetch and check if there is changes on master
    echo "BabelCDB found, fetching new data..."
    git -C ${SCRIPT_DIR}/BabelCDB/ checkout master
    git -C ${SCRIPT_DIR}/BabelCDB/ fetch
    if [[ $(git -C ${SCRIPT_DIR}/BabelCDB/ diff HEAD...origin/master) ]] || !([ -d "${SCRIPT_DIR}/../ygorec-data.db" ]); then
        # if there is changes or if ygorec-data.db is not found, pull them and import the db
        git -C ${SCRIPT_DIR}/BabelCDB/ pull
        echo "Importing new card data..."
        go run ${SCRIPT_DIR}/setup_database.go
        mkdir -p /tmp/ygorec_cards_data/ && cp ${SCRIPT_DIR}/BabelCDB/*.cdb /tmp/ygorec_cards_data/
        ${SCRIPT_DIR}/import_cards_data.py /tmp/ygorec_cards_data/
        rm -f /tmp/cards_data/*.cdb
    else
        # else exit the script
        echo "Card database up-to-date!"
    fi
else
    # folder not found, clone it
    echo "BabelCDB not found, cloning repository..."
    git clone "https://github.com/ProjectIgnis/BabelCDB.git" ${SCRIPT_DIR}/BabelCDB/
    git -C ${SCRIPT_DIR}/BabelCDB/ checkout master
    # upsert the database
    echo "Repository ready, importing card data..."
    go run ${SCRIPT_DIR}/setup_database.go
    mkdir -p /tmp/ygorec_cards_data/ && cp ${SCRIPT_DIR}/BabelCDB/*.cdb /tmp/ygorec_cards_data/
    ${SCRIPT_DIR}/import_cards_data.py /tmp/ygorec_cards_data/
    rm -f /tmp/ygorec_cards_data/*.cdb
fi
