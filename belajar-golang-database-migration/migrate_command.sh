# Install golang migrate
go install -tags "postgres,mysql,mongodb" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration table
migrate create -ext sql -dir db/migrations create_table_category

# Migration up
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up

# Migration down
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down

# Create 3 migrations
migrate create -ext sql -dir db/migrations create_table_first
migrate create -ext sql -dir db/migrations create_table_second
migrate create -ext sql -dir db/migrations create_table_third

# Migration up all versions
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up

# Migration down all versions
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down

# Migration up specific versions
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up 2

# Migration down specific versions
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down 1

# Dirty state
migrate create -ext sql -dir db/migrations sample_dirty_state
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down 1 # error dirty state

# Force update version
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force 20251030093931

# Check current version
migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations version
