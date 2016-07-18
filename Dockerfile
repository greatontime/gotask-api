FROM golang
ADD . /go/src/github.com/greatontime/gotask-api
WORKDIR /go/src/github.com/greatontime/gotask-api
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/greatontime/gotask-api
ENTRYPOINT /go/bin/gotask-api
EXPOSE 8080
