export CONFIG_PATH="./config/local.yaml" && go run ./cmd/api </br>

go build -o ./bin/api/broker-service ./cmd/api </br>

docker build -f build/docker/Dockerfile -t broker-service . </br>


docker run \
  -v $(pwd)/config:/usr/local/bin/config \
  -e CONFIG_PATH=/usr/local/bin/config/local.yaml \
  -p4000:4000 broker-service
