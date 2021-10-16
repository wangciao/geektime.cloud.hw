FROM golang:1.11.2-alpine

COPY src/main.go main.go
RUN go build main.go
EXPOSE 80
ENTRYPOINT ["./main"]
