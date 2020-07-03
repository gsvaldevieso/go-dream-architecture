init:
	cp .env-example .env

test:
	go test -cover ./...

fmt:
	go fmt ./...

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

logs:
	docker logs -f app