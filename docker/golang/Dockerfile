FROM golang:latest

RUN apt-get update -qq && \
    apt-get install -y mysql-client vim
    
WORKDIR /go/src/app
COPY . /go/src/app

ENV GO111MODULE=on

RUN go mod download

EXPOSE 8080

# CMD ["go", "run", "/go/src/app/server/cmd/main.go"]