FROM golang:latest AS builder
WORKDIR /inventory
COPY inventory/go.mod inventory/go.sum ./
RUN go mod download
COPY inventory/. .
RUN go build -o main ./cmd 

EXPOSE 8081  
CMD [ "./main" ]  