FROM golang:latest
RUN mkdir app
ADD . ./app
WORKDIR ./app
RUN go mod tidy

ENTRYPOINT go run cmd/bank-microservice/main.go

EXPOSE 3000
