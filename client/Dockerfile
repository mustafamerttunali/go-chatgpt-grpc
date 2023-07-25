FROM golang:alpine as builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o main .
# ------------------------------------------------------
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main /main
ENTRYPOINT ["/main"]