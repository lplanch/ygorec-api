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
  python3 python3-requests \
  sqlite3 libsqlite3-dev
COPY --from=builder . ./usr/src/app
ARG BABELCDB_PATH
ARG DATABASE_PATH
RUN make upsert-data
RUN make goprod
EXPOSE 4000
CMD ["./main"]