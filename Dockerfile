FROM golang:1.21.7 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./target/go-oauth2

EXPOSE 8080

CMD ["/app/target/go-oauth2"]