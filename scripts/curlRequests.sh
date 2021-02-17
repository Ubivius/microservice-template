#!/bin/bash

curl localhost:9090/products
curl localhost:9090/products -XPOST -d '{"name":"addName", "price":1.00, "sku":"abc-abc-abcd"}'
curl localhost:9090/products/1
curl localhost:9090/products/3
curl localhost:9090/delete/1
curl localhost:9090/products -XPUT -d '{"id":1, "name":"addName", "price":1.00, "sku":"abc-abc-abcd"}'