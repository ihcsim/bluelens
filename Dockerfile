FROM golang:1.8
MAINTAINER Ivan Sim, ihcsim@gmail.com

ARG DEBIAN_FRONTEND=noninteractive

COPY . /go/src/github.com/ihcsim/bluelens
WORKDIR /go/src/github.com/ihcsim/bluelens

RUN apt update && \
    apt install -y make && \
		wget -O /usr/bin/glide https://s3-us-west-2.amazonaws.com/go.ihcsim.repo/glide-v0.12.3 && \
    chmod +x /usr/bin/glide && \
		make goagen
