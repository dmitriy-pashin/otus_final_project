help:
	@echo "Commands:"
	@echo "make build - build project"
	@echo "make run-web - run web app"
	@echo "make tests - run all tests"
	@echo "make tests-unit - run unit tests"
	@echo "make tests-functional - run functional and integration tests"
	@echo "make migrate-up - migrations up"
	@echo "make migrate-down - migrations down"

build:
	rm -r -f bin/*
	go build -i -o bin/otus_final_project ./web/app

run-web:
	./bin/otus_final_project

tests: tests-unit tests-functional

tests-functional:
	go test -tags functional ./... -v -count=1 -parallel=1

tests-unit:
	go test -race -count 100 ./...

migrate-up:
	goose -dir migrations postgres "host=localhost port=5432 user=otus password=otus dbname=antibf sslmode=disable" up

migrate-down:
	goose -dir migrations postgres "host=localhost port=5432 user=otus password=otus dbname=antibf sslmode=disable" down
