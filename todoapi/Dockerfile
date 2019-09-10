FROM golang:1.13

WORKDIR /

RUN go get github.com/gihyodocker/todoapi

COPY . /go/src/github.com/gihyodocker/todoapi

RUN ls -al /go/src/github.com/gihyodocker/todoapi/
RUN cd /go/src/github.com/gihyodocker/todoapi && go build -o bin/todoapi cmd/main.go
RUN cd /go/src/github.com/gihyodocker/todoapi && cp bin/todoapi /usr/local/bin/

CMD ["todoapi"]
