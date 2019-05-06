FROM golang:1.11 AS building-stage

WORKDIR /go/src/github.com/orsa-scholis/orsum-inflandi-II/backend

ENV CGO_ENABLED=0
ENV GO111MODULE=on

COPY *.go ./
COPY go.* ./

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/orsum-inflandi-ii .

FROM alpine AS serving-stage

WORKDIR /root/
COPY --from=building-stage /go/bin/orsum-inflandi-ii .

EXPOSE '4560'
ENTRYPOINT ["./orsum-inflandi-ii"]
