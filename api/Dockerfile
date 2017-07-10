FROM golang:1.8.2

WORKDIR /go/src/github.com/scriptnull/badgeit/api

COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]