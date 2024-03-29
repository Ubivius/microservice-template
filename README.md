# Microservice-Template
Template for microservices.

This template can be used to create another microservice without having to write any of the boilerplate code.

## Product endpoints

`GET` `/products` Returns json data about every product.

`GET` `/products/{id}` Returns json data about a specific product. `id=[string]`

`GET` `/health/live` Returns a Status OK when live.

`GET` `/health/ready` Returns a Status OK when ready or an error when dependencies are not available.

`POST` `/products` Add new product with specific data.</br>
__Data Params__
```json
{
  "name":        "string, required",
  "sku":         "string, required",
  "description": "string",
  "price":       "float",
}
```

`PUT` `/products` Update product data</br>
__Data Params__
```json
{
  "id":          "string, required",
  "name":        "string",
  "sku":         "string",
  "description": "string",
  "price":       "float",
}
```

`DELETE` `/products/{id}` Delete product.  `id=[string]`

__Fuzzing__
To start fuzzing, simply build and run:
```console
$ cd pkg/fuzzing
$ go-fuzz-build
$ go-fuzz
```
This will generate and test inputs in an infinite loop.
You need at least an initial input named "0" inside the `corpus` folder.
For more information, refer to [go-fuzz](https://github.com/dvyukov/go-fuzz/tree/master) documentation.
