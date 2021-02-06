# BUILD_TYPE can have these values: local, test or prod
# if BUILD_TYPE is empty, set to local
ARG BUILD_TYPE=local

FROM golang:stretch as build-env

COPY . ./src

RUN apt update
WORKDIR /go/src
echo "Setup build environnement"
mkdir bin
export PATH=$PATH:/go/bin
export GO111MODULE=on
go mod init
echo "Building Microsevice..."
go build -o /go/bin/app -v ./...
echo "First Docker build-stage is now done"


FROM gcr.io/distroless/base as prod

FROM golang:alpine as test


FROM golang:stretch as local


FROM ${BUILD_TYPE} AS exit_artefact
COPY --from=build-env /go/bin/app /app
CMD ["app"]