FROM golang:latest AS builder
WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod download
COPY ../ ./
RUN go build -o main ./cmd  

EXPOSE 8080  
CMD [ "./main" ]  