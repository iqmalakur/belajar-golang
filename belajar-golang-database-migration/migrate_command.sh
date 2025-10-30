# Install golang migrate
go install -tags "postgres,mysql,mongodb" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration table
migrate create -ext sql -dir db/migrations create_table_category
