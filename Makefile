migrate-init:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate-up:
	migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3307)/mydb" up

migrate-down:
	migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3307)/mydb" down

.PHONY: migrate-init migrate-up migrate-down