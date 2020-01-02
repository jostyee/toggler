FROM golang:alpine AS builder

COPY . /src
WORKDIR /src/cmd/toggler
RUN apk --no-cache add build-base ca-certificates git gcc \
    && go get -d -v ./... \
    && CGO_ENABLED=1 CC=gcc go build -o toggler

FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /src/cmd/toggler/toggler .
ENTRYPOINT ["toggler"]
CMD ["http-server", "-port", "8080"]
EXPOSE 8080
