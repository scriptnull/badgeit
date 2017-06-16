run:
	go run main.go

test-formatters:
	go test github.com/scriptnull/badgeit/formatters

test-contracts:
	go test github.com/scriptnull/badgeit/contracts

test: test-formatters test-contracts