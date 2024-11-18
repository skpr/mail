FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY skprmail /usr/local/bin/
RUN chmod +x /usr/local/bin/skprmail

CMD ["/usr/local/bin/skprmail"]
