#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/products
curl localhost:9090/products/a2181017-5c53-422b-b6bc-036b27c04fc8
curl localhost:9090/products -XPOST -d '{"name":"addName", "price":1.00, "sku":"abc-abc-abcd"}'
curl localhost:9090/products -XPUT -d '{"id":"a2181017-5c53-422b-b6bc-036b27c04fc8", "name":"newName", "price":2.00, "sku":"abc-abc-abcde"}'
curl localhost:9090/products/a2181017-5c53-422b-b6bc-036b27c04fc8 -XDELETE
