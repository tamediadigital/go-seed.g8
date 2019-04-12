FROM golang:alpine
EXPOSE 8081
RUN apk add --no-cache dumb-init
ADD $name$ /usr/local/bin/
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/usr/local/bin/$name$"]
