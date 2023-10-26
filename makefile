
include .env
export

start:
	go run cmd/main.go

tests/local: 
	go test -tags=tests ./internal/adapters/http

tests/integration: 
	go test -tags=testing ./internal/tests

tests/integration/testify: 
	go test -tags=testify ./internal/tests

tests/integration/containers: 
	go test -tags=containers ./internal/tests

tests/integration/bdd: 
	go test -tags=bdd ./internal/tests	

tests/integration/verbose: 
	go test -v ./internal/tests

docs:
	docker build --tag swaggo/swag:1.8.1 . --file swaggo.Dockerfile && \
	docker run --rm --volume ${PWD}:/app --workdir /app swaggo/swag:1.8.1 /root/swag init \
		--parseDependency \
		--parseInternal \
		--dir ./internal/adapters/http \
		--generalInfo swagger.go \
		--output ./api/swagger/public \
		--parseDepth 1

oapi-generate:
	oapi-codegen  --config ./oapi-codegen.yaml ./api/openapi3/openapi.yml