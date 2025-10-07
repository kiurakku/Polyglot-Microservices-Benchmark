FROM rust:1.76-alpine AS build
WORKDIR /app
RUN apk add --no-cache musl-dev
COPY Cargo.toml .
COPY src ./src
RUN cargo build --release

FROM alpine:3.20
WORKDIR /srv
COPY --from=build /app/target/release/compute /usr/local/bin/compute
ENTRYPOINT ["/usr/local/bin/compute"]

