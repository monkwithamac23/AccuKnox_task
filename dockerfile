# Use a minimal base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Expose the port on which your application runs
EXPOSE 8080

# Define the command to run the application
CMD ["./app"]
