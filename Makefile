setup:
	# Create go environment
	go mod init venv

install:
	# Install library
	go mod download

test:
	# Run test
	go test ./src/controllers -v

lint:
	# See local hadolint install instructions:   https://github.com/hadolint/hadolint
	# This is linter for Dockerfiles
	hadolint Dockerfile
	# This is a linter for Golang source code linter: https://golangci-lint.run/
	golangci-lint run --disable-all -E errcheck

