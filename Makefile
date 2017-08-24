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
	./badgeit samples/semaphore/valid

badge-codecov: build
	./badgeit samples/codecov/valid

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

sample-codecov-clean:
	rm -rf samples/codecov

sample-codecov-init:
	mkdir -p samples/codecov/valid  samples/codecov/invalid && cd samples/codecov/valid && git init && git remote add origin git@github.com:sindresorhus/make-dir.git && \
	cd ../invalid && git init && git remote add origin git@github.com:scriptnull/compilex.git;

init-samples: sample-github-init sample-gitter-init sample-travis-init sample-circle-init sample-semaphore-init sample-codecov-init

clean-samples: sample-github-clean sample-gitter-clean sample-travis-clean sample-circle-clean sample-semaphore-clean sample-codecov-clean

init: init-samples

clean: clean-samples

api-run:
	export REDIS_HOSTNAME="localhost" && \
	export REDIS_PORT="6379" && \
	go run ./api/main.go

worker-run: build
	export CLONE_DIR="`pwd`/worker/storage" && \
	export REDIS_HOSTNAME="localhost" && \
	export REDIS_PORT="6379" && \
	go run ./worker/main.go

redis-run:
	docker run --name badgeit-redis -p 6379:6379 -d redis