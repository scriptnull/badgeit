#!/bin/bash
set -e

if [ "$IS_PULL_REQUEST" == true ]
then
    echo "Skipping building of image for PRs"
    exit 0
fi

API_CHANGED_FILE_COUNT=`git diff --name-only HEAD~1..HEAD api/ | wc -l`;
WORKER_CHANGED_FILE_COUNT=`git diff --name-only HEAD~1..HEAD worker/ contracts/ common/ formatters/ | wc -l`;

build_worker() {
    # Build binary and worker
    docker build -t "scriptnull/badgeit-worker:$BRANCH.$BUILD_NUMBER" .
    docker images
    docker push scriptnull/badgeit-worker:$BRANCH.$BUILD_NUMBER
    ./telegram.sh "New Worker image available: scriptnull/badgeit-worker:$BRANCH.$BUILD_NUMBER"
}

if [ $API_CHANGED_FILE_COUNT -gt 0 ]; then
    echo "API has changes"
    # Build API
    docker build -t "scriptnull/badgeit-api:$BRANCH.$BUILD_NUMBER" ./api/.
    docker images
    docker push scriptnull/badgeit-api:$BRANCH.$BUILD_NUMBER
    ./telegram.sh "New API image available: scriptnull/badgeit-api:$BRANCH.$BUILD_NUMBER"
fi

if [ $WORKER_CHANGED_FILE_COUNT -gt 0 ]; then
    echo "Worker has changes"
    build_worker
fi

if [ $FORCE_WORKER_BUILD == true ]; then
    echo "FORCE_WORKER_BUILD: $FORCE_WORKER_BUILD"
    build_worker
fi

echo "Build Image script End"

