run:
	go run main.go

build:
	go build

badge-npm: build
	./badgeit samples/npm

test-formatters:
	go test github.com/scriptnull/badgeit/formatters

test-contracts:
	go test github.com/scriptnull/badgeit/contracts

test: test-formatters test-contracts