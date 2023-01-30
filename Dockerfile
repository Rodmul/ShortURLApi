FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./short_link ./cmd/main.go

EXPOSE 4000

ENTRYPOINT ["./short_link"]
