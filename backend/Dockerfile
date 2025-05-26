# Use official Go Alpine image
FROM golang:1.21-alpine

# Install necessary dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /src/app

# Install Air (hot reload dev tool)
RUN go install github.com/cosmtrek/air@latest

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies early
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Expose port (optional, depending on your app)
EXPOSE 3000

# Default command using air for hot reload
CMD ["air"]
