#! /bin/sh

URL=host.docker.internal:8088

while true
do
  curl -o /dev/null $URL
done
