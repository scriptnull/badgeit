run:
	go run main.go

build:
	go build

badge-github: build
	./badgeit samples/github

badge-gitter: build
	./badgeit samples/gitter

badge-npm: build
	./badgeit samples/npm

test-formatters:
	go test -cover github.com/scriptnull/badgeit/formatters

test-contracts:
	go test -cover github.com/scriptnull/badgeit/contracts

test: test-formatters test-contracts

sample-github-clean:
	rm -rf samples/github

sample-github-init:
	mkdir -p samples/github && cd samples/github && git init && git remote add origin git@github.com:atom/atom.git

sample-gitter-clean:
	rm -rf samples/gitter

sample-gitter-init:
	mkdir -p samples/gitter && cd samples/gitter && git init && git remote add origin git@github.com:scriptnull/badgeit.git

init-samples: sample-github-init sample-gitter-init

clean-samples: sample-github-clean sample-gitter-clean

init: init-samples

clean: clean-samples