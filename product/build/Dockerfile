FROM golang:latest AS builder
WORKDIR /product
COPY product/go.mod product/go.sum ./
RUN go mod download
COPY product/. .
RUN go build -o main ./cmd

EXPOSE 8080  
CMD [ "./main" ]  