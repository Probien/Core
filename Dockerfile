FROM golang:1.18.2 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/Probien-Backend

COPY go.mod .
RUN go mod download
COPY . .

RUN go install

FROM scratch
COPY --from=builder /go/bin/Probien-Backend .
EXPOSE 9000
ENTRYPOINT ["./server"]