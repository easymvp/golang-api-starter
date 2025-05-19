FROM golang:alpine as builder
WORKDIR /app

# This will download all certificates (ca-certificates) and builds it in a
# single file under /etc/ssl/certs/ca-certificates.crt (update-ca-certificates)
# I also add git so that we can download with `go mod download` and
# tzdata to configure timezone in final image
RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN  GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Golang can run in a scratch image, so that, the only thing that your docker
# image contains is your executable
FROM golang:alpine
ARG ENVIRONMENT
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# This line will copy all certificates to final image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 7788
CMD ["./main"]
