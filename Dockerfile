FROM golang:1.21.7 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /www

COPY . ./

RUN go mod download

RUN go build -o app .

FROM alpine

EXPOSE 8088

COPY --from=builder /www/app /
COPY --from=builder /www/configs /configs

CMD ["/app"]