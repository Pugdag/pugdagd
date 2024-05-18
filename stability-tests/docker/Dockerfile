ARG KASPAD_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPAD_IMAGE} as pugdagd
FROM ${KASPAMINER_IMAGE} as pugdagminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/pugdag/pugdagd/

WORKDIR /go/src/github.com/pugdag/pugdagd/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=pugdagd /app/ /app/
COPY --from=pugdagminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/pugdag/pugdagd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]