#!/bin/bash
docker build --rm -t haskellwebapp .

docker run --detach=false -it --name haskellwebapp haskellwebapp
