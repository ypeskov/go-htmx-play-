#!/usr/bin/env bash

source ../.env

db_name=${DATABASE_URL:-database.db}

if [ -z "$db_name" ]; then
  echo "No database name specified."
  exit 1
fi

action=$1

if [ -z "$action" ]; then
  echo "No action specified: up or down."
  exit 1
fi

if [ "$action" != "up" ] && [ "$action" != "down" ]; then
  echo "Invalid action: $action."
  echo "Possible values: up, down."
  exit 1
fi

migration_number=$2

db_url="sqlite3://../$db_name"

if [ -z "$migration_number" ]; then
  migrate -database "$db_url" -path migrations $action
else
  migrate -database "$db_url" -path migrations $action $migration_number
fi
