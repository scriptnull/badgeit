FROM golang:1.8.2

WORKDIR /go/src/github.com/scriptnull/badgeit/worker

COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

RUN go-wrapper build /go/src/github.com/scriptnull/badgeit

CMD ["go-wrapper", "run"] # ["app"]