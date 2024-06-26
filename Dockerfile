# Start by using the official Golang image as base
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /go-authentication

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
