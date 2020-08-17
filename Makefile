help:
	@echo "Commands:"
	@echo "make init - init project"
	@echo "make build - build project"
	@echo "make run-web - run web app"
	@echo "make tests - run all tests"
	@echo "make tests-unit - run unit tests"
	@echo "make tests-functional - run functional and integration tests"

build: dependencies-install
	rm -r -f bin/*
	go build -i -o bin/otus_final_project ./cmd/web

run-web:
	./bin/otus_final_project_web

tests: tests-unit tests-functional

tests-functional:
	go test -tags functional ./... -v -count=1 -parallel=1

tests-unit:
	go test ./... -v
