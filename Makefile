run:
	go run ./cmd/app/main.go
test:
	mkdir -p .output
	go test ./internal/... -coverprofile=.output/cover.out
	go tool cover -html=.output/cover.out -o=.output/cover.html
## 開発環境のマイグレーション実行
migrate-up:
	sql-migrate up -config=./internal/infrastructure/migration/dbconfig.yml
migrate-up-dryrun:
	sql-migrate up -dryrun -config=./internal/infrastructure/migration/dbconfig.yml
migrate-down:
	sql-migrate down -config=./internal/infrastructure/migration/dbconfig.yml
migrate-down-dryrun:
	sql-migrate down -dryrun -config=./internal/infrastructure/migration/dbconfig.yml
migrate-status:
	sql-migrate status -config=./internal/infrastructure/migration/dbconfig.yml
## sqlcによるSQL文からのコード生成
sqlc-gen:
	sqlc generate -f ./internal/infrastructure/sqlc/sqlc.yml
sqlc-compile:
	sqlc compile -f ./internal/infrastructure/sqlc/sqlc.yml
