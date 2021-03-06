FROM golang:latest
RUN mkdir "/books-api"
ADD . /books-api/
WORKDIR /books-api/cmd
RUN go build -o main .
ENTRYPOINT go run main.go