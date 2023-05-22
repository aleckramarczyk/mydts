#Latest golang base image as builder
FROM golang:latest as builder

WORKDIR /app

COPY ./go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

#Create app container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

#Copy the binary artifact from builder container to app container
COPY --from=builder /app/main .

#Copy configuration
COPY conf conf

EXPOSE 8080

CMD ["./main"]