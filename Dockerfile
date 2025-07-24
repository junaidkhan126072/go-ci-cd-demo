# ------------ Stage 1: Build the Go binary ------------
FROM golang:1.21 AS builder

WORKDIR /app

# Copy Go modules and download dependencies first (cache optimization)
COPY go.mod ./
RUN go mod download

# Copy the rest of the app source code
COPY . .

# Build the Go app
RUN go build -o app

# ------------ Stage 2: Create a small runtime image ------------
FROM debian:bookworm-slim

WORKDIR /app

# Copy the compiled Go binary from builder stage
COPY --from=builder /app/app .

# Tell Docker to expose port 8080
EXPOSE 8080

# Set the default command to run the app
CMD ["./app"]
