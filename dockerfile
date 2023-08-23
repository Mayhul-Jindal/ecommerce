# TODO
# Make it simpler get rid of this cgo shit dunno what it is 
# Refer to percy and the python docker blog (by snyk)

FROM golang:1.21 AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final stage
FROM alpine:latest

# Install required certificates for Go
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Set the entry point to the binary
ENTRYPOINT ["./app"]