FROM golang:1.21-alpine

RUN apk update && apk add --no-cache \
    build-base \
    git \
    gcc \
    musl-dev \
    opencv-dev \
    pkgconfig

# Set the Go environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the application code to the working directory
COPY . .

# Install GoCV
RUN go get -u -d gocv.io/x/gocv

RUN go clean -modcache && go mod tidy && go mod vendor

# Build the Go application
RUN go build -o main .

# Expose the port your application runs on
EXPOSE 8000

# Command to run the executable
ENTRYPOINT ["/app/main", "start"]