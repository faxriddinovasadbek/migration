DB_URL := "postgres://asadbek:1234@localhost:5432/handlar?sslmode=disable"

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate-file:
	migrate create -ext sql -dir migrations/ -seq work_history