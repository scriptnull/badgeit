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

badge-bower: build
	./badgeit samples/bower

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

api-run:
	export RABBIT_USERNAME="user" && \
	export RABBIT_PASSWORD="password" && \
	export RABBIT_HOSTNAME="localhost" && \
	export RABBIT_PORT="5672" && \
	go run ./api/main.go

worker-run: build
	export RABBIT_USERNAME="user" && \
	export RABBIT_PASSWORD="password" && \
	export RABBIT_HOSTNAME="localhost" && \
	export RABBIT_PORT="5672" && \
	go run ./worker/main.go

docker-queue-init:
	docker run -d --hostname my-rabbit --name badgeit-rabbit -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password rabbitmq:3-management