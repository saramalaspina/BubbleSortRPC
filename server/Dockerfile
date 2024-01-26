FROM golang:1.21

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o server

CMD ["./server"]


