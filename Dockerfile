 FROM golang:1.22-alpine as builder

 WORKDIR /app

 COPY go.mod go.sum ./

 RUN go mod tidy

 COPY . .
 RUN go install github.com/swaggo/swag/cmd/swag@latest

 RUN swag init -g cmd/airspace_challenge/main.go

 RUN go build -o airspace_challenge cmd/airspace_challenge/main.go

 FROM alpine:latest

 RUN apk add --no-cache libpq bash curl

 WORKDIR /root/

 COPY --from=builder /app/airspace_challenge .
 COPY infra/postgres/National_Security_UAS_Flight_Restrictions.geojson ./

 COPY wait-for-it.sh /usr/local/bin/wait-for-it

 RUN chmod +x /usr/local/bin/wait-for-it

 EXPOSE 8080

 CMD ["wait-for-it", "postgres:5432", "--", "./airspace_challenge"]
