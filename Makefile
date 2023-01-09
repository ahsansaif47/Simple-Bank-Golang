postgres:
	docker run --name sbankpostgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=61926114 -d postgres

migrateup:
	migrate -path db/migration -database "postgresql://postgres:61926114@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:61926114@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlcinit:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc init

sqlc: 
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate