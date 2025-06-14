#!/bin/bash

DB_URL="postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
MIGRATIONS_DIR="../migrations"

migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" up