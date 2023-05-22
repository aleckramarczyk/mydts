FROM golang:1.19.9-alpine3.18

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /root/

COPY ./go.mod go.sum ./
COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

EXPOSE 8080

CMD ["./main"]