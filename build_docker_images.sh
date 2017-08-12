API_CHANGED_FILE_COUNT=`git diff --name-only HEAD~1..HEAD api/ | wc -l`;
WORKER_CHANGED_FILE_COUNT=`git diff --name-only HEAD~1..HEAD worker/ contracts/ common/ formatters/ | wc -l`;

if [ $API_CHANGED_FILE_COUNT -gt 0 ]; then
    echo "API has changes"
    # Build API
    docker build -t "scriptnull/badgeit-api:$BRANCH.$BUILD_NUMBER" ./api/.
    docker images
    docker push scriptnull/badgeit-api:$BRANCH.$BUILD_NUMBER
fi

if [ $WORKER_CHANGED_FILE_COUNT -gt 0 ]; then
    echo "Worker has changes"
    # Build binary and worker
    docker build -t "scriptnull/badgeit-worker:$BRANCH.$BUILD_NUMBER" .
    docker images
    docker push scriptnull/badgeit-worker:$BRANCH.$BUILD_NUMBER
fi

