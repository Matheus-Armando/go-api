FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy source code
COPY . .

# If go.mod exists, get dependencies, otherwise initialize
RUN if [ -f go.mod ]; then \
  go get github.com/gin-gonic/gin && \
  go get github.com/joho/godotenv && \
  go mod tidy; \
  else \
  go mod init github.com/Matheus-Armando/go-api && \
  go get github.com/gin-gonic/gin && \
  go get github.com/joho/godotenv && \
  go mod tidy; \
  fi

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