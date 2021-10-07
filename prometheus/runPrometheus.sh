#!/bin/bash

docker run -d -p 9090:9090 -v ~/ubivius/microservice-template/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus