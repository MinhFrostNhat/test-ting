FROM golang:1.21

WORKDIR /app

COPY . .

# Build the Go application
RUN go build -o main ./main.go

CMD ["./main"]

EXPOSE 8081
