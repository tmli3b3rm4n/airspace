# Base image with PostgreSQL and PostGIS
FROM postgis/postgis:15-3.3

# Install GDAL for loading GeoJSON
RUN apt-get update && apt-get install -y gdal-bin && apt-get clean

# Copy initialization script and GeoJSON file
COPY init.sql /docker-entrypoint-initdb.d/init.sql


# Expose PostgreSQL port
EXPOSE 5432