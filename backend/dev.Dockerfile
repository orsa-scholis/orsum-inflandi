FROM golang:1.12

WORKDIR /go/src/github.com/orsa-scholis/orsum-inflandi-II/backend

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go get github.com/oxequa/realize

EXPOSE 4560
ENTRYPOINT ["realize", "start"]
