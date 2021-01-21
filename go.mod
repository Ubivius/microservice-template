module example

go 1.15

replace handlers => ./handlers

require (
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20201229214741-2366c2514674 // indirect
	github.com/gorilla/mux v1.8.0
	handlers v1.0.0
)
