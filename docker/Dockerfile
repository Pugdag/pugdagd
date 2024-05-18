# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/pugdag/pugdagd

WORKDIR /go/src/github.com/pugdag/pugdagd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache pugdagd dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o pugdagd .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/pugdag/pugdagd/pugdagd /app/
COPY --from=build /go/src/github.com/pugdag/pugdagd/infrastructure/config/sample-pugdagd.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
