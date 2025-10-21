# Use a Go-Alpine image to compile the app. Named this stage build
FROM golang:1.25-alpine AS build
# Run all following commands from /app inside this stage
WORKDIR /app    

# Copy dependency files first to enable cached module downloads
COPY go.mod go.sum ./
# Download Go modules (skips on cache hits when mods don’t change)
RUN go mod download

# Copy the rest of the source code into the build workspace
COPY . .
# Compile a single self-contained binary named "main"
RUN CGO_ENABLED=0 go build -o main

# Start a runtime image
FROM alpine:3.20
# Work from /app in the runtime container
WORKDIR /app
# Copy the compiled binary from the build stage into this image
COPY --from=build /app/main .
# app listens on port 8080
EXPOSE 8080
# Run the compiled program when the container starts (docker run ...)
CMD ["./main"]