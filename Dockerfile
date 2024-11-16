# Use official Golang image as a build stage
FROM golang:1.22-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o airspace_challenge cmd/airspace_challenge/main.go

# Use the official Golang image for runtime
FROM alpine:latest

# Install necessary libraries (e.g., libpq for PostgreSQL)
RUN apk add --no-cache libpq

# Set the working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/airspace_challenge .
COPY infra/postgres/National_Security_UAS_Flight_Restrictions.geojson ./

# Expose the port your app will run on
EXPOSE 8080

# Command to run the application
CMD ["./airspace_challenge"]