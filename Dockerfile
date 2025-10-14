FROM golang:1.25-alpine AS build
WORKDIR /app

# cache
COPY go.mod go.sum ./
RUN go mod download

# copy source and build static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o main

# where it runs
FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]