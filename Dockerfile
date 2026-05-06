# This image uses alpine linux btw
# Этот образ диска использует alpine linux
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY ./ ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-app .

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/gin-app .

# Disable debug features
# ENV GIN_MODE=release

# Copy the static files
COPY ./static ./static

# Expose the application port
EXPOSE 3257

# Command to run the application
CMD ["./gin-app"]
