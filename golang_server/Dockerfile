# Use an official Golang image as a base
FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o server cmd/api/main.go

# Expose the application port
EXPOSE 8888

# Command to run the application
CMD ["./server"]
