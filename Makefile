start:
	go run cmd/app/main.go

copy-hook:
	ln pre-commit-hook .git/hooks/pre-commit