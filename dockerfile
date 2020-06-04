# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Henri Touchard"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Copy go source files 
COPY pkg ./pkg
COPY main.go ./

# Copy react build folder
COPY frontBuild ./frontBuild

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
# COPY . .

RUN ls -la 

# Build the Go app 
#GOOS=linux
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .
COPY --from=builder /app/frontBuild ./frontBuild

RUN ls -la

# Expose port to the outside world
EXPOSE 5000

#Command to run the executable
CMD [ "./main" ]