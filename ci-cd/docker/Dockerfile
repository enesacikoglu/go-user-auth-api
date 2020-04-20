FROM golang:1.14 AS builder

ENV GOPATH /go
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -v main.go

FROM alpine:latest AS alpine

ENV LANG C.UTF-8
ENV GOPATH /go

RUN apk --no-cache add tzdata ca-certificates
COPY --from=builder /app/main   /app/main
COPY --from=builder /app/infrastructure/resource /app/infrastructure/resource/

WORKDIR /app

RUN chmod +x main

EXPOSE 6161

ENTRYPOINT ["./main"]