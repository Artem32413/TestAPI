FROM golang:latest AS builder
WORKDIR /analytics
COPY analytics/go.mod analytics/go.sum ./
RUN go mod download
COPY analytics/. .
RUN go build -o main ./cmd  

EXPOSE 8080  
CMD [ "./main" ]  