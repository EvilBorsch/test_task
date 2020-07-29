TARGET_PORT:=8080

.PHONY: run_local
run_local:
	go run src/main.go

.PHONY: build
build:
	docker build -t go .

.PHONY: run
run:
	docker run -p $(TARGET_PORT):$(TARGET_PORT) go


.PHONY: test
test:
	go test -covermode=atomic -coverpkg=./... -coverprofile=cover ./... && go tool cover -func=cover && rm cover



