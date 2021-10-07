#!/bin/bash

docker build -t my-prometheus .
docker run -d -p 60000:9090 my-prometheus