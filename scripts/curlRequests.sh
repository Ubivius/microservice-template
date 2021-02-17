#!/bin/bash

curl localhost:9090/products
curl localhost:9090/products/1
curl localhost:9090/products -XPOST --header "Content-Type: application/json" -d '{"name":"addName", "price":1.00, "sku":"abc-abc-abcd"}'
curl localhost:9090/products
curl localhost:9090/delete/1
curl localhost:9090/products
curl localhost:9090/products -XPUT -d '{"id":2, "name":"newName", "price":1.00, "sku":"abc-abc-abcd"}'
