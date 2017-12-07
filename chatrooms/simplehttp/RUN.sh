#!/bin/bash

set -x


# Kill any connection for postgresql container to connect
sudo fuser -n tcp -k 5432

# Kill all containers
docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
sudo docker stop $(sudo docker ps -aq) && sudo docker rm -v $(sudo docker ps -aq)

# Set up postgresql container
#go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
rm web_service
CGO_ENABLED=0 go build --tags netgo -a -installsuffix cgo
#go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
sudo docker-compose up -d --build

sleep 5 && sudo docker-compose logs api && sudo docker-compose ps
sudo docker-compose logs -f api
