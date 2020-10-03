#!/bin/bash

PWD=$(dirname "$0")
docker-compose -f $PWD/docker-compose.yml up -d --build

while ! docker exec -it gfg_mysql mysql -uuser -ppassword -e "SELECT 1" &> /dev/null ; do
    echo "Waiting for database connection..."
    sleep 1
done

docker exec -it gfg_mysql mysql -uuser -ppassword -e "$(cat $PWD/init_database.sql)"

docker-compose -f $PWD/docker-compose.yml stop

echo Done!