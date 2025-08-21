FROM golang:latest AS builder
WORKDIR /warehouses
COPY warehouses/go.mod warehouses/go.sum ./
RUN go mod download
COPY warehouses/. .
RUN go build -o main ./cmd 

EXPOSE 8080  
CMD [ "./main" ]  