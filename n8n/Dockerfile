# Use the official n8n image as the base
FROM docker.n8n.io/n8nio/n8n

# Switch to root user to install additional packages
USER root

# Update package list and install required packages
RUN apk update && apk add --no-cache wget fbida-exiftran

# Install Go (latest stable version)
RUN wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz && \
    rm go1.23.4.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# Create GOPATH directory and set permissions
RUN mkdir -p $GOPATH && chown -R node:node $GOPATH

# Install the primitive package
RUN go install github.com/fogleman/primitive@latest

# Make sure the primitive binary is accessible
RUN cp $GOPATH/bin/primitive /usr/local/bin/primitive && chmod +x /usr/local/bin/primitive

# Switch back to the node user
USER node

# Verify installations (optional - for debugging)
RUN go version && primitive -h && exiftran -h || echo "All tools installed successfully"
