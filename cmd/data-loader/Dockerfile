FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o data-loader main.go

FROM ghcr.io/osgeo/gdal:ubuntu-small-3.10.0

RUN apt-get update && apt-get install -y \
    postgresql-client \
    netcat-openbsd \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/data-loader /usr/local/bin/data-loader

COPY National_Security_UAS_Flight_Restrictions.geojson /data/National_Security_UAS_Flight_Restrictions.geojson
ENV GEOJSON_FILE=/data/National_Security_UAS_Flight_Restrictions.geojson

VOLUME ["/data"]

COPY wait-for-it.sh /usr/local/bin/wait-for-it
RUN chmod +x /usr/local/bin/wait-for-it

# Default command
CMD ["wait-for-it", "postgres:5432", "--", "data-loader"]