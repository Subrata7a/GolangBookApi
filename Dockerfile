# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Subrata Sutradhar <ssubrata.cm@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app

# CGO_ENABLED=0
# Disables CGO (C Go) to ensure a statically linked binary.
# CGO is the bridge between Go and C, and disabling it results in a binary that doesn't depend on external C libraries.

# GOOS=linux -> GOOS reders toDockerfile builder Operating System is LInux

# go build -a :
# Forces rebuilding of all packages, even if they are up-to-date.
# This ensures that the binary is rebuilt with the specified configurations.

# -installsuffix cgo
# Appends "cgo" to the output directory during the build.
# This is commonly used to create separate output directories for binaries built with and without CGO.

# -o book_api : Specifies the output name of the binary as "book_api".
# " ." : Indicates the source directory where the Go code is located. In this case, it's the current directory.

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o GolangBookApi .


######## Start a new stage #######
FROM alpine:latest

# --no-cache
# Avoid caching the package index locally.
# This reduces the image size by removing the package index files after installation.

# ca-certificates
# package provides a bundle of trusted CA certificates.
# Installing this package ensures that the required CA certificates are available in the system,
# allowing your application to validate SSL/TLS certificates during communication.
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/GolangBookApi .

EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./GolangBookApi", "start"]