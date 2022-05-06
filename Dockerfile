# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint compared to Ubuntu
# Example: FROM golang:1.18-alpine
# FROM golang:1.18-stretch
FROM golang:1.18

WORKDIR /app

# Download necessary Go modules
COPY ./src/go.mod ./
COPY ./src/go.sum ./

RUN go mod download

#COPY ./src/*.go ./

EXPOSE 1323

# Build the Go app
#RUN go build -o app/main .

#CMD ["./app/main"]
