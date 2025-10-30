# Install golang migrate
go install -tags "postgres,mysql,mongodb" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration table
migrate create -ext sql -dir db/migrations create_table_category

# Migration up
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up