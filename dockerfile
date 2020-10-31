
# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Robert Sanchez <rasr12@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /

# Copy files
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get -d github.com/gorilla/mux

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]