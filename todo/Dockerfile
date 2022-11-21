# Start from golang base image
FROM golang:1.19-alpine3.15

# Add Maintainer info
LABEL maintainer="umarov.doniyor.2001@gmail.com"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /home/app
WORKDIR /home/app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build cmd/main.go

# Expose port 8081 to the outside world
EXPOSE 8081

# Run the executable
CMD [ "./main" ]

# FROM golang:1.19-alpine3.15
# LABEL owner="umarov.doniyor.2001@gmail.com"

# RUN mkdir -p /home/app

# WORKDIR /home/app

# COPY go.sum .
# COPY go.mod .

# RUN go mod download

# COPY . .

# RUN go build cmd/main.go

# CMD ["./main"]
