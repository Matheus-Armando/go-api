#!/bin/sh

# Start json-server in the background
json-server --watch /app/data/db.json --host 0.0.0.0 --port 3000 &

# Start the Go API
/app/go-api