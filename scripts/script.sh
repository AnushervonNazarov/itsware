#!/bin/bash

# Database connection details
DB_NAME="postgres"
DB_USER="postgres"
DB_PASSWORD="q123"
DB_HOST="localhost"
DB_PORT="5432"

# Export password so psql doesn't prompt for it
export PGPASSWORD="$DB_PASSWORD"

echo "Seeding database..."

# Run the SQL script
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f /home/hp/go/src/itsware/migrations/db.sql

if [ $? -eq 0 ]; then
    echo "Seeding completed successfully!"
else
    echo "Seeding failed!"
fi
