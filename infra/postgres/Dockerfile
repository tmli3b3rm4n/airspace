# Use the official PostgreSQL image for Linux amd64
FROM postgres:latest

# Set environment variables if needed
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

# Copy any additional configuration files if needed
COPY pg_hba.conf /etc/postgresql/pg_hba.conf

# Expose the PostgreSQL port
EXPOSE 5432