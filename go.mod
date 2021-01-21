module example

go 1.15

replace handlers => ./handlers

require (
	github.com/gorilla/mux v1.8.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	handlers v1.0.0
)
