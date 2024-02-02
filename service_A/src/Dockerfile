# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY ./controllers ./controllers
COPY ./models ./models

RUN CGO_ENABLED=0 GOOS=linux go build -o /bitcoin-server

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /bitcoin-server /bitcoin-server

EXPOSE 8080

ENTRYPOINT ["/bitcoin-server"]
