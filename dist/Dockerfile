FROM golang:latest as builder

RUN CGO_ENABLED=0 GOOS=linux go install github.com/jackc/tern@latest

FROM alpine

COPY --from=builder /go/bin/tern /opt/bin/tern
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY migrations /migrations
COPY bin/trailrcore /opt/bin/trailrcore

EXPOSE 8080

CMD ["sh", "-c", "/opt/bin/tern migrate --config /migrations/tern.conf --migrations /migrations && /opt/bin/trailrcore"]

LABEL org.opencontainers.image.source https://github.com/arthureichelberger/trailrcore