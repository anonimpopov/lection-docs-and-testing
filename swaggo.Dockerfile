# Start from the latest golang base image
FROM golang:1.17-alpine as builder

# Set Environment Variables
ENV HOME /app
ENV GOBIN /app
ENV CGO_ENABLED 0
ENV GOOS linux

# Set the Current Working Directory inside the container
WORKDIR /app

# Install git and swag
RUN apk add --no-cache git && \
  go install -v -a -installsuffix cgo github.com/swaggo/swag/cmd/swag@v1.8.1

######## Start a new stage from golang base image #######
FROM golang:1.17-alpine

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/swag .
