up:
	docker-compose --file build/docker-composer.yml up -d
down:
	docker-compose --file build/docker-composer.yml down
status:
	docker-compose --file build/docker-composer.yml ps
api:
	docker exec -it bookstore_api bash
run:
	go run cmd/main.go
