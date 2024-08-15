FROM golang:1.20.10

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN ls

COPY . .
WORKDIR /app/cmd

RUN go build -o app .

WORKDIR /app

EXPOSE 9000

ENTRYPOINT ["./cmd/app", "-migrate=true"]