# Use an official Golang runtime as a parent image
FROM golang:1.19.11-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Download dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o app ./cmd

# Create a non-root user and switch to that user
RUN adduser -D api-gateway
USER api-gateway

# Expose port 3000 for the application
EXPOSE 3000

# Define the command to run the executable
CMD ["./app"]
