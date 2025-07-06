# Stage 1: Builder
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
# CGO_ENABLED=0 is important for static linking, which is good for Alpine
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/main .

# Stage 2: Final Image
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the port that the application will listen on
# Cloud Run injects the PORT environment variable
EXPOSE 8080

# Run the application
ENTRYPOINT ["/app/main"]
