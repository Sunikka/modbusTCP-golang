# Use the official Ubuntu base image
FROM ubuntu:latest

# Set the working directory inside the container
WORKDIR /app

# Expose the ModbusTCP default port
EXPOSE 502

# Set the command to run the Go binary
CMD ["./mbServer"]
