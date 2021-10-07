#!/bin/bash

docker run -d -p 60000:9090 -v ~/ubivius/microservice-template/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus