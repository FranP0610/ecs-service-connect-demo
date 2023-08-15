# syntax=docker/dockerfile:1
FROM public.ecr.aws/docker/library/golang:1.20-alpine3.18
#FROM golang:1.20-alpine3.18

# Install curl
RUN apk add curl
# Set default destination for all the subsequent commands
WORKDIR /app
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
# Copy the source code
COPY /cmd/api-a/ ./

# Build a binary file named webapp and
# located in the root of the filesystem of the image
RUN CGO_ENABLED=0 GOOS=linux go build -o /webapp

EXPOSE 8080

# Run
CMD ["/webapp"]