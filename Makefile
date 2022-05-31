.SILENT: migrateup
.SILENT: migratedown
up:
	docker-compose --file build/docker-composer.yml up -d && docker exec -it bookstore_api bash
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
install:
	go mod tidy && go mod vendor
test:
	go test ./... 
serve:
	go run cmd/main.go
