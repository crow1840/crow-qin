FROM golang:alpine

MAINTAINER liuwei
WORKDIR /
ADD crow-qin /crow-qin
ENTRYPOINT ["/crow-qin"]
#CMD ["/crow-qin"]