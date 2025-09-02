FROM golang:tip-alpine3.21
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o demo
EXPOSE 8080
CMD ["./demo"]

