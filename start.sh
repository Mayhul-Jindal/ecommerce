#!/bin/sh

set -e 

echo "run database migration"
migrate -path /app/database/migrations -database "$DB_URL" -verbose up

echo "start the app"
exec "$@"
