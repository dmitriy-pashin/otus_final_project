## builder
FROM golang:1.14 as builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN make build

EXPOSE 9000

CMD ["./bin/otus_final_project"]