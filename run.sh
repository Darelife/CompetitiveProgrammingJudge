#!/usr/bin/env bash
set -e

ROOT="$(cd "$(dirname "$0")" && pwd)"
cd "$ROOT"

echo "=== Building server and client ==="
go build -o /tmp/judge-server ./cmd/server
go build -o /tmp/judge-client ./cmd/client

echo "=== Starting judge server ==="
/tmp/judge-server &
SERVER_PID=$!

# Give the server a moment to start
sleep 1

echo "=== Running client submissions ==="
/tmp/judge-client

echo ""
echo "=== Done. Stopping server. ==="
kill $SERVER_PID
