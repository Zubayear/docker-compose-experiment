FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o /hello

EXPOSE 9000

CMD [ "/hello" ]