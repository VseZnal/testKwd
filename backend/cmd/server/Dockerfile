FROM golang:1.18-alpine AS builder

WORKDIR /testKwd

COPY ./backend/ ./backend

RUN cd backend && go mod download
RUN cd ./backend/cmd/server/ && go build -o /compiled

FROM alpine:3.16 AS runtime

WORKDIR /app

COPY --from=builder /compiled /app

ENTRYPOINT ["/app/compiled"]
