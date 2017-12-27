#!/bin/bash

case `uname -s` in 
	Linux*)  sslConfig=/etc/ssl/openssl.cnf;;
esac

openssl req \
	-newkey rsa:2048 \
	-x509 \
	-nodes \
	-keyout server.key \
	-new \
	-out server.pem \
	-subj /CN=localhost \
	-reqexts SAN \
	-extensions SAN \
	-config <(cat $sslConfig <(printf '[SAN]\nsubjectAltName=DNS:localhost')) \
	-sha256 \
	-days 360


