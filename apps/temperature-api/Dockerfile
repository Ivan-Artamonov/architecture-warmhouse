# Build stage
FROM golang:1.22-alpine  AS builder

# Set working directory
WORKDIR /app

COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/temperature-sensor

# Runtime stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/temperature-sensor .

# Expose the port the app runs on
EXPOSE 80

# Command to run the application
CMD ["./temperature-sensor"]