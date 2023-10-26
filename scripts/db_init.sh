#!/bin/bash
set -e

# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
# 	CREATE DATABASE auth;
# 	GRANT ALL PRIVILEGES ON DATABASE auth TO admin;
# EOSQL

echo "END SETUP"