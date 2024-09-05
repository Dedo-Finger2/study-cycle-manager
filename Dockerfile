FROM golang:1.23-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o scm ./cmd/scm/main.go

CMD ["./scm"]
