FROM golang:1.19.0 as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/myapp .

COPY --from=builder /usr/src/app/views ./views

COPY --from=builder /usr/src/app/public ./public

EXPOSE 8080

CMD ["./myapp"]
