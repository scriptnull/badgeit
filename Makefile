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

badge-travis: build
	./badgeit samples/travis

badge-circle: build
	./badgeit samples/circle

badge-semaphore: build
	./badgeit samples/semaphore

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

sample-travis-clean:
	rm -rf samples/travis

sample-travis-init:
	mkdir -p samples/travis && cd samples/travis && git init && git remote add origin git@github.com:rust-lang/cargo.git && touch .travis.yml

sample-circle-clean:
	rm -rf samples/circle

sample-circle-init:
	mkdir -p samples/circle && cd samples/circle && git init && git remote add origin git@github.com:circleci/frontend.git && touch circle.yml

sample-semaphore-clean:
	rm -rf samples/semaphore

sample-semaphore-init:
	mkdir -p samples/semaphore/valid  samples/semaphore/invalid && cd samples/semaphore/valid && git init && git remote add origin git@github.com:argonlaser/badgeit-front.git && \
	cd ../invalid && git init && git remote add origin git@github.com:scriptnull/compilex.git;

init-samples: sample-github-init sample-gitter-init sample-travis-init sample-circle-init sample-semaphore-init

clean-samples: sample-github-clean sample-gitter-clean sample-travis-clean sample-circle-clean sample-semaphore-clean

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
	export CLONE_DIR="`pwd`/worker/storage" && \
	go run ./worker/main.go

queue-run:
	docker run -d --hostname my-rabbit --name badgeit-rabbit -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password rabbitmq:3-management