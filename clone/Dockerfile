FROM golang:1.22.5

WORKDIR /app

RUN go mod init example/m

COPY . .

RUN go mod tidy

RUN go build -o main

EXPOSE 8001

CMD ["./main"]
