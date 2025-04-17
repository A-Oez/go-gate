FROM golang:1.23-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o /app/cmd/main ./cmd

EXPOSE 5000

CMD ["/app/cmd/main"]

#build image: docker build -t go-gate .
#start container: docker run -p 5000:3030 go-gate 