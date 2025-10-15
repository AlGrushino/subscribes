FROM golang:1.25.2

RUN mkdir /subscribe

WORKDIR /subscribe

COPY . .

RUN touch ".env"

RUN go mod tidy

RUN go build -o /main ./cmd/main.go

COPY start.sh /start.sh
RUN chmod +x /start.sh

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["/start.sh"]