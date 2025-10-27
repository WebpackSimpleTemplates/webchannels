#!/bin/bash

docker build . -t webchannels:v1
docker image tag webchannels:v1 harbor.a-7.tech/weebee/webchannels
docker push harbor.a-7.tech/weebee/webchannels
