#! /bin/sh

URL="localhost:8088"

while true
do
  curl -o /dev/null $URL
  sleep 1
done
