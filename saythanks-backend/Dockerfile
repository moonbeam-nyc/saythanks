# Use the latest stable Go image
FROM golang:1.23

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o main .

# Command to run the executable with air
CMD ["air"]