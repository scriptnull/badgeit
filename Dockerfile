FROM golang:1.8.2

WORKDIR /go/src/github.com/scriptnull/badgeit

COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

RUN go build

RUN cd worker && go-wrapper download
RUN cd worker && go-wrapper install

CMD cd worker && go run main.go