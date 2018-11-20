FROM golang:1.11.0-alpine

WORKDIR /go/src/github.com/orsa-scholis/orsum-inflandi-II/backend
COPY . .

ENV CGO_ENABLED=0

RUN go install -v -i

EXPOSE '4560'

CMD 'backend'
