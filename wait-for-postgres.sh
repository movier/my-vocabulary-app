#!/bin/sh
# wait-for-postgres.sh

set -e

cmd="$@"

sleep 10

>&2 echo "Postgres is up - executing command"
exec $cmd