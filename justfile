dev:
  @air

build: test
  @go build -o ./bin/mes3hacklab.ssh ./...

test:
 @go test ./... 

run:
  @go run ./...

