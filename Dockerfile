# Use the official Go image as a build stage
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o bin/main .

# Use a more recent minimal base image for the final stage
FROM debian:bookworm-slim

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/bin/main /app/bin/

# Copy the .env file into the container
COPY .env /app/

# Set the working directory
WORKDIR /app

# Expose port 8000
EXPOSE 8000

# Command to run the executable
ENTRYPOINT ["./bin/main"]