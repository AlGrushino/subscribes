#!/bin/sh
set -e

echo "Checking if migrations are needed..."

if migrate -path ./migrations -database "postgres://$DB_USER:$DB_PASSWORD@db:$DB_PORT/$DB_NAME?sslmode=disable" up 2>/dev/null; then
    echo "Migrations applied successfully"
else
    echo "No new migrations or error occurred"
fi

echo "Starting application..."
exec /main