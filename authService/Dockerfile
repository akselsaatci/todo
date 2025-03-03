# Stage 1: Builder
FROM golang:alpine AS builder
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install Air for hot-reloading
RUN go install github.com/cosmtrek/air@latest

# Copy the source code into the container
COPY . .

# Build the Go binary (optional, as Air will run the code directly)
RUN go build -o main ./cmd/main.go

# Stage 2: Final Image
FROM alpine:latest

# Arguments for environment variables
ARG DATABASE_HOST
ARG DATABASE_PORT
ARG DATABASE_USER
ARG DATABASE_PASSWORD
ARG DATABASE_NAME
ARG JWT_SECRET_KEY
ARG SERVER_PORT

# Environment Variables
ENV DB_HOST=${DATABASE_HOST}
ENV DB_PORT=${DATABASE_PORT}
ENV DB_USER=${DATABASE_USER}
ENV DB_PASSWORD=${DATABASE_PASSWORD}
ENV DB_NAME=${DATABASE_NAME}
ENV JWT_SECRET=${JWT_SECRET_KEY}
ENV PORT=${SERVER_PORT}

WORKDIR /root/

# Copy Air from the builder stage
COPY --from=builder /app/main .

# Install dependencies in the final image (including Air)
RUN apk --no-cache add bash

# Copy Air binary from the builder stage
COPY --from=builder /go/bin/air /usr/local/bin/air

# Expose the server port
EXPOSE 8080

# Command to run Air for hot-reloading
CMD ["air", "-c", ".air.toml"]

