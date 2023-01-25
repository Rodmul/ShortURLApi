create_migration:
# make create_migration name=name_your_migration_without_spaces
	migrate create -ext sql -dir db/migrations -seq ${name}
migrate:
# make migrate password=postgres_password host=your_host port=.... mode=up/down
	migrate -database 'postgres://postgres:${password}@${host}:${port}/short_link?sslmode=disable' -path ./db/migrations ${mode}
build_image:
	docker build -t rodmul/short_link:v1 .
run_with_db:
	docker run -d -p 4000:4000 --name short_link_container rodmul/short_link:v1 "--use_db_storage"
run_with_im:
	docker run -d -p 4000:4000 --name short_link_container rodmul/short_link:v1
run_im:
	go build -o . cmd/main.go
	./main --use_db_storage
fmt:
	go fmt ./...