migrate-init:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate-up:
	migrate -path sql/migrations -database "mysql://root:root@tcp(mysql:3306)/mydb" up

migrate-down:
	migrate -path sql/migrations -database "mysql://root:root@tcp(mysql:3306)/mydb" down

.PHONY: migrate-init migrate-up migrate-down