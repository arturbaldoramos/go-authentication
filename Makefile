build:
	@go build -o bin/go-authentication.exe ./cmd/app/main.go

run: build
	@./bin/go-authentication

clean:
	@go clean
	rmdir /s /q bin

db_up:
	docker-compose up -d

db_down:
	docker-compose stop