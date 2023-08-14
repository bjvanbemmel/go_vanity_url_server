FROM golang:1.20.7-alpine3.17

WORKDIR /proxy

RUN go install github.com/cosmtrek/air@latest

COPY go.mod .air.toml ./

RUN go mod download

CMD [ "air", "-c", ".air.toml" ]
