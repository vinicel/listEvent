FROM golang:1.19.3

# Enable Go Modules
ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR go/src/github.com/vinicel/listEvent

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 8900

# command to be used to execute when the image is used to start a container
CMD ["go", "run", "main.go"]