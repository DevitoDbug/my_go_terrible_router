FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum  /app/

RUN go mod download

COPY cmd/ pkg/ /app/

RUN go build -o /app/bin /app/cmd/main/main.go

EXPOSE 8080

CMD ["/app/bin"]