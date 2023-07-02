# Use the official Go image from the Docker Hub
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Install make
RUN apt-get update && apt-get install -y make

# Download dependencies
RUN go mod download

# Get swag, generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Add the $GOPATH/bin directory to the PATH
ENV PATH=$PATH:$GOPATH/bin

# Generate swagger docs
RUN swag init -g cmd/main.go -o cmd/docs -d .

# Using make command to build the application
RUN make build

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/bin/server"]