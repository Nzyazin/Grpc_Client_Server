# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /

COPY ./template ./
ADD template ./template
RUN go version

COPY ./ ./
RUN go mod download
RUN go build -o /grpc_client ./template/main.go

EXPOSE 8080

CMD ["./grpc_client"]