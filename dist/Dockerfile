FROM scratch

COPY --from=golang:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY bin/trailrcore /opt/bin/trailrcore

EXPOSE 8080

CMD ["/opt/bin/trailrcore"]

LABEL org.opencontainers.image.source https://github.com/arthureichelberger/trailrcore