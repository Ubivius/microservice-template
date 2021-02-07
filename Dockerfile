# BUILD_TYPE can have these values: local, test or prod
# if BUILD_TYPE is empty, set to local
ARG BUILD_TYPE=local
FROM golang:stretch as build-env

COPY . ./src
RUN pwd
RUN apt update
RUN mkdir bin
WORKDIR /go/src
RUN pwd
RUN echo "Setup build environnement"
RUN export PATH=$PATH:/go/bin
RUN export GO111MODULE=on
RUN echo "Building Microsevice..."
RUN go build -o /go/bin/app -v ./...
RUN echo "First Docker build-stage is now done"


FROM gcr.io/distroless/base as prod

FROM golang:alpine as test

FROM golang:stretch as local

FROM ${BUILD_TYPE} AS exit_artefact
COPY --from=build-env /go/bin/app /app
CMD ["app"]