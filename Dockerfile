# Use an official Golang runtime as a parent image
FROM golang:1.23.4

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the API
RUN go build -o github-api-manager

# Expose the port used by the API
EXPOSE 8080

# Run the API
CMD ["./github-api-manager"]
