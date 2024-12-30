FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy go mod and sum files
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
EXPOSE 8000

# Build the Go app
RUN go build -o main cmd/main.go

# Run the Go app
CMD ["./main"]