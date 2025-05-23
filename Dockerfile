FROM golang:alpine as builder

RUN apk add --no-cache git 

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]

