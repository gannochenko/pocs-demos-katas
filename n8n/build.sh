#!/bin/bash

# Build and start the n8n container with Golang and primitive support
echo "Building custom n8n image with Golang and primitive support..."

# Build the image
docker-compose build --no-cache

echo "Starting the services..."
docker-compose up -d

echo "Waiting for containers to start..."
sleep 10

echo "Testing if primitive CLI is available..."
docker exec n8n primitive -h

echo "Testing Go installation..."
docker exec n8n go version

echo "Testing exiftran installation..."
docker exec n8n exiftran -h

echo "Setup complete! You can access n8n at http://localhost:5678"
echo "Available tools in n8n workflows (use Execute Command node):"
echo "  - primitive -h (geometric primitive art generation)"
echo "  - exiftran -h (EXIF-based image transformations)"
echo "  - go version (Go programming language)"
