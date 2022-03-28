FROM golang:1.18.0-alpine3.15 AS build

WORKDIR /code

# copy depdencies files
COPY go.* .

# download dependencies
RUN go mod download

COPY *.go .

RUN go build -o go-ci

FROM alpine:3.14

WORKDIR /

COPY --from=build /code/go-ci /go-ci

ENTRYPOINT [ "/go-ci" ]
