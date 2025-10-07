FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod .
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod CGO_ENABLED=0 GOOS=linux go build -o /bin/api

FROM alpine:3.20
WORKDIR /srv
COPY --from=build /bin/api /usr/local/bin/api
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/api"]

