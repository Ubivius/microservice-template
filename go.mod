module github.com/Ubivius/microservice-template

go 1.15

require (
	github.com/Ubivius/pkg-telemetry v1.0.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	go.mongodb.org/mongo-driver v1.7.3
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.25.0
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.25.0
	go.opentelemetry.io/otel v1.0.1
	golang.org/x/crypto v0.0.0-20210317152858-513c2a44f670 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	sigs.k8s.io/controller-runtime v0.10.2
)
