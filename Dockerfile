FROM alpine:3.9

# set workdir to root
WORKDIR /root

# update certificates
RUN apk update && apk add tzdata ca-certificates && rm -rf /var/cache/apk/*

# copy binary file
COPY nats-logger /root

ENTRYPOINT [ "/root/nats-logger" ]
