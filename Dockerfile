FROM golang:1.11 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED 0
RUN cd /build/cmd/grenade && \
    go build && \
    chmod +x grenade

FROM alpine:3.9
COPY --from=builder /build/cmd/grenade/grenade /grenade
ENTRYPOINT ["/grenade"]
