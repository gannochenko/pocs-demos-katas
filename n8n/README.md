# n8n with Golang and Primitive Support

This setup extends the official n8n Docker image to include Golang and the `primitive` CLI tool.

## What's Included

- **n8n**: The official n8n workflow automation platform
- **Golang**: Latest stable version for executing Go programs
- **Primitive**: A CLI tool for reproducing images with geometric primitives
- **Exiftran**: A CLI tool for lossless EXIF-based image transformations

## Files

- `Dockerfile`: Extends the official n8n image with Go and primitive
- `docker-compose.yml`: Docker Compose configuration
- `build.sh`: Build and test script

## Quick Start

1. **Build and start the services:**

   ```bash
   ./build.sh
   ```

2. **Or manually:**

   ```bash
   docker-compose build --no-cache
   docker-compose up -d
   ```

3. **Access n8n:**
   Open http://localhost:5678 in your browser

## Testing the Setup

You can test that Go and primitive are available by running:

```bash
# Test Go installation
docker exec n8n go version

# Test primitive CLI
docker exec n8n primitive -h

# Test exiftran CLI
docker exec n8n exiftran -h
```

## Using in n8n Workflows

In your n8n workflows, you can use the **Execute Command** node to run:

- `go version` - Check Go version
- `primitive -h` - See primitive help
- `primitive input.jpg -o output.png -n 100` - Generate primitive art
- `exiftran -h` - See exiftran help
- `exiftran -a -i image.jpg` - Auto-rotate image based on EXIF orientation

## Environment Variables

The setup includes the following environment variables for n8n:

- Database connection to PostgreSQL
- Timezone settings
- Volume mounts for data persistence

## Volumes

- `n8n_data`: Persists n8n configuration and workflows
- `postgres_data`: Persists PostgreSQL database
- Host directories mounted for file access

## Services

- **n8n**: Main application on port 5678
- **postgres**: Database on port 5432

## Troubleshooting

If you encounter issues:

1. Check container logs:

   ```bash
   docker-compose logs n8n
   ```

2. Rebuild without cache:

   ```bash
   docker-compose build --no-cache
   ```

3. Verify Go, primitive, and exiftran installation:
   ```bash
   docker exec n8n which go
   docker exec n8n which primitive
   docker exec n8n which exiftran
   ```
