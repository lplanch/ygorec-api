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

upsert-data:
	${FETCH_CARDS}

.PHONY: dcb dcuf dcubf dcu dcd goinstall godev goprod gotest goformat upsert-data