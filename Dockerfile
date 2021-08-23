FROM golang:alpine as builder

WORKDIR /go/src/github.com/MichaelS11/go-tetris
COPY . .

RUN apk add --no-cache git
RUN go get ./
RUN go build -ldflags="-extldflags=-static" -o /go/bin/go-tetris

FROM scratch
WORKDIR /
COPY --from=builder /go/bin/go-tetris .

ENTRYPOINT ["./go-tetris"]
