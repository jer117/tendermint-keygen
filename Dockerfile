# Stage 1: Build the Go binary
FROM golang:latest AS build

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY key_generation.go .

RUN go get -d -v ./...
RUN go build -o key_generation key_generation.go

# Stage 2: Create the final Docker image
FROM alpine:latest

# Install necessary libraries if required by the binary (comment out if not needed)
RUN apk --no-cache add ca-certificates

# Copy the binary from the previous stage
COPY --from=build /go/src/app/key_generation /usr/local/bin/priv_key_gen

# Set the working directory for the binary (optional)
WORKDIR /data

# Run the binary when the container starts
CMD ["priv_key_gen"]
