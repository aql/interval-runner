FROM golang:1.18.0-bullseye AS builder

RUN mkdir -p /src
COPY ./src /src

WORKDIR /src
RUN go build runner.go

FROM debian:11.2-slim

# Install runner
RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /src/runner /app/runner
RUN chmod +x /app/runner

ENTRYPOINT ["/app/runner"]
