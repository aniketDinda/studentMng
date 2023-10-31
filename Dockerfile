FROM golang:latest

WORKDIR /app

COPY . . 

RUN go get -d -v ./...

RUN go build -o api .

EXPOSE 8085

CMD ["./api"]