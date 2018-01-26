FROM alpine

COPY gopath/bin/kuberdemo2 /go/bin/kuberdemo2

ENTRYPOINT /go/bin/kuberdemo2
