up:
	docker-compose --file docker/docker-composer.yaml up -d
down:
	docker-compose --file docker/docker-composer.yaml down
status:
	docker-compose --file docker/docker-composer.yaml ps
api:
	docker exec -it bookstore_api bash
run:
	go run app/main.go
