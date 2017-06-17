run:
	go run main.go

build:
	go build

badge-github: build
	./badgeit samples/github

badge-npm: build
	./badgeit samples/npm

test-formatters:
	go test -cover github.com/scriptnull/badgeit/formatters

test-contracts:
	go test -cover github.com/scriptnull/badgeit/contracts

test: test-formatters test-contracts