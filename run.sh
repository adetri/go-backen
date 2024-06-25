#!/bin/bash

# Function to start the Go server
start_server() {
  echo "Starting Go server..."
  go run main.go &
  SERVER_PID=$!
}

# Function to stop the Go server
stop_server() {
  if [ -n "$SERVER_PID" ]; then
    echo "Stopping Go server..."
    kill $SERVER_PID
  fi
}

# Trap signals and stop the server
trap stop_server EXIT

# Run swag init to generate Swagger documentation
echo "Generating Swagger documentation..."
swag init

# Check if swag init was successful
if [ $? -eq 0 ]; then
  echo "Swagger documentation generated successfully."
  start_server
  # Wait for the server process to complete
  wait $SERVER_PID
else
  echo "Failed to generate Swagger documentation."
  exit 1
fi
