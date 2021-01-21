module example

go 1.15

replace handlers => ./handlers

require (
	github.com/gorilla/mux v1.8.0
	handlers v1.0.0
)
