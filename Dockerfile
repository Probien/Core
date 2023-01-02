FROM golang:1.18.2 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN ls

COPY . .
RUN go build -o server .

FROM scratch
COPY --from=builder /app/server .
EXPOSE 9000

CMD ["./server"]