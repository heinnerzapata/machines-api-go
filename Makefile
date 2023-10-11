include local.env

migrateup:
	migrate -path db/migration/ -database "$(DB_STRING_CONN)" -verbose up

migratedown:
	migrate -path db/migration/ -database "$(DB_STRING_CONN)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migratedown test sqlc