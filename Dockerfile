FROM golang:1.9

ENV PATH=$PATH:$GOPATH/bin
WORKDIR $GOPATH/src/github.com/shavit/stockwell

RUN apt-get update \
  && apt-get -y install \
    nano

ADD $PWD/ $GOPATH/src/github.com/shavit/stockwell
RUN go get ./...

# Arguments will be appended
# ENTRYPOINT ["go", "run", "main.go"]

# Arguments will override
# CMD ["go", "run", "main.go"]
CMD ["bash"]
