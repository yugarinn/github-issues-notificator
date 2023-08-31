FROM golang:1.20

ENV GOPATH=/go

RUN mkdir /notificator
WORKDIR /notificator

ADD go.mod ./go.mod
ADD go.sum ./go.sum

ADD . .

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.10.0/wait /usr/local/bin/wait
RUN chmod +x /usr/local/bin/wait

# Install dev dependencies
RUN go mod download && go mod verify
RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install gotest.tools/gotestsum@latest
RUN go install github.com/codegangsta/gin@latest

ENV NOTIFICATOR_ENV=development
EXPOSE 3000

CMD /usr/local/bin/wait && ./setup.sh
