# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 creates a statically linked binary, which is ideal for minimal images
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application listens on (e.g., 8080)
EXPOSE 8080

# Run the compiled application
CMD ["./main"]