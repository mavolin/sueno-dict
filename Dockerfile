FROM golang:1.19-alpine

WORKDIR /dict

COPY . .

RUN go build --tags prod -o dict

CMD ["./dict"]