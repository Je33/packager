.PHONY: build build-prof build-run gen test test-html bench run lint dc

build:
	go build -o ./build/api ./cmd/api/main.go

build-prof: build
	go tool pprof —text ./build/api

build-run: build
	./build/api

gen:
	go generate ./...

test:
	go test -v -coverprofile cover.out ./...

test-html: test
	go tool cover -html=cover.out

bench:
	go test -bench=. -benchmem ./...

run:
	go run -race ./cmd/api/main.go

lint:
	golangci-lint run

dc:
	docker-compose up