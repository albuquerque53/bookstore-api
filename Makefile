.SILENT: migrateup
.SILENT: migratedown
up:
	docker-compose --file build/docker-composer.yml up -d
down:
	docker-compose --file build/docker-composer.yml down
migrateup:
	migrate -path internal/infra/database/migration -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:3306)/${MYSQL_DATABASE}" -verbose up
migratedown:
	migrate -path internal/infra/database/migration -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:3306)/${MYSQL_DATABASE}" -verbose down
status:
	docker-compose --file build/docker-composer.yml ps
api:
	docker exec -it bookstore_api bash
test:
	go test ./... 
serve:
	go run cmd/main.go
