# Use the official Ubuntu 22.04 as the base image
FROM ubuntu:22.04

# Set environment variables to avoid interactive prompts during package installation
ENV DEBIAN_FRONTEND=noninteractive

# Install necessary dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    git \
    telnetd -y \
    build-essential && \
    rm -rf /var/lib/apt/lists/*

# Download and install Go
ENV GOLANG_VERSION=1.20
RUN curl -LO https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz

# Set up Go environment variables
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Create the working directory
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

RUN touch .env

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -o main.app /app/src/project

# Expose the port the application will run on
EXPOSE 1323

# Command to run the application
CMD ["./main.app"]
