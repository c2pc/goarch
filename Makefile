.PHONY:
.SILENT:
.DEFAULT_GOAL := run

run:
	docker-compose up --remove-orphans app

debug:
	docker-compose up --remove-orphans debug

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage