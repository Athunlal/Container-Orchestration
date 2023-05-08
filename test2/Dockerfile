
# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the container
EXPOSE 8080

# Run the Go app when the container starts
CMD ["./main"]
