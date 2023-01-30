FROM golang:1.19.5

WORKDIR /go/src/github.com/vinicel/listEvent

COPY go.* ./

RUN go mod download

COPY . .

EXPOSE 8900