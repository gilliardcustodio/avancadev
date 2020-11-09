FROM golang:1.15

WORKDIR go/src/serviceb

COPY . .


RUN go mod download

RUN GOOS=linux go build b.go

ADD start.sh ./start.sh

RUN chmod +x ./start.sh

ENTRYPOINT ["sh", "./start.sh"]


