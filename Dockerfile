FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

COPY . .

RUN go build -o pricefetcher

FROM ubuntu:latest

LABEL authors="tuvshno"

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/pricefetcher /pricefetcher

EXPOSE 3000

CMD ["/pricefetcher", "-listenaddr", ":3000"]
