.SILENT:

run:
	go run cmd/main.go

migrate_create:
	migrate create -ext sql -dir db/migrations -seq $(NAME)

migrate_up:
	migrate -source file:/Users/amir/Documents/work/Self/Go_projects/mini_chat/db/migrations -database postgres://localhost:5432/chat?sslmode=disable up $(LEVEL)

migrate_down:
	yes y | migrate -source file:/Users/amir/Documents/work/Self/Go_projects/mini_chat/db/migrations -database postgres://localhost:5432/chat?sslmode=disable down $(LEVEL)