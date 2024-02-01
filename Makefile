createmigrations:
	@echo "Creating migrations"
	@atlas migrate diff --env gorm
	@echo "Migrations creation done"

migrate:
	@echo "Migrating..."
	@atlas migrate apply --env gorm
	@echo "Migrations done"

start:
	@which CompileDaemon > /dev/null 2>&1 && CompileDaemon --build="go build -o ./bin/app ./cmd/" --command=./bin/app || (echo "Using go run" && go run ./cmd/main.go)
	
test:
	# Set env to test
	@export ENV=test
	@go test -v ./...