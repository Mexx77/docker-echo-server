# Docker HTTP echo server

A simple Dockerfile for mocking http-based web apps. The http echo server
is written in Go and responds with a 200 OK status code to every request.

## Run
```
PORT=8080
docker run \
 -p $PORT:$PORT \
 -e PORT=$PORT \
 --name docker-echo-server \
 --rm \
 mexx77/docker-echo-server
```

## Build locally
```
git clone git@github.com:Mexx77/docker-echo-server.git
cd docker-echo-server
docker build -t docker-echo-server .
```

## Run locally built server
```
PORT=8080
docker run \
 -p $PORT:$PORT \
 -e PORT=$PORT \
 --name docker-echo-server \
 --rm \
 docker-echo-server
```

## Test echo server
```
curl http://localhost:$PORT
```