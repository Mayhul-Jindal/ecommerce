# builder stage
FROM golang:1.21.0-alpine3.18 as builder

WORKDIR /app

# these layers can be reused because of caching mechanism
COPY go.mod go.sum ./
RUN go mod download
RUN apk add curl
RUN  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

COPY . .
RUN go build -o main .
RUN chmod +x main

# Final stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 /bin/migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY database/migrations ./database/migrations

EXPOSE 3000

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]