# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ADD . /go/src/github.com/points-org/dropbox

RUN go get -v github.com/points-org/dropbox

RUN go install github.com/points-org/dropbox

ENTRYPOINT /go/bin/dropbox

EXPOSE 8081