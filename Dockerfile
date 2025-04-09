# --- Builder Stage ---
    FROM golang:1.23.8-alpine AS builder

    # Metadata
    LABEL maintainer="okeny <op58692@gmail.com>"
    
    # Enable Go modules and disable checksum database (optional)
    ENV GO111MODULE=on \
        GOSUMDB=off \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64
    
    # Install required packages
    RUN apk update && apk add --no-cache git
    
    # Set work directory
    WORKDIR /app
    
    # Copy go.mod and go.sum first (to leverage Docker cache)
    COPY go.mod go.sum ./
    
    # Download dependencies
    RUN go mod download
    
    # Copy source code
    COPY . .
    
    # Build the binary
    RUN go build -a -installsuffix cgo -o bms-service .
    
    # --- Final Image Stage ---
    FROM alpine:latest
    
    # Install certs for HTTPS
    RUN apk --no-cache add ca-certificates
    
    # Set work directory
    WORKDIR /app
    
    # Copy binary and .env file
    COPY --from=builder /app/bms-service .
    COPY --from=builder /app/.env .

    # Optional: mark the binary as executable
    RUN chmod +x ./bms-service
    
    # Expose API port
    EXPOSE 8000
    
    # Entrypoint: Run migrations before starting the API service
    ENTRYPOINT ["sh", "-c", "./bms-service migrations up && ./bms-service api"]
    
    