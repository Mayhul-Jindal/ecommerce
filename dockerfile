FROM golang:1.21 AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application with optimizations
RUN go build  -o bin/book-store-microservice

EXPOSE 3000

CMD ["./bin/book-store-microservice"]