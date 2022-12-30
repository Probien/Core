FROM golang:1.18.2

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN ls

COPY . .
RUN go build -o app server.go

EXPOSE 9000

CMD ["./app"]