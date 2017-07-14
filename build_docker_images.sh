# Build binary and worker
docker build -t "scriptnull/badgeit-worker:$BRANCH.$BUILD_NUMBER" .
docker images
docker push scriptnull/badgeit-worker:$BRANCH.$BUILD_NUMBER

# Build API
docker build -t "scriptnull/badgeit-api:$BRANCH.$BUILD_NUMBER" ./api/.
docker images
docker push scriptnull/badgeit-api:$BRANCH.$BUILD_NUMBER