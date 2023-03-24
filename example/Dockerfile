FROM alpine:latest
COPY ./curl.sh ./
RUN apk update && apk upgrade && apk add curl && chmod 755 curl.sh
CMD ["./curl.sh"]
