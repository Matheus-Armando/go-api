FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api .

# Use a smaller image for the final build
FROM alpine:latest

WORKDIR /app

# Install json-server
RUN apk add --no-cache nodejs npm && \
  npm install -g json-server

# Copy the built executable
COPY --from=builder /go-api /app/go-api

# Copy the db.json file
COPY data/db.json /app/data/db.json

# Copy any other necessary files
COPY docker/entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh

EXPOSE 8080
EXPOSE 3000

ENTRYPOINT ["/app/entrypoint.sh"]