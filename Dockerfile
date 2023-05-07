# Build the application from source
FROM docker.io/library/golang:1.19 AS build-stage

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-world-http

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM registry.access.redhat.com/ubi8/ubi-micro AS build-release-stage

WORKDIR /

COPY --from=build-stage /hello-world-http /hello-world-http

EXPOSE 8080

USER 65532:65532

ENTRYPOINT ["/hello-world-http"]