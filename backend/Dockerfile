# Stage 1: Build
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go.mod and go.sum from backend folder into /app in container
COPY backend/go.mod backend/go.sum ./

RUN go mod download

# Copy the entire backend folder contents into /app
COPY backend/. .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app

# Stage 2: Run
FROM alpine:latest

RUN apk add --no-cache ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/app /app/app

EXPOSE 8080

ENTRYPOINT ["/app/app"]
