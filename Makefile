up:
	docker-compose --file docker/docker-composer.yml up -d
down:
	docker-compose --file docker/docker-composer.yml down
status:
	docker-compose --file docker/docker-composer.yml ps
api:
	docker exec -it bookstore_api bash
run:
	go run cmd/main.go
