#================================
#== DOCKER ENVIRONMENT
#================================
DOCKER := @docker
COMPOSE := @docker-compose

dcb:
	${COMPOSE} build

dcuf:
ifdef f
	${COMPOSE} up -d --${f}
endif

dcubf:
ifdef f
	${COMPOSE} up -d --build --${f}
endif

dcu:
	${COMPOSE} up -d --build

dcd:
	${COMPOSE} down

#================================
#== GOLANG ENVIRONMENT
#================================
GO := @go

goinstall:
	${GO} get .

godev:
	${GO} run main.go

goprod:
	${GO} build -o main .

gotest:
	${GO} test -v

goformat:
	${GO} fmt ./...

#================================
#== SCRIPTS
#================================

FETCH_CARDS := ./scripts/fetch_cards.sh
FETCH_ENUMS := ./scripts/import_enums_data.py
FETCH_DECKS := ./scripts/fetch_banlists.sh
FETCH_DECKS := ./scripts/fetch_decks.sh

upsert-data:
	${FETCH_CARDS}
	${FETCH_ENUMS}
	${FETCH_BANLISTS}
	${FETCH_DECKS}

.PHONY: dcb dcuf dcubf dcu dcd goinstall godev goprod gotest goformat upsert-data