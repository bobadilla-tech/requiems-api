#!/bin/sh
# Runs automatically on first init of a fresh db_data volume (official
# postgres image convention: everything in /docker-entrypoint-initdb.d/ runs
# once, only against a brand-new data directory). Creates a second database
# so Go's integration tests (which create self-contained
# api_keys/subscriptions/plans/usage_logs fixture tables — see agents.md) can
# run against TEST_DATABASE_URL without colliding with Rails' real migrations
# of those same table names in the main database.
#
# An existing dev volume from before this file was added won't get this
# automatically — run once manually:
#   docker exec requiem-dev-db-1 createdb -U "$POSTGRES_USER" requiem_test
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	SELECT 'CREATE DATABASE requiem_test OWNER ' || quote_ident('$POSTGRES_USER')
	WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'requiem_test')\gexec
EOSQL
