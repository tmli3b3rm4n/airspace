# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o data-loader main.go

# Stage 2: Final Image
FROM ghcr.io/osgeo/gdal:ubuntu-small-3.10.0

# Install required packages
RUN apt-get update && apt-get install -y \
    postgresql-client \
    netcat-openbsd \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/data-loader /usr/local/bin/data-loader

# Copy your GeoJSON data into the container
COPY National_Security_UAS_Flight_Restrictions.geojson /data/National_Security_UAS_Flight_Restrictions.geojson

# Set environment variable for GeoJSON file location
ENV GEOJSON_FILE=/data/National_Security_UAS_Flight_Restrictions.geojson

# Expose the /data directory (optional for mounted volumes)
VOLUME ["/data"]

# Copy wait-for-it script and make it executable
COPY wait-for-it.sh /usr/local/bin/wait-for-it
RUN chmod +x /usr/local/bin/wait-for-it

# Default command
CMD ["wait-for-it", "postgres:5432", "--", "data-loader"]