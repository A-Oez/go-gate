FROM golang:1.24.2-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o /app/main 

EXPOSE 5000
ENV BACKEND_PORT=5000

CMD ["/app/main"]

#build image: docker build -t go-gate .
#start container: docker run -p 5000:5000 go-gate 