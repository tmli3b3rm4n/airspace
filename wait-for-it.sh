#!/usr/bin/env bash
# wait-for-it.sh: Wait for a service to become available.

set -e

HOST="postgres"
PORT="5432"
shift 2
CMD="$@"

if [[ -z "$HOST" || -z "$PORT" ]]; then
  echo "Usage: $0 host port [cmd...]"
  exit 1
fi

echo "Waiting for $HOST:$PORT to be available..."

while ! nc -z "$HOST" "$PORT"; do
  sleep 1
done

echo "$HOST:$PORT is available. Starting the command..."
exec $CMD