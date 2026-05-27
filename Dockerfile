# Build stage

FROM golang:1.19-alpine AS builder



WORKDIR /app



# Copy go mod and sum files

COPY go.mod go.sum ./



# Download dependencies

RUN go mod download



# Copy source code

COPY . .



# Build the application

RUN go build -o main .



# Run stage

FROM alpine:latest



WORKDIR /app



# Copy the binary from builder

COPY --from=builder /app/main .



# Expose port (as seen in main.go, default is 2020)

EXPOSE 2020



# Command to run

CMD ["./main"]

