FROM alpine:latest
MAINTAINER Jorge Alarcon
# go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
ADD web_service /usr/bin/web_service

WORKDIR app/
ADD setup.sql .
ENTRYPOINT ["web_service"]
