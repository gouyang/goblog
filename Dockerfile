from fedora

RUN yum update -y
RUN yum install git golang -y
RUN mkdir -p /go
ENV GOPATH /go
RUN go get github.com/ouyanggh/goblog
EXPOSE 8080
WORKDIR /go/src/github.com/ouyanggh/goblog
CMD /go/bin/goblog
