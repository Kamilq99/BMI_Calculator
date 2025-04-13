FROM golang:1.23.8-bookworm

WORKDIR /calculate

COPY main.go ./
COPY go.mod ./
COPY go.sum ./
COPY handlers ./handlers
COPY models ./models
COPY tests ./tests

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]