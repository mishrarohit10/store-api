FROM golang:1.19.0

WORKDIR /user/src/app

# RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy

CMD go run .

