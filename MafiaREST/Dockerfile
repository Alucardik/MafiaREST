FROM golang:1.18

WORKDIR /usr/local/mafia_rest_server

COPY config config
COPY utils utils
COPY db db
COPY schemes schemes
COPY endpoints endpoints
COPY msgbroker msgbroker
COPY pdfgen pdfgen
COPY server server
COPY worker worker
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod tidy
RUN go build .

CMD ["bash", "-c", "while ! curl -s $QUEUE_HOST:15672 > /dev/null; do echo waiting for rabbitmq; sleep 3; done; ./MafiaREST --mode=$MafiaREST_MODE"]