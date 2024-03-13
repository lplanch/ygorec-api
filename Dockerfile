# ======================
#  GO FIRST STAGE
# ======================

FROM --platform=linux/amd64 golang:latest as builder
USER ${USER}
WORKDIR /usr/src/app
COPY go.mod \
  go.sum ./
RUN go mod download
COPY . ./
ENV GO111MODULE="on" \
  GOARCH="amd64" \
  GOOS="linux" \
  CGO_ENABLED="1"
RUN apt-get clean \
  && apt-get remove

# ======================
#  GO FINAL STAGE
# ======================

FROM builder
WORKDIR /usr/src/app
RUN apt-get update \
  && apt-get install -y \
  make \
  vim \
  build-essential \
  git \
  python3 python3-pip \
  python3-requests \
  sqlite3 libsqlite3-dev
RUN pip install --break-system-packages \
  sqlite3-to-mysql \
  mysql-connector \
  aiohttp

COPY --from=builder . ./usr/src/app

ARG PORT=4000
ARG BABELCDB_PATH=./data/BabelCDB
ARG DB_HOST=127.0.0.1
ARG DB_PORT=3306
ARG DB_USER=root
ARG DB_PASSWORD=123456
ARG DB_NAME=railway

RUN make goprod
EXPOSE $PORT
CMD ["sh", "-c", "make upsert-data && make generate-views && ./main"]