# Use the Raspberry Pi Debian base image
FROM balenalib/raspberry-pi-debian:latest

# Install Go
RUN apt-get update && apt-get install -y golang

# Set the working directory within the container
WORKDIR /app

# Copy all the contents of your repository to the container
COPY . .

RUN go mod download
# RUN go mod tidy

# Build the Go executable
RUN go build -o app api_driver.go
RUN chmod +x app
EXPOSE 8080

# Specify the command to run when the image starts
CMD ["app"]
