FROM alpine:3.15

RUN apk add --no-cache ca-certificates

COPY skprmail /usr/local/bin/
RUN chmod +x /usr/local/bin/skprmail

CMD ["/usr/local/bin/skprmail"]
