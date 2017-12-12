#!/bin/bash
docker build --rm -t golangwebapp .

docker run -p 8080:80 --name golangwebapp golangwebapp
