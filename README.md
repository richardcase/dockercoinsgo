[![Build Status](https://travis-ci.org/richardcase/dockercoinsgo.svg?branch=master)](https://travis-ci.org/richardcase/dockercoinsgo)

# Docker Coins using Go

A Go port of https://github.com/jpetazzo/orchestration-workshop/tree/master/dockercoins


# Notes

Certificates generated using EasyRSA as detailed here: http://www.hydrogen18.com/blog/your-own-pki-tls-golang.html

# Running locally

HASHER
$GOPATH/bin/hasher -cert=$(pwd)/../../certs/localhost:50051.crt -key=$(pwd)/../../certs/localhost:50051.key -ca=$(pwd)/../../certs/ca.crt

RNG
$GOPATH/bin/rng -cert=$(pwd)/../../certs/localhost:50052.crt -key=$(pwd)/../../certs/localhost:50052.key -ca=$(pwd)/../../certs/ca.crt

REDIS
docker run --name some-redis -p 6379:6379 -d redis

WORKER
$GOPATH/bin/worker -cert=$(pwd)/../../certs/client0.crt -key=$(pwd)/../../certs/client0.key -ca=$(pwd)/../../certs/ca.crt

UI
DCKR_WEB_STATIC=$(pwd)/../../webui DCKR_CACHE_ADDR=localhost:6379 $GOPATH/bin/webui

# Perf testing

ab -n 100000 -c 20 -T application/json -p ~/rng.json  https://localhost:500
52/v1/rng

