FROM golang:1.22

WORKDIR /thinker

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux go build -o thinker ./cmd/thinker/main.go

EXPOSE 8080

CMD ["./thinker"]