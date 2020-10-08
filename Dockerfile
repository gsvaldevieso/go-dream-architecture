FROM golang:1.14-stretch

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go mod download
RUN go get github.com/pilu/fresh