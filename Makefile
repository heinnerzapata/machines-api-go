migrateup:
	migrate -path db/migration/ -database "postgres://amxolmvh:Ifc6YHUH5PdD1fTrVb0046CpbpuAAJXz@trumpet.db.elephantsql.com/amxolmvh" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgres://amxolmvh:Ifc6YHUH5PdD1fTrVb0046CpbpuAAJXz@trumpet.db.elephantsql.com/amxolmvh" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migratedown test sqlc