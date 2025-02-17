# syntax=docker/dockerfile:1
FROM golang:1.20-bullseye AS builder

WORKDIR /src

COPY --link . .

# App build
RUN --mount=type=cache,target=/go/pkg/ \
    --mount=type=cache,target=/root/.cache/go-build/ <<-EOF
    go mod download
    CGO_ENABLED=0 go build -ldflags="-w -s" -o /graphql_demo_app .
EOF

# Final image
FROM debian:bullseye-slim

RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

COPY --link --from=builder /graphql_demo_app /graphql_demo_app
ENTRYPOINT ["/graphql_demo_app"]
