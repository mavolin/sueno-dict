FROM golang:1.19-alpine

WORKDIR /dict

COPY . /dict

RUN go build -o dict github.com/mavolin/sueno-dict/cmd/dict

EXPOSE 8080

CMD ["./dict"]