FROM golang:1.15

WORKDIR go/src/servicea

COPY . .

RUN go mod download

RUN GOOS=linux go build a.go

EXPOSE 9090

ENTRYPOINT ["./a"]